package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func hasOneOfFileTypes(file string, filetypes []string) bool {
	if len(filetypes) == 0 {
		return false // No file types specified, so no files to skip
	}
	for _, filetype := range filetypes {
		if strings.HasSuffix(file, filetype) {
			return true // File matches one of the specified types
		}
	}
	return false // No match found
}

// Display shows the directory tree structure for the given path
func Display(path string, filetypes []string) error {
	fmt.Printf("Showing file tree for: %s\n", path)
	err := filepath.WalkDir(path, func(file string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !hasOneOfFileTypes(file, filetypes) && !info.IsDir() {
			return nil // Skip files that match the specified file types
		}
		if path == file {
			return nil // Skip the root directory itself
		}

		relFilename, err := filepath.Rel(path, file)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err) // coverage-ignore
		}
		depth := len(strings.Split(relFilename, string(os.PathSeparator))) - 1
		depthLabel := strings.Repeat("  ", depth)
		name := filepath.Base(relFilename)
		if info.IsDir() {
			name += "/"
		}

		fmt.Println(depthLabel + name)

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}

	return nil
}
