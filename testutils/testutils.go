// Package testutils provides utility functions for creating test files and directories.
package testutils

import (
	"os"
	"path/filepath"
	"testing"
)

func CreateTestFiles(t *testing.T, testFiles []string, tmpDir string) {
	for _, file := range testFiles {
		fullPath := filepath.Join(tmpDir, file)
		err := os.MkdirAll(filepath.Dir(fullPath), 0755)
		if err != nil {
			t.Fatalf("Failed to create dir for %s: %v", file, err)
		}

		f, err := os.Create(fullPath)
		if err != nil {
			t.Fatalf("Failed to create file %s: %v", file, err)
		}
		f.Close()
	}
}

func CreateTempDir(t *testing.T) string {
	tmpDir, err := os.MkdirTemp("", "tree_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	return tmpDir
}
