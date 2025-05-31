package tree

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDisplay(t *testing.T) {
	// Create temporary directory structure for testing
	tmpDir := createTempDir(t)

	// Create test structure
	testFiles := []string{
		"file1.txt",
		"file2.jpg",
		"README.md",
		"subdir/file3.txt",
		"subdir/file4.png",
		"subdir/nested/file5.txt",
	}

	createTestFiles(t, testFiles, tmpDir)

	tests := []struct {
		name      string
		path      string
		filetypes []string
		wantError bool
	}{
		{
			name:      "valid directory without filter",
			path:      tmpDir,
			filetypes: []string{},
			wantError: false,
		},
		{
			name:      "valid directory with single extension filter",
			path:      tmpDir,
			filetypes: []string{".txt"},
			wantError: false,
		},
		{
			name:      "valid directory with multiple extension filter",
			path:      tmpDir,
			filetypes: []string{".jpg", ".png"},
			wantError: false,
		},
		{
			name:      "valid directory with non-matching filter",
			path:      tmpDir,
			filetypes: []string{".pdf", ".doc"},
			wantError: false,
		},
		{
			name:      "valid directory with full filename filter",
			path:      tmpDir,
			filetypes: []string{"README.md"},
			wantError: false,
		},
		{
			name:      "non-existent directory",
			path:      "/non/existent/path/that/definitely/does/not/exist",
			filetypes: []string{},
			wantError: true,
		},
		{
			name:      "invalid path with null bytes",
			path:      "/tmp/test\x00invalid",
			filetypes: []string{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Display(tt.path, tt.filetypes)
			if (err != nil) != tt.wantError {
				t.Errorf("Display() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func createTestFiles(t *testing.T, testFiles []string, tmpDir string) {
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

func createTempDir(t *testing.T) string {
	tmpDir, err := os.MkdirTemp("", "tree_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)
	return tmpDir
}

