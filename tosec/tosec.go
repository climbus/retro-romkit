package tosec

import (
	"fmt"
	"strings"
	
	"tosec-manager/tree"
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

