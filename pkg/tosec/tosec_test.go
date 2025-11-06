package tosec

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/climbus/retro-romkit/testutils"
)

func TestGetStatsEmptyDirectory(t *testing.T) {
	tmpDir := testutils.CreateTempDir(t)
	defer os.RemoveAll(tmpDir)

	tosecFolder := Create(tmpDir, "")

	stats, err := tosecFolder.GetStats()
	if err != nil {
		t.Fatalf("GetStats failed: %v", err)
	}

	if stats.TotalFiles != 0 {
		t.Errorf("Expected 0 total files, got %d", stats.TotalFiles)
	}

	if stats.DirectoryCounts["/"] != 0 {
		t.Errorf("Expected 0 files in root directory, got %d", stats.DirectoryCounts["/"])
	}
}

func TestGetStatsOnlyDirectories(t *testing.T) {
	tmpDir := testutils.CreateTempDir(t)
	defer os.RemoveAll(tmpDir)

	err := os.MkdirAll(filepath.Join(tmpDir, "dir1"), 0755)
	if err != nil {
		t.Fatalf("Failed to create dir1: %v", err)
	}
	err = os.MkdirAll(filepath.Join(tmpDir, "dir2", "subdir1"), 0755)
	if err != nil {
		t.Fatalf("Failed to create dir2/subdir1: %v", err)
	}

	tosecFolder := Create(tmpDir, "")

	stats, err := tosecFolder.GetStats()
	if err != nil {
		t.Fatalf("GetStats failed: %v", err)
	}

	if stats.TotalFiles != 0 {
		t.Errorf("Expected 0 total files, got %d", stats.TotalFiles)
	}

	if reflect.DeepEqual(Stats{
		TotalFiles:      3,
		DirectoryCounts: map[string]int{"/": 0, "dir1": 0, "dir1/subdir1": 0, "dir2": 0},
	}, stats.DirectoryCounts) {
		t.Errorf("Expected directory counts to match, got %v", stats.DirectoryCounts)
	}
}

func TestTosec_GetStats(t *testing.T) {

	tmpDir := testutils.CreateTempDir(t)
	defer os.RemoveAll(tmpDir)

	testFiles := []string{
		"file1.crt",
		"file2.txt",
		"dir1/file3.d64",
		"dir1/file4.txt",
		"dir1/subdir1/file5.prg",
		"dir2/file6.txt",
	}

	testutils.CreateTestFiles(t, testFiles, tmpDir)

	tests := []struct {
		name     string
		path     string
		platform string
		want     Stats
		wantErr  bool
	}{
		{
			name:     "Test filter by platform",
			path:     tmpDir,
			platform: "c64",
			want: Stats{
				TotalFiles:      3,
				DirectoryCounts: map[string]int{"/": 1, "dir1": 1, "dir1/subdir1": 1, "dir2": 0},
			},
			wantErr: false,
		},
		{
			name:     "Test filter by platform, whhen no files match",
			path:     tmpDir,
			platform: "coleco",
			want: Stats{
				TotalFiles:      0,
				DirectoryCounts: map[string]int{"/": 0, "dir1": 0, "dir1/subdir1": 0, "dir2": 0},
			},
			wantErr: false,
		},
		{
			name: "Test filter without platform",
			path: tmpDir,
			want: Stats{
				TotalFiles:      6,
				DirectoryCounts: map[string]int{"/": 2, "dir1": 2, "dir1/subdir1": 1, "dir2": 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			to := Create(tt.path, tt.platform)
			got, gotErr := to.GetStats()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetStats() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetStats() succeeded unexpectedly")
			}

			if got.TotalFiles != tt.want.TotalFiles || !reflect.DeepEqual(got.DirectoryCounts, tt.want.DirectoryCounts) {
				t.Errorf("GetStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFiles(t *testing.T) {
	tmpDir := testutils.CreateTempDir(t)
	defer os.RemoveAll(tmpDir)

	testFiles := []string{
		"Game One (1990)(Publisher A).zip",
		"Game Two (1991)(Publisher B)(Europe)(en).zip",
		"Game Three (1992)(Publisher C)[a].zip",
		"InvalidFileName.txt",
		"subdir/Game Four (1993)(Publisher D).zip",
	}

	testutils.CreateTestFiles(t, testFiles, tmpDir)

	tests := []struct {
		name          string
		platform      string
		expectedCount int
		wantErr       bool
	}{
		{
			name:          "Get all files without platform filter",
			platform:      "",
			expectedCount: 4, // 4 valid TOSEC files, 1 invalid
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tosecFolder := Create(tmpDir, tt.platform)
			files, err := tosecFolder.GetFiles()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(files) != tt.expectedCount {
				t.Errorf("GetFiles() got %d files, want %d", len(files), tt.expectedCount)
			}

			// Verify parsed files have correct structure
			for _, file := range files {
				if file.Title == "" {
					t.Errorf("File has empty title: %+v", file)
				}
				if file.Date == "" {
					t.Errorf("File has empty date: %+v", file)
				}
				if file.Publisher == "" {
					t.Errorf("File has empty publisher: %+v", file)
				}
			}
		})
	}
}

func TestFormatTree(t *testing.T) {
	tmpDir := testutils.CreateTempDir(t)
	defer os.RemoveAll(tmpDir)

	testFiles := []string{
		"file1.zip",
		"file2.zip",
		"subdir/file3.zip",
		"subdir/nested/file4.zip",
	}

	testutils.CreateTestFiles(t, testFiles, tmpDir)

	tests := []struct {
		name         string
		platform     string
		expectFiles  []string
		expectHeader bool
	}{
		{
			name:         "Format tree with all files",
			platform:     "",
			expectFiles:  []string{"file1.zip", "file2.zip", "subdir/", "file3.zip", "subdir/nested/", "file4.zip"},
			expectHeader: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tosecFolder := Create(tmpDir, tt.platform)
			lines := tosecFolder.FormatTree()

			var foundLines []string
			for line := range lines {
				foundLines = append(foundLines, line)
			}

			if len(foundLines) == 0 {
				t.Error("FormatTree() returned no lines")
			}

			// Check if header is present
			if tt.expectHeader {
				if len(foundLines) == 0 || !reflect.DeepEqual(foundLines[0], fmt.Sprintf("Showing file tree for: %s", tmpDir)) {
					t.Errorf("FormatTree() missing or incorrect header, got: %v", foundLines[0])
				}
			}

			// Verify some expected files are present in output
			output := ""
			for _, line := range foundLines {
				output += line + "\n"
			}

			for _, expectedFile := range tt.expectFiles {
				found := false
				for _, line := range foundLines {
					if reflect.DeepEqual(line, expectedFile) || reflect.DeepEqual(line, "  "+expectedFile) || reflect.DeepEqual(line, "    "+expectedFile) {
						found = true
						break
					}
				}
				if !found {
					t.Logf("Full output:\n%s", output)
					t.Errorf("FormatTree() missing expected file: %s", expectedFile)
				}
			}
		})
	}
}

func TestFormatTreeError(t *testing.T) {
	tosecFolder := Create("/non/existent/path", "")
	lines := tosecFolder.FormatTree()

	var foundLines []string
	var hasError bool
	for line := range lines {
		foundLines = append(foundLines, line)
		if len(line) >= 5 && line[:5] == "Error" {
			hasError = true
		}
	}

	if !hasError {
		t.Errorf("FormatTree() should return error line for non-existent path, got: %v", foundLines)
	}
}

func TestParseFileName(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     *File
		wantErr  bool
	}{
		{
			"Test standad filename",
			"Zynaps (1987)(Hewson Consultants).zip",
			&File{
				FileName:  "Zynaps (1987)(Hewson Consultants).zip",
				Title:     "Zynaps",
				Date:      "1987",
				Publisher: "Hewson Consultants",
				Region:    "",
				Language:  "",
				Format:    "zip",
				Flags:     []string{},
			},
			false,
		},
		{
			"Test filename with region and language",
			"Zynaps (1987)(Hewson Consultants)(Europe)(en).zip",
			&File{
				FileName:  "Zynaps (1987)(Hewson Consultants)(Europe)(en).zip",
				Title:     "Zynaps",
				Date:      "1987",
				Publisher: "Hewson Consultants",
				Region:    "Europe",
				Language:  "en",
				Format:    "zip",
				Flags:     []string{},
			},
			false,
		},
		{
			"Test filename with flags",
			"Zynaps (1987)(Hewson Consultants)[a][Aka kota].zip",
			&File{
				FileName:  "Zynaps (1987)(Hewson Consultants)[a][Aka kota].zip",
				Title:     "Zynaps",
				Date:      "1987",
				Publisher: "Hewson Consultants",
				Format:    "zip",
				Flags:     []string{"a", "Aka kota"},
			},
			false,
		},
		{
			"Test bad filename",
			"InvalidFileName.txt",
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, gotErr := ParseFileName(tt.fileName)

			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ParseFileName() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ParseFileName() succeeded unexpectedly")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
