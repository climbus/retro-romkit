package tosec

import (
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

func TestParseFileName(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     *TosecFile
		wantErr  bool
	}{
		{
			"Test standad filename",
			"Zynaps (1987)(Hewson Consultants).zip",
			&TosecFile{
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
			&TosecFile{
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
			&TosecFile{
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
