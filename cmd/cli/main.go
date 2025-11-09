package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"maps"
	"os"
	"slices"

	"github.com/climbus/retro-romkit/pkg/tosec"
)

func printUsage() {
	fmt.Println(`Usage: tosec <command> [<args>]

Available commands:

	show <path>		Show file tree of the specified path
	stats <path>		Show statistics about files in the specified path
	list <path>		List all files in the specified path
	copy <path>		Copy files from the specified path to the output directory
	help			Show this help message`)
}

func parsePlatformFlag() string {
	platform := flag.StringP("platform", "p", "", "Platform to filter by (optional)")
	flag.Parse()
	return *platform
}

func main() {
	if len(os.Args) == 1 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "show":
		path := getPath()
		platform := parsePlatformFlag()

		tosecFolder := tosec.Create(path, platform)

		lines := tosecFolder.FormatTree()
		for line := range lines {
			fmt.Println(line)
		}
	case "stats":
		path := getPath()
		platform := parsePlatformFlag()

		tosecFolder := tosec.Create(path, platform)

		stats, err := tosecFolder.GetStats()

		if err != nil {
			fmt.Printf("Error retrieving stats: %v\n", err)
			return
		}
		fmt.Printf("Total files: %d\n", stats.TotalFiles)
		for _, key := range slices.Sorted(maps.Keys(stats.DirectoryCounts)) {
			fmt.Printf("%s (%d)\n", key, stats.DirectoryCounts[key])
		}
	case "list":
		path := getPath()
		tosecFolder := tosec.Create(path, "")

		files, err := tosecFolder.GetFiles()
		if err != nil {
			fmt.Printf("Error retrieving files: %v\n", err)
			return
		}
		for _, file := range files {
			fmt.Printf("%s (%s) - %s - r:%s l:%s : %s\n", file.Title, file.Date, file.Publisher, file.Region, file.Language, file.FileName)
		}
	case "copy":
		path := getPath()
		platform := parsePlatformFlag()
		tosecFolder := tosec.Create(path, platform)

		// outputDir := flag.StringP("output", "o", "", "Output directory to copy files to")
		limit := flag.IntP("limit", "l", 0, "Limit the number of files per directory")
		unzip := flag.BoolP("unzip", "u", false, "Unzip files before copying")
		flag.Parse()

		tosecFolder.BuildTree(tosec.CopyOptions{Limit: *limit, Unzip: *unzip})

	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printUsage()
	}
}

func getPath() string {
	if len(os.Args) < 3 {
		fmt.Println("Error: '" + os.Args[1] + "' command requires a path argument.\n")
		printUsage()
		os.Exit(1)
	}
	path := os.Args[2]

	// Validate that the path exists
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Error: Path '%s' does not exist.\n", path)
		} else {
			fmt.Printf("Error: Cannot access path '%s': %v\n", path, err)
		}
		os.Exit(1)
	}

	// Validate that the path is a directory
	if !info.IsDir() {
		fmt.Printf("Error: Path '%s' is not a directory.\n", path)
		os.Exit(1)
	}

	return path
}
