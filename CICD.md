# CI/CD Documentation

## Overview

This project uses GitHub Actions for continuous integration and automated releases with **semantic versioning** based on **Conventional Commits**. The CI/CD pipeline automatically analyzes commits, determines version bumps, generates changelogs, builds binaries, and creates releases.

## 🤖 Automated Release Process

### How It Works

1. **Developer commits** using Conventional Commits format
2. **CI validates** commit messages in pull requests
3. **On merge to main**, semantic-release:
   - Analyzes commit messages
   - Determines version bump (major/minor/patch)
   - Generates CHANGELOG.md
   - Creates git tag
   - Creates GitHub release
4. **Build job** compiles binaries for all platforms
5. **Assets uploaded** to GitHub release

### No Manual Tagging Required! 🎉

Version numbers and tags are created automatically based on your commit messages.

## 📝 Conventional Commits

This project follows the [Conventional Commits](https://www.conventionalcommits.org/) specification.

### Commit Message Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Commit Types

| Type | Description | Version Bump | Example |
|------|-------------|--------------|---------|
| `feat` | New feature | Minor (0.X.0) | `feat: add copy command` |
| `fix` | Bug fix | Patch (0.0.X) | `fix: resolve memory leak` |
| `perf` | Performance improvement | Patch (0.0.X) | `perf: optimize file parsing` |
| `refactor` | Code refactoring | Patch (0.0.X) | `refactor: simplify error handling` |
| `docs` | Documentation only | No release | `docs: update README` |
| `style` | Code style/formatting | No release | `style: run gofmt` |
| `test` | Add/update tests | No release | `test: add parser tests` |
| `build` | Build system changes | No release | `build: update Makefile` |
| `ci` | CI configuration | No release | `ci: add commitlint` |
| `chore` | Other changes | No release | `chore: update dependencies` |
| `revert` | Revert previous commit | Patch (0.0.X) | `revert: undo feature X` |

### Breaking Changes

For **MAJOR** version bumps (X.0.0), use one of these formats:

**Option 1: Exclamation mark**
```
feat!: change API response format
```

**Option 2: BREAKING CHANGE footer**
```
feat: change API response format

BREAKING CHANGE: API now returns JSON instead of XML
```

### Examples

✅ **Good commit messages:**
```
feat: add user authentication
fix: resolve race condition in file parser
docs: add installation instructions
refactor: extract common flag parsing function
perf: cache compiled regex patterns
feat(cli): add --output flag to copy command
fix(tosec)!: change ParseFileName return type

BREAKING CHANGE: ParseFileName now returns error as second value
```

❌ **Bad commit messages:**
```
Update code
Fixed bug
WIP
asdf
add feature
```

### Version Bump Examples

**Starting version: 1.2.3**

| Commits | New Version | Reason |
|---------|-------------|--------|
| `fix: bug fix` | 1.2.4 | Patch bump |
| `feat: new feature` | 1.3.0 | Minor bump |
| `feat!: breaking change` | 2.0.0 | Major bump |
| `docs: update README` | 1.2.3 | No release |
| `fix: bug A`<br>`fix: bug B` | 1.2.4 | Patch bump (multiple) |
| `feat: feature A`<br>`fix: bug B` | 1.3.0 | Minor bump (highest) |

## Workflows

### 1. CI Workflow (`.github/workflows/ci.yml`)

**Triggers:**
- Push to `main` branch
- Pull requests to `main` branch

**Jobs:**

#### Test Job
- Runs on Ubuntu latest
- Sets up Go 1.24
- Caches Go modules for faster builds
- Runs `go vet` for static analysis
- Checks code formatting with `gofmt`
- Runs tests with race detection and generates coverage report
- Uploads coverage to Codecov
- Verifies the project builds successfully

#### Lint Job
- Runs on Ubuntu latest
- Sets up Go 1.24
- Runs `golangci-lint` with configured linters (see `.golangci.yml`)

### 2. Commitlint Workflow (`.github/workflows/commitlint.yml`)

**Trigger:** Pull requests

**Purpose:** Validates that all commits in a PR follow Conventional Commits format

**Behavior:**
- Checks each commit message in the PR
- Fails if any commit doesn't follow the format
- Adds helpful comment to PR with examples
- Prevents merging of invalid commits

### 3. Release Workflow (`.github/workflows/release.yml`)

**Trigger:** Push to `main` branch (automatic after PR merge)

**Process:**

#### Job 1: Semantic Release
1. Checks out code with full history
2. Installs semantic-release and plugins
3. Analyzes commits since last release
4. Determines version bump (major/minor/patch/none)
5. Generates CHANGELOG.md
6. Creates VERSION file
7. Commits CHANGELOG and VERSION to repository
8. Creates git tag (e.g., `v1.2.3`)
9. Creates GitHub release with generated notes

#### Job 2: Build (only if new release created)
1. Checks out code at new version
2. Sets up Go 1.24
3. Runs all tests
4. Builds binaries for 6 platforms:
   - Linux (amd64, arm64)
   - macOS (amd64, arm64/Apple Silicon)
   - Windows (amd64, arm64)
5. Creates compressed archives (tar.gz, zip)
6. Generates SHA256 checksums
7. Uploads all artifacts to GitHub release

## Creating a Release

### The New Way (Automatic) ✨

1. **Write code with conventional commits:**
   ```bash
   git add .
   git commit -m "feat: add new awesome feature"
   git push origin your-branch
   ```

2. **Create PR and merge to main:**
   - Commitlint validates your messages
   - CI runs tests
   - Merge PR

3. **Automatic release happens:**
   - Semantic-release analyzes commits
   - Version determined automatically
   - CHANGELOG generated
   - Tag created
   - Release published
   - Binaries built and uploaded

### The Old Way (Manual - Still Works)

You can still manually create tags if needed:
```bash
git tag v1.0.0
git push origin v1.0.0
```

But semantic versioning is recommended! 🚀

## Version Numbering

We follow [Semantic Versioning](https://semver.org/):
- `vMAJOR.MINOR.PATCH`
- Example: `v1.2.3`
  - MAJOR: Breaking changes (from `feat!:` or `BREAKING CHANGE:`)
  - MINOR: New features (from `feat:`)
  - PATCH: Bug fixes (from `fix:`, `perf:`, `refactor:`)

## Configuration Files

### `.releaserc.json`
Semantic-release configuration:
- Commit analysis rules
- Changelog generation settings
- Plugin configuration
- Release note formatting

### `commitlint.config.js`
Commitlint configuration:
- Validates commit message format
- Enforces Conventional Commits
- Used in PR validation

### `.golangci.yml`
Linter configuration with enabled checks:
- Code formatting (gofmt)
- Static analysis (govet, staticcheck)
- Error checking (errcheck)
- Unused code detection
- Complexity analysis (gocyclo)
- Spelling (misspell)
- And more...

## Makefile Commands

The enhanced Makefile provides several commands for local development:

### Build Commands
```bash
make build          # Build for current platform
make build-all      # Build for all platforms (Linux, macOS, Windows)
make install        # Install to $GOPATH/bin
```

### Test Commands
```bash
make test           # Run tests with race detection
make test-coverage  # Generate HTML coverage report
```

### Quality Commands
```bash
make fmt            # Format code with gofmt
make vet            # Run go vet
make lint           # Run all linters (requires golangci-lint)
```

### Release Commands
```bash
make release        # Build all platforms and create archives
make clean          # Remove build artifacts
```

### Help
```bash
make help           # Show all available commands
```

## Platform Support

### Supported Platforms
- **Linux**: amd64, arm64
- **macOS**: amd64 (Intel), arm64 (Apple Silicon)
- **Windows**: amd64, arm64

### Binary Naming Convention
- Linux: `romkit-linux-{arch}`
- macOS: `romkit-darwin-{arch}`
- Windows: `romkit-windows-{arch}.exe`

## Continuous Integration Checks

Every push and pull request runs:
- ✅ Commit message validation (Conventional Commits)
- ✅ Static analysis (`go vet`)
- ✅ Code formatting check (`gofmt`)
- ✅ Race condition detection
- ✅ Unit tests with coverage
- ✅ Build verification
- ✅ Linting (golangci-lint)

## Troubleshooting

### Release Not Created

**Problem:** Pushed to main but no release created
- **Reason:** No commits that trigger a release (e.g., only `docs:`, `chore:`, etc.)
- **Solution:** Commits must use `feat:`, `fix:`, `perf:`, or `refactor:` to trigger releases

### Commitlint Fails in PR

**Problem:** PR blocked by commitlint
- **Solution:**
  1. Check commit messages follow Conventional Commits format
  2. Use interactive rebase to fix messages:
     ```bash
     git rebase -i HEAD~3  # Fix last 3 commits
     # Change 'pick' to 'reword' for commits to fix
     # Save and update commit messages
     git push --force-with-lease
     ```

### Wrong Version Bump

**Problem:** Expected minor version but got patch
- **Solution:** Make sure you used `feat:` not `fix:`
  - `feat:` → minor (0.X.0)
  - `fix:` → patch (0.0.X)
  - `feat!:` → major (X.0.0)

### Manual Tag Conflicts

**Problem:** Manually created tag conflicts with automatic tag
- **Solution:** Don't mix manual and automatic tags. Choose one approach:
  - Remove manual tag: `git tag -d v1.0.0 && git push --delete origin v1.0.0`
  - Or stick to manual tags only

## Best Practices

### Commit Messages
1. ✅ **Use conventional commits** for all commits
2. ✅ **Be descriptive** in subject line
3. ✅ **Add body** for complex changes
4. ✅ **Reference issues** in footer (e.g., `Fixes #123`)
5. ✅ **Mark breaking changes** explicitly

### Development Workflow
1. **Create feature branch** from main
2. **Make commits** using conventional format
3. **Test locally:** `make test`
4. **Format code:** `make fmt`
5. **Run linters:** `make lint`
6. **Create PR** to main
7. **Wait for CI** to validate commits
8. **Merge PR**
9. **Automatic release** happens!

### Example Workflow

```bash
# Create feature branch
git checkout -b feat/add-export-command

# Make changes and commit using conventional format
git add .
git commit -m "feat(cli): add export command for ROM metadata"

# Add tests
git add .
git commit -m "test(cli): add tests for export command"

# Update docs
git add .
git commit -m "docs: document export command usage"

# Push and create PR
git push origin feat/add-export-command

# After PR approval and merge to main:
# → Semantic-release analyzes commits
# → Sees "feat:" commit → Minor version bump (e.g., 1.2.0 → 1.3.0)
# → Generates changelog entry
# → Creates tag v1.3.0
# → Creates GitHub release
# → Builds and uploads binaries
```

## GitHub Secrets

No additional secrets required! The workflows use the default `GITHUB_TOKEN` which is automatically provided by GitHub Actions.

## Changelog

The `CHANGELOG.md` file is automatically generated and updated with each release. It includes:
- 🚀 Features
- 🐛 Bug Fixes
- ⚡ Performance Improvements
- ♻️ Code Refactoring
- 📝 Documentation
- ⏪ Reverts
- 💥 Breaking Changes

## Future Improvements

Potential enhancements to consider:
- [ ] Docker image builds and publication to GitHub Container Registry
- [ ] Homebrew formula auto-update
- [ ] Snap/Flatpak packaging for Linux
- [ ] Code signing for macOS and Windows binaries
- [ ] Performance benchmarking in CI
- [ ] Automated dependency updates (Dependabot)
- [ ] Release announcements to Slack/Discord

## Additional Resources

- [Conventional Commits Specification](https://www.conventionalcommits.org/)
- [Semantic Versioning](https://semver.org/)
- [Semantic Release](https://semantic-release.gitbook.io/)
- [Commitlint](https://commitlint.js.org/)
