# ─── Variables ──────────────────────────────────────────────
APP_NAME    := crm_go
MAIN        := ./main.go
BINARY      := ./bin/$(APP_NAME)
DB_DSN      := postgres://crm_user:crm_password@localhost:5432/crm?sslmode=disable
MIGRATIONS  := db/postgres/migrations

# ─── Dev ────────────────────────────────────────────────────

## run: Run the application (go run)
.PHONY: run
run:
	go run $(MAIN)

## build: Build the binary
.PHONY: build
build:
	go build -o $(BINARY) $(MAIN)

## clean: Remove build artifacts
.PHONY: clean
clean:
	rm -rf ./bin

# ─── Docker ─────────────────────────────────────────────────

## up: Start docker-compose services (postgres, redis, ...)
.PHONY: up
up:
	docker compose up -d

## down: Stop docker-compose services
.PHONY: down
down:
	docker compose down

## restart: Restart docker-compose services
.PHONY: restart
restart: down up

## logs: Follow docker-compose logs
.PHONY: logs
logs:
	docker compose logs -f

## ps: Show running docker-compose containers
.PHONY: ps
ps:
	docker compose ps

# ─── Database ───────────────────────────────────────────────

## migrate-up: Apply all up migrations
.PHONY: migrate-up
migrate-up:
	@for f in $(MIGRATIONS)/*.up.sql; do \
		echo "Applying $$f ..."; \
		docker exec -i crm_postgres psql -U crm_user -d crm < "$$f"; \
	done

## migrate-down: Apply all down migrations
.PHONY: migrate-down
migrate-down:
	@for f in $(MIGRATIONS)/*.down.sql; do \
		echo "Reverting $$f ..."; \
		docker exec -i crm_postgres psql -U crm_user -d crm < "$$f"; \
	done

## db-shell: Open psql shell inside postgres container
.PHONY: db-shell
db-shell:
	docker exec -it crm_postgres psql -U crm_user -d crm

## db-reset: Drop and recreate database tables
.PHONY: db-reset
db-reset: migrate-down migrate-up

# ─── Swagger ────────────────────────────────────────────────

## swagger: Generate swagger docs
.PHONY: swagger
swagger:
	swag init --parseDependency --parseInternal

## swagger-fmt: Format swagger annotations
.PHONY: swagger-fmt
swagger-fmt:
	swag fmt

# ─── Code Quality ──────────────────────────────────────────

## fmt: Format Go source files
.PHONY: fmt
fmt:
	go fmt ./...

## vet: Run go vet
.PHONY: vet
vet:
	go vet ./...

## lint: Run golangci-lint (install: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
.PHONY: lint
lint:
	golangci-lint run ./...

## test: Run all tests
.PHONY: test
test:
	go test -v ./...

## test-cover: Run tests with coverage report
.PHONY: test-cover
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# ─── Dependencies ──────────────────────────────────────────

## deps: Download and tidy dependencies
.PHONY: deps
deps:
	go mod download
	go mod tidy

# ─── Composite ─────────────────────────────────────────────

## dev: Start everything for development (docker + swagger + run)
.PHONY: dev
dev: up swagger run

## setup: First time project setup (docker + wait + migrate + swagger + deps)
.PHONY: setup
setup: deps up
	@echo "Waiting for postgres to be ready..."
	@until docker exec crm_postgres pg_isready -U crm_user -d crm > /dev/null 2>&1; do sleep 1; done
	@$(MAKE) migrate-up
	@$(MAKE) swagger
	@echo ""
	@echo "✓ Setup complete! Run 'make run' to start the server."

## check: Run all quality checks (fmt + vet + test)
.PHONY: check
check: fmt vet test

# ─── Help ───────────────────────────────────────────────────

## help: Show this help message
.PHONY: help
help:
	@echo ""
	@echo "Usage: make <target>"
	@echo ""
	@sed -n 's/^## /  /p' $(MAKEFILE_LIST) | column -t -s ':'
	@echo ""

.DEFAULT_GOAL := help
