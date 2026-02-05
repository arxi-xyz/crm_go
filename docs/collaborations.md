# Collaborations

## Commit Messages

Please use this format for your commit messages:

```cmd
git commit -m "<type>(<scope>): <short, imperative description>"
e.g:
feat(auth): add login endpoint
fix(users): handle duplicate phone numbers
```

### Operations

- feat
    Add a new capability or behavior
    Feature does NOT need to be fully complete

- fix
    Fix a bug or incorrect behavior
    Applies to both development and production issues

- hotfix
    Critical fix for production incidents
    Use sparingly and only for urgent cases

- refactor
    Code change that neither fixes a bug nor adds a feature

- chore
    Maintenance tasks (deps, configs, tooling)

- wip
    Temporary commits
    Must not be merged into main branches

### Rules

- wip commits are allowed only on feature branches
- main branch must not contain wip commits
- Keep commits small and focused
- One logical change per commit

---
