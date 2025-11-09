# Retro RomKit Documentation

This directory contains auto-generated documentation for the Retro RomKit packages.

## Package Documentation

### Public Packages

- [tosec](packages/tosec.md) - Functionality for analyzing and displaying file trees and statistics for TOSEC ROM collections

### Internal Packages

- [tree](packages/tree.md) - Directory traversal functionality with filtering capabilities

## Generating Documentation

Documentation is automatically generated from Go source code comments using [gomarkdoc](https://github.com/princjef/gomarkdoc).

To regenerate the documentation:

```bash
make docs
```

Or manually:

```bash
gomarkdoc ./pkg/tosec > docs/packages/tosec.md
gomarkdoc ./internal/tree > docs/packages/tree.md
```

## Documentation Guidelines

When adding documentation to Go code:

1. **Package comments**: Start with `// Package <name>` describing the package purpose
2. **Type comments**: Describe what the type represents
3. **Function comments**: Start with the function name and describe what it does
4. **Exported items**: All exported types, functions, and methods should have comments

Example:

```go
// Package example provides example functionality.
package example

// User represents a user in the system.
type User struct {
    Name string
    Age  int
}

// NewUser creates a new User with the given name and age.
func NewUser(name string, age int) *User {
    return &User{Name: name, Age: age}
}
```

For more information on Go documentation conventions, see:
- [Effective Go - Commentary](https://golang.org/doc/effective_go#commentary)
- [Go Doc Comments](https://go.dev/doc/comment)
