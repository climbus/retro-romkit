// Package tree provides directory traversal functionality with filtering capabilities.
package tree

import (
	"os"
	"path/filepath"
	"strings"
)

// Entry represents a single tree entry
type Entry struct {
	Name   string
	Depth  int
	IsDir  bool
	Folder string
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
		if strings.Contains(file, string(os.PathSeparator)+".") || strings.HasPrefix(file, ".") {
			return nil // Skip hidden files and directories
		}
		if file == path {
			return nil // Skip the root directory itself
		}

		if !hasOneOfFileTypes(file, filetypes) && !info.IsDir() {
			return nil // Skip files that don't match the specified file types
		}

		relFilename, err := filepath.Rel(path, file)
		if err != nil {
			return err
		}
		depth := len(strings.Split(relFilename, string(os.PathSeparator))) - 1

		var name string
		if info.IsDir() {
			name = relFilename
		} else {
			name = filepath.Base(relFilename)
		}
		folder := filepath.Dir(relFilename)

		entries <- Entry{
			Name:   name,
			Depth:  depth,
			IsDir:  info.IsDir(),
			Folder: folder,
		}

		return nil
	})

	return err
}
