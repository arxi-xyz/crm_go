CREATE TABLE IF NOT EXISTS users
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    uuid       UUID        NOT NULL DEFAULT gen_random_uuid(),

    phone      TEXT        NOT NULL,
    first_name TEXT,
    last_name  TEXT,
    password   TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX users_uuid_unique
    ON users (uuid);

CREATE UNIQUE INDEX users_phone_unique
    ON users (phone)
    WHERE deleted_at IS NULL;

CREATE INDEX users_created_at_idx
    ON users (created_at DESC);

-- todo: the default user must be read from config file
insert into users (phone, first_name, last_name, password, deleted_at)
values ('09130108631', 'hossein', 'sharif', 'secret', null);