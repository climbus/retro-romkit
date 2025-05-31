package tree

import (
	"os"
	"path/filepath"
	"strings"
)

// Entry represents a single tree entry
type Entry struct {
	Name  string
	Depth int
	IsDir bool
}

func hasOneOfFileTypes(file string, filetypes []string) bool {
	if len(filetypes) == 0 {
		return true // No file types specified, so no files to skip
	}
	for _, filetype := range filetypes {
		if strings.HasSuffix(file, filetype) {
			return true // File matches one of the specified types
		}
	}
	return false // No match found
}

// Walk traverses the directory tree and sends entries to the provided channel
func Walk(path string, filetypes []string, entries chan<- Entry) error {
	defer close(entries)
	
	err := filepath.WalkDir(path, func(file string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !hasOneOfFileTypes(file, filetypes) && !info.IsDir() {
			return nil // Skip files that don't match the specified file types
		}
		if path == file {
			return nil // Skip the root directory itself
		}

		relFilename, err := filepath.Rel(path, file)
		if err != nil {
			return err
		}
		depth := len(strings.Split(relFilename, string(os.PathSeparator))) - 1
		name := filepath.Base(relFilename)

		entries <- Entry{
			Name:  name,
			Depth: depth,
			IsDir: info.IsDir(),
		}

		return nil
	})

	return err
}

