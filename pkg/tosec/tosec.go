package tosec

import (
	"fmt"
	"strings"

	"tosec-manager/internal/tree"
)

// GetFileTree returns a channel of tree entries for the given path
func GetFileTree(path string) <-chan tree.Entry {
	entries := make(chan tree.Entry, 100)

	go func() {
		tree.Walk(path, []string{}, entries)
	}()

	return entries
}

// FormatTree returns a channel of formatted text lines for the tree
func FormatTree(path string) <-chan string {
	lines := make(chan string, 100)

	go func() {
		defer close(lines)

		lines <- fmt.Sprintf("Showing file tree for: %s", path)

		entries := GetFileTree(path)
		for entry := range entries {
			depthLabel := strings.Repeat("  ", entry.Depth)
			name := entry.Name
			if entry.IsDir {
				name += "/"
			}
			lines <- depthLabel + name
		}
	}()

	return lines
}

type Stats struct {
	TotalFiles      int
	DirectoryCounts map[string]int
}

// GetStats returns statistics about the files in the given path
func GetStats(path string) (Stats, error) {
	stats := Stats{
		TotalFiles:      0,
		DirectoryCounts: make(map[string]int),
	}
	stats.DirectoryCounts["/"] = 0 // Initialize root directory count

	entries := GetFileTree(path)
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

	return stats, nil
}
