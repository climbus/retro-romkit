// Package tosec provides functionality for analyzing and displaying file trees and statistics.
package tosec

import (
	"fmt"
	"strings"

	"tosec-manager/internal/tree"
)

// GetFileTree returns a channel of tree entries for the given path
func GetFileTree(path string) (<-chan tree.Entry, <-chan error) {
	entries := make(chan tree.Entry, 100)
	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)
		if err := tree.Walk(path, []string{}, entries); err != nil {
			errCh <- err
		}
	}()

	return entries, errCh
}

// FormatTree returns a channel of formatted text lines for the tree
func FormatTree(path string) <-chan string {
	lines := make(chan string, 100)

	go func() {
		defer close(lines)

		lines <- fmt.Sprintf("Showing file tree for: %s", path)

		entries, errCh := GetFileTree(path)
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
func GetStats(path string) (Stats, error) {
	stats := Stats{
		TotalFiles:      0,
		DirectoryCounts: make(map[string]int),
	}
	stats.DirectoryCounts["/"] = 0 // Initialize root directory count

	entries, errCh := GetFileTree(path)
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
