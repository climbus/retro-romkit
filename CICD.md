# CI/CD Documentation

## Overview

This project uses GitHub Actions for continuous integration and deployment. The CI/CD pipeline automatically builds, tests, and releases the RetroRomkit application.

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

### 2. Release Workflow (`.github/workflows/release.yml`)

**Trigger:** Push of version tags (format: `v*`, e.g., `v1.0.0`)

**Process:**
1. Checks out code with full history
2. Sets up Go 1.24
3. Extracts version from git tag
4. Runs all tests
5. Builds binaries for multiple platforms:
   - Linux (amd64, arm64)
   - macOS (amd64, arm64/Apple Silicon)
   - Windows (amd64, arm64)
6. Creates compressed archives:
   - `.tar.gz` for Linux/macOS
   - `.zip` for Windows
7. Generates SHA256 checksums
8. Generates changelog from git commits
9. Creates GitHub release with all artifacts

## Creating a Release

To create a new release:

1. **Commit all changes** to the main branch
2. **Create and push a version tag:**
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```
3. **GitHub Actions will automatically:**
   - Run tests
   - Build binaries for all platforms
   - Create release archives
   - Generate checksums
   - Create a GitHub release
   - Upload all artifacts

## Version Numbering

We follow [Semantic Versioning](https://semver.org/):
- `vMAJOR.MINOR.PATCH`
- Example: `v1.2.3`
  - MAJOR: Breaking changes
  - MINOR: New features (backwards compatible)
  - PATCH: Bug fixes

## Makefile Commands

The enhanced Makefile provides several commands for local development and release preparation:

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

## Local Release Testing

To test the release process locally:

```bash
# Build for all platforms
make build-all

# Create release archives
make release

# Check generated files
ls -lh bin/
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
- ✅ Static analysis (`go vet`)
- ✅ Code formatting check (`gofmt`)
- ✅ Race condition detection
- ✅ Unit tests with coverage
- ✅ Build verification
- ✅ Linting (golangci-lint)

## Configuration Files

### `.golangci.yml`
Linter configuration with enabled checks:
- Code formatting (gofmt)
- Static analysis (govet, staticcheck)
- Error checking (errcheck)
- Unused code detection
- Complexity analysis (gocyclo)
- Spelling (misspell)
- And more...

### `.github/workflows/ci.yml`
Continuous integration workflow for testing and linting.

### `.github/workflows/release.yml`
Automated release workflow triggered by version tags.

## Troubleshooting

### Release Workflow Fails

**Problem:** Release workflow fails to create release
- **Solution:** Ensure the tag follows the format `v*` (e.g., `v1.0.0`)
- Check that tests pass locally: `make test`
- Verify builds work: `make build-all`

### Lint Failures

**Problem:** Linting fails in CI
- **Solution:** Run locally: `make lint`
- Fix formatting: `make fmt`
- Address vet warnings: `make vet`

### Test Failures

**Problem:** Tests fail in CI but pass locally
- **Solution:** Run with race detection: `go test -race ./...`
- Check for platform-specific issues
- Ensure all dependencies are properly versioned in `go.mod`

## Best Practices

1. **Always run tests before pushing:** `make test`
2. **Format code before committing:** `make fmt`
3. **Run linters locally:** `make lint`
4. **Test cross-compilation:** `make build-all`
5. **Review changelog before release:** Check git log since last tag
6. **Use semantic versioning:** Follow MAJOR.MINOR.PATCH format

## GitHub Secrets

No additional secrets are required. The workflows use the default `GITHUB_TOKEN` which is automatically provided by GitHub Actions.

## Codecov Integration

Coverage reports are automatically uploaded to Codecov for pull requests and main branch commits. To view coverage:

1. Visit the Codecov dashboard for this repository
2. Badge can be added to README.md (optional)

## Future Improvements

Potential enhancements to consider:
- [ ] Docker image builds and publication to GitHub Container Registry
- [ ] Automated changelog generation from conventional commits
- [ ] Release notes templates
- [ ] Homebrew formula auto-update
- [ ] Snap/Flatpak packaging for Linux
- [ ] Code signing for macOS and Windows binaries
- [ ] Performance benchmarking in CI
