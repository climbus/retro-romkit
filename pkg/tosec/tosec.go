// Package tosec provides functionality for analyzing and displaying file trees and statistics.
package tosec

import (
	"fmt"
	"strings"

	"tosec-manager/internal/tree"
)

type Tosec struct {
	Path      string
	Platform  string
	FileTypes []string
}

func Create(path, platform string) *Tosec {
	platforms := map[string][]string{
		"amiga":   {"adf", "dms", "ipf", "lha", "lzx"},
		"atari":   {"st", "msa", "zip"},
		"c64":     {"d64", "t64", "prg", "crt"},
		"nes":     {"nes", "unif"},
		"gameboy": {"gb", "gbc", "gba"},
		"sega":    {"md", "smd", "gen", "bin"},
		"pc":      {"exe", "com", "bat", "zip", "rar"},
		"psx":     {"iso", "bin", "cue"},
		"coleco":  {"col", "rom"},
		"golang":  {"go"},
	}

	return &Tosec{
		Path:      path,
		Platform:  platform,
		FileTypes: platforms[platform],
	}
}

func (t *Tosec) FileTypesWithArchives() []string {
	if len(t.FileTypes) == 0 {
		return t.FileTypes
	}
	// Add common archive formats to the file types
	archiveTypes := []string{"zip", "rar", "7z", "tar", "gz", "bz2"}
	fileTypes := make([]string, len(t.FileTypes)+len(archiveTypes))
	copy(fileTypes, t.FileTypes)
	copy(fileTypes[len(t.FileTypes):], archiveTypes)
	return fileTypes
}

// GetFileTree returns a channel of tree entries for the given path
func (tosecFolder *Tosec) GetFileTree() (<-chan tree.Entry, <-chan error) {
	entries := make(chan tree.Entry, 100)
	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)
		if err := tree.Walk(tosecFolder.Path, tosecFolder.FileTypesWithArchives(), entries); err != nil {
			errCh <- err
		}
	}()

	return entries, errCh
}

// FormatTree returns a channel of formatted text lines for the tree
func (t *Tosec) FormatTree() <-chan string {
	lines := make(chan string, 100)

	go func() {
		defer close(lines)

		lines <- fmt.Sprintf("Showing file tree for: %s", t.Path)

		entries, errCh := t.GetFileTree()
		for entry := range entries {
			depthLabel := strings.Repeat("  ", entry.Depth)
			name := entry.Name
			if entry.IsDir {
				name += "/"
			}
			lines <- depthLabel + name
		}

		// Check for errors after processing entries
		select {
		case err := <-errCh:
			if err != nil {
				lines <- fmt.Sprintf("Error: %v", err)
			}
		default:
			// No error
		}
	}()

	return lines
}

type Stats struct {
	TotalFiles      int
	DirectoryCounts map[string]int
}

// GetStats returns statistics about the files in the given path
func (t *Tosec) GetStats() (Stats, error) {
	stats := Stats{
		TotalFiles:      0,
		DirectoryCounts: make(map[string]int),
	}
	stats.DirectoryCounts["/"] = 0 // Initialize root directory count

	entries, errCh := t.GetFileTree()
	for entry := range entries {
		if entry.IsDir {
			stats.DirectoryCounts[entry.Name] = 0
		} else {
			stats.TotalFiles++
			if entry.Depth > 0 {
				stats.DirectoryCounts[entry.Folder]++
			} else {
				stats.DirectoryCounts["/"]++ // Root directory
			}
		}
	}

	// Check for errors after processing entries
	select {
	case err := <-errCh:
		if err != nil {
			return stats, err
		}
	default:
		// No error
	}

	return stats, nil
}
