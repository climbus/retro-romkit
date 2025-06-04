// Package tosec provides functionality for analyzing and displaying file trees and statistics.
package tosec

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/climbus/retro-romkit/internal/tree"
)

const REGEX_MAIN_DATA = `^(.*?) \((.*?)\)\((.*?)\).*\.(.*)$`
const REGEX_FLAG = `\[(.*?)\]`
const REGEX_OPTION = `\((.*?)\)`
const REGEX_LANGUAGE = `^` + LANGUAGE_NAMES + `(-` + LANGUAGE_NAMES + `)?$`

const REGEX_REGION = `(Japan|USA|Europe|World|International|Asia|Australia|Brazil|China|Korea|Taiwan)`
const LANGUAGE_NAMES = `(en|fr|de|es|it|ja|zh|ko|pt|ru|nl|pl|sv|no|da|fi|tr|ar|he|hi|th|vi|id|ms|cs|hu|ro|bg|el|uk|hr|sk|sl|lt|lv|et|fa|ur)`

type TosecFolder struct {
	Path      string
	Platform  string
	FileTypes []string
}

type TosecFile struct {
	FileName  string
	Title     string
	Date      string
	Publisher string
	Platform  string
	Format    string
	Flags     []string
	Region    string
	Language  string
}

type Stats struct {
	TotalFiles      int
	DirectoryCounts map[string]int
}

// ParseFileName parses a file name according to the TOSEC naming convention
func ParseFileName(fileName string) (*TosecFile, error) {

	// TODO: Move compiled regex to package level variable
	re := regexp.MustCompile(REGEX_MAIN_DATA)
	re_flags := regexp.MustCompile(REGEX_FLAG)
	re_options := regexp.MustCompile(REGEX_OPTION)

	matches := re.FindStringSubmatch(fileName)
	if matches == nil {
		return nil, errors.New("invalid file name format")
	}
	tf := &TosecFile{
		FileName:  fileName,
		Title:     strings.TrimSpace(matches[1]),
		Date:      strings.TrimSpace(matches[2]),
		Publisher: strings.TrimSpace(matches[3]),
		Format:    strings.TrimSpace(matches[4]),
	}

	rest := tf.extractRestPartOfName()

	flags_res := re_flags.FindAllStringSubmatch(rest, -1)
	flags := extractValues(flags_res)
	tf.Flags = flags

	options_res := re_options.FindAllStringSubmatch(rest, -1)
	options := extractValues(options_res)

	for _, opt := range options {
		opt = strings.TrimSpace(opt)
		if tf.Region == "" && regexp.MustCompile(REGEX_REGION).MatchString(opt) {
			tf.Region = opt
		} else if tf.Language == "" && regexp.MustCompile(REGEX_LANGUAGE).MatchString(opt) {
			tf.Language = opt
		}
	}

	// fmt.Println("Rest of the file name:", rest)
	// fmt.Println("Options", options)
	return tf, nil
}

// Create initializes a TosecFolder with the given path and platform.
func Create(path, platform string) *TosecFolder {

	// TODO: Move platform list to a package-level variable or config
	platforms := map[string][]string{
		"amiga":   {"adf", "dms", "ipf", "lha", "lzx"},
		"atari":   {"st", "msa", "zip"},
		"c64":     {"d64", "t64", "prg", "crt"},
		"nes":     {"nes", "unif"},
		"gameboy": {"gb", "gbc", "gba"},
		"sega":    {"md", "smd", "gen", "bin"},
		"pc":      {"exe", "com", "bat", "zip", "rar"},
		"psx":     {"iso", "bin", "cue"},
		"coleco":  {"col", "rom"},
		"golang":  {"go"},
	}

	return &TosecFolder{
		Path:      path,
		Platform:  platform,
		FileTypes: platforms[platform],
	}
}

// GetFileTree returns a channel of tree entries for the given path
func (tosecFolder *TosecFolder) GetFileTree() (<-chan tree.Entry, <-chan error) {
	entries := make(chan tree.Entry, 100)
	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)
		if err := tree.Walk(tosecFolder.Path, tosecFolder.fileTypesWithArchives(), entries); err != nil {
			errCh <- err
		}
	}()

	return entries, errCh
}

// GetFiles returns a slice of TosecFile objects parsed from the file names in the folder
func (t *TosecFolder) GetFiles() ([]TosecFile, error) {
	entries, errCh := t.GetFileTree()
	var fileList []TosecFile

	for entry := range entries {
		if !entry.IsDir {
			tf, err := ParseFileName(entry.Name)
			if err != nil {
				fmt.Println("error parsing file name: " + entry.Name + " Error: " + err.Error())
				continue
			}
			fileList = append(fileList, *tf)
		}
	}

	// Check for errors after processing entries
	// TODO: Check if it's necessary to wait for the error channel
	select {
	case err := <-errCh:
		if err != nil {
			return nil, err
		}
	}

	return fileList, nil
}

// FormatTree returns a channel of formatted text lines for the tree
func (t *TosecFolder) FormatTree() <-chan string {
	lines := make(chan string, 100)

	go func() {
		defer close(lines)

		lines <- fmt.Sprintf("Showing file tree for: %s", t.Path)

		entries, errCh := t.GetFileTree()
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
		}
	}()

	return lines
}

// GetStats returns statistics about the files in the given path
func (t *TosecFolder) GetStats() (Stats, error) {
	stats := Stats{
		TotalFiles:      0,
		DirectoryCounts: make(map[string]int),
	}
	stats.DirectoryCounts["/"] = 0 // Initialize root directory count

	entries, errCh := t.GetFileTree()
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
	}

	return stats, nil
}

func (tf *TosecFile) extractRestPartOfName() string {
	publisherStr := fmt.Sprintf("(%s)", tf.Publisher)
	idx := strings.LastIndex(tf.FileName, publisherStr)
	rest := tf.FileName[idx+len(publisherStr) : len(tf.FileName)-len(tf.Format)-1]
	return rest
}

func extractValues(elements [][]string) []string {
	values := make([]string, 0)
	if elements != nil {
		for _, val := range elements {
			values = append(values, strings.TrimSpace(val[1]))
		}
	}
	return values
}

func (t *TosecFolder) fileTypesWithArchives() []string {
	if len(t.FileTypes) == 0 {
		return t.FileTypes
	}
	// Add common archive formats to the file types
	// TODO: Consider making this configurable or extensible
	archiveTypes := []string{"zip", "rar", "7z", "tar", "gz", "bz2"}
	fileTypes := make([]string, len(t.FileTypes)+len(archiveTypes))
	copy(fileTypes, t.FileTypes)
	copy(fileTypes[len(t.FileTypes):], archiveTypes)
	return fileTypes
}
