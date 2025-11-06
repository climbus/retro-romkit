# Contributing to RetroRomkit

Thank you for your interest in contributing to RetroRomkit! 🎮

## 📝 Commit Message Guidelines

This project uses [Conventional Commits](https://www.conventionalcommits.org/) for automated versioning and changelog generation.

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- **feat**: New feature (triggers MINOR version bump)
- **fix**: Bug fix (triggers PATCH version bump)
- **perf**: Performance improvement (triggers PATCH version bump)
- **refactor**: Code refactoring (triggers PATCH version bump)
- **docs**: Documentation changes (no version bump)
- **style**: Code formatting (no version bump)
- **test**: Adding or updating tests (no version bump)
- **build**: Build system changes (no version bump)
- **ci**: CI configuration changes (no version bump)
- **chore**: Other changes (no version bump)
- **revert**: Revert previous commit (triggers PATCH version bump)

### Breaking Changes

For breaking changes (triggers MAJOR version bump):

```
feat!: change API response format
```

or

```
feat: change API response format

BREAKING CHANGE: API now returns JSON instead of XML
```

### Examples

✅ **Good:**
```
feat: add copy command with limit option
fix: resolve race condition in file parser
docs: add installation instructions
refactor: extract common flag parsing function
perf: cache compiled regex patterns
test: add tests for GetFiles method
```

❌ **Bad:**
```
update code
fixed bug
WIP
asdf
```

## 🔧 Development Workflow

### 1. Fork and Clone

```bash
git clone https://github.com/YOUR_USERNAME/retro-romkit.git
cd retro-romkit
```

### 2. Create Feature Branch

```bash
git checkout -b feat/your-feature-name
# or
git checkout -b fix/your-bug-fix
```

### 3. Make Changes

- Write your code
- Add tests for new features
- Ensure all tests pass: `make test`
- Format code: `make fmt`
- Run linters: `make lint`

### 4. Commit with Conventional Commits

```bash
git add .
git commit -m "feat: add awesome new feature"
```

### 5. Push and Create Pull Request

```bash
git push origin feat/your-feature-name
```

Then create a PR on GitHub.

### 6. CI Validation

Your PR will be automatically checked for:
- ✅ Conventional commit format
- ✅ Tests passing
- ✅ Code formatting
- ✅ Linting

### 7. Merge to Main

After approval and merge:
- Semantic-release automatically analyzes your commits
- Creates new version based on commit types
- Generates CHANGELOG entry
- Creates GitHub release
- Builds and uploads binaries

## 🧪 Testing

### Run Tests

```bash
make test           # Run all tests
make test-coverage  # Generate coverage report
```

### Write Tests

- Add tests in `*_test.go` files
- Use table-driven tests when possible
- Aim for high coverage (currently 89-96%)

## 🎨 Code Style

- Follow Go conventions
- Use `gofmt` for formatting: `make fmt`
- Run `go vet`: `make vet`
- Pass golangci-lint: `make lint`

## 📦 Building

```bash
make build          # Build for current platform
make build-all      # Build for all platforms
```

## 📚 Documentation

- Update README.md for user-facing changes
- Add comments to exported functions/types
- Update CICD.md for CI/CD changes

## 🐛 Reporting Bugs

Create an issue with:
- Clear description
- Steps to reproduce
- Expected vs actual behavior
- Your environment (OS, Go version)

## 💡 Suggesting Features

Create an issue with:
- Use case description
- Proposed solution
- Alternative approaches considered

## ✅ Pull Request Checklist

Before submitting a PR:

- [ ] Code follows Go conventions
- [ ] Tests added/updated
- [ ] All tests pass (`make test`)
- [ ] Code formatted (`make fmt`)
- [ ] Linters pass (`make lint`)
- [ ] Commit messages follow Conventional Commits
- [ ] Documentation updated if needed

## 🔄 Version Bumps

You don't need to manually update versions! The system automatically determines version bumps based on your commit types:

| Commit Type | Version Change | Example |
|-------------|----------------|---------|
| `fix:` | Patch (0.0.X) | 1.2.3 → 1.2.4 |
| `feat:` | Minor (0.X.0) | 1.2.3 → 1.3.0 |
| `feat!:` or `BREAKING CHANGE:` | Major (X.0.0) | 1.2.3 → 2.0.0 |

## 📄 License

By contributing, you agree that your contributions will be licensed under the project's license.

## 🙏 Thank You!

Your contributions make RetroRomkit better for everyone!
