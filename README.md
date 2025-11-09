# 🎮 The RetroRomkit

**Prepare your ROMset. Rule your retro.**

The RetroRomkit is a smart, lightweight toolkit for organizing, renaming, and structuring TOSEC-based ROM sets.  
Whether you're managing thousands of disk images for Amiga, Commodore 64, Atari ST, or other classic platforms — this tool helps you bring order to chaos.

---

## ✨ Features

- 🗂 **Parse TOSEC-style filenames**
- 🧠 **Extract metadata** (title, year, publisher, language, etc.)
- 📁 **Generate custom folder structures**
  - Sort by platform, letter, publisher, or custom rules
  - Split ROMs into folders of N files
- 🧪 **Preview mode** (dry-run) before applying changes
- 🔄 **Reversible** (no destructive operations)
- 🎯 Designed for **multi-platform** collections

---

## 🧰 Example Use Cases

- Prepare a clean, emulator-ready ROMset for MiSTer or RetroArch
- Organize your collection by platform, then by title
- Build subsets from full TOSEC dumps

---

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

## 📚 Documentation

Package documentation is available in the [docs/](docs/) directory:

- [tosec](docs/packages/tosec.md) - Core functionality for TOSEC ROM analysis
- [tree](docs/packages/tree.md) - Directory traversal utilities

To regenerate the documentation from source code:

```bash
make docs
```

Documentation is automatically updated on every push to main branch.

## Build

```bash
make build
```

The compiled binary will be available in the `bin/` directory.
