# retro-romkit

A CLI tool for managing and analyzing TOSEC file collections.

## Installation

```bash
make build
```

## Usage

```bash
romkit <command> [<args>]
```

### Commands

- `show <path>` - Show file tree of the specified path
- `stats <path>` - Show statistics about files in the specified path  
- `help` - Show help message

### Examples

```bash
# Display file tree
romkit show /path/to/directory

# Show file statistics
romkit stats /path/to/directory
```

## Build

```bash
make build
```

The compiled binary will be available in the `bin/` directory.
