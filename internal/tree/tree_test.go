package tree

import (
	"reflect"
	"testing"
	"github.com/climbus/retro-romkit/testutils"
)

func TestWalk(t *testing.T) {
	tmpDir := testutils.CreateTempDir(t)

	testFiles := []string{
		".git/config",
		"file1.txt",
		"file2.jpg",
		"README.md",
		"subdir/file3.txt",
		"subdir/file4.png",
		"subdir/nested/file5.txt",
	}

	testutils.CreateTestFiles(t, testFiles, tmpDir)

	tests := []struct {
		name          string
		filetypes     []string
		expectedFiles []string
		wantError     bool
	}{
		{
			name:          "no filter - show all files",
			filetypes:     []string{},
			expectedFiles: []string{"README.md", "file1.txt", "file2.jpg", "subdir", "file3.txt", "file4.png", "subdir/nested", "file5.txt"},
		},
		{
			name:          "filter txt files only",
			filetypes:     []string{".txt"},
			expectedFiles: []string{"file1.txt", "subdir", "file3.txt", "subdir/nested", "file5.txt"},
		},
		{
			name:          "filter jpg and png files",
			filetypes:     []string{".jpg", ".png"},
			expectedFiles: []string{"file2.jpg", "subdir", "file4.png", "subdir/nested"},
		},
		{
			name:          "filter non-existent extension - show only dirs",
			filetypes:     []string{".pdf"},
			expectedFiles: []string{"subdir", "subdir/nested"},
		},
		{
			name:          "filter by full filename",
			filetypes:     []string{"README.md"},
			expectedFiles: []string{"README.md", "subdir", "subdir/nested"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries := make(chan Entry, 100)

			go func() {
				err := Walk(tmpDir, tt.filetypes, entries)
				if tt.wantError && err == nil {
					t.Errorf("Expected error but got none")
				}
				if !tt.wantError && err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}()

			var foundFiles []string
			for entry := range entries {
				foundFiles = append(foundFiles, entry.Name)
			}

			equals := reflect.DeepEqual(foundFiles, tt.expectedFiles)
			if !equals {
				t.Errorf("Expected files %v, but got %v", tt.expectedFiles, foundFiles)
			}
		})
	}

	// Test error case
	t.Run("non-existent directory", func(t *testing.T) {
		entries := make(chan Entry, 100)

		err := Walk("/non/existent/path", []string{}, entries)
		if err == nil {
			t.Error("Expected error for non-existent directory")
		}

		// Drain channel
		for range entries {
		}
	})
}
