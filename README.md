# tosec-manager

A CLI tool for managing and analyzing TOSEC file collections.

## Installation

```bash
make build
```

## Usage

```bash
tosec-manager <command> [<args>]
```

### Commands

- `show <path>` - Show file tree of the specified path
- `stats <path>` - Show statistics about files in the specified path  
- `help` - Show help message

### Examples

```bash
# Display file tree
tosec-manager show /path/to/directory

# Show file statistics
tosec-manager stats /path/to/directory
```

## Build

```bash
make build
```

The compiled binary will be available in the `bin/` directory.
