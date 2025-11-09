# Changelog

All notable changes to this project will be documented in this file. See [Conventional Commits](https://conventionalcommits.org) for commit guidelines.

## [0.1.2](https://github.com/climbus/retro-romkit/compare/v0.1.1...v0.1.2) (2025-11-09)

### 🐛 Bug Fixes

* add write permissions for docs workflow ([abf643b](https://github.com/climbus/retro-romkit/commit/abf643ba72bc7951bff8b66a18667045a1adacdc))
* fixed help formatting ([caf27d4](https://github.com/climbus/retro-romkit/commit/caf27d4f78e3f558671d1fb37da7fa9e30438780))

### 📝 Documentation

* update generated documentation [skip ci] ([56a01cc](https://github.com/climbus/retro-romkit/commit/56a01cc6dcb38606e0401ca49f3efa1723357304))

## [0.1.1](https://github.com/climbus/retro-romkit/compare/v0.1.0...v0.1.1) (2025-11-09)

### 🐛 Bug Fixes

* fixed help formatting ([02f0cb8](https://github.com/climbus/retro-romkit/commit/02f0cb83b4656a027cff0c3694e2198f5338c2e4))

### 📝 Documentation

* add automated package documentation generation ([735b918](https://github.com/climbus/retro-romkit/commit/735b918378b19f10a358b65bbec2c295cade195f))

### ♻️ Code Refactoring

* extracted root dir const and removed comments ([d900a58](https://github.com/climbus/retro-romkit/commit/d900a580595aeb0929bdcc912f5d02ed584437d3))

## 0.1.0 (2025-11-06)

### ⚠ BREAKING CHANGES

* Releases are no longer created automatically on every push to main.

This commit changes the release process to be manually triggered, giving
developers control over when releases are published.

CHANGES:
- Release workflow now uses workflow_dispatch (manual trigger)
- No longer triggers on push to main
- Releases must be explicitly triggered using `make release` command

NEW MAKEFILE COMMANDS:
- make release: Triggers the release workflow (requires gh CLI)
- make release-dry-run: Preview what release would be created
- make build-archives: Build archives locally (renamed from release)

WORKFLOW IMPROVEMENTS:
- Added dry_run input option to test releases without publishing
- Added version_type input (auto/major/minor/patch) for future use
- Workflow can be triggered via GitHub UI or gh CLI

DOCUMENTATION UPDATES:
- CICD.md: Updated to reflect manual release process
  * Added step-by-step release instructions
  * Added 3 ways to trigger releases (Makefile/CLI/UI)
  * Updated troubleshooting section
- CONTRIBUTING.md: Added release section for maintainers
- Makefile: Updated help text with new commands

BENEFITS:
✅ Control over when releases happen
✅ Can accumulate multiple PRs before releasing
✅ Ability to preview releases with dry-run
✅ No surprise releases from documentation commits
✅ Maintainers decide release timing

USAGE:
# Preview what would be released
make release-dry-run

# Trigger a release
make release

# Or use GitHub Actions UI to manually run workflow

### 🚀 Features

* add automated semantic versioning with conventional commits ([501a7af](https://github.com/climbus/retro-romkit/commit/501a7af2f2ffc9a37f0d8dbe744f57e373652eeb))

### 🐛 Bug Fixes

* change release workflow to manual trigger instead of automatic ([f45eb8d](https://github.com/climbus/retro-romkit/commit/f45eb8d81d11a91c26fb772de0c61dbb68a9aee4))
* correct JSON syntax in .releaserc.json ([3b04bb3](https://github.com/climbus/retro-romkit/commit/3b04bb3cbf9f52b4462eb1e6861e0a369341da1f))
