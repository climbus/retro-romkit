package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"maps"
	"os"
	"slices"

	"tosec-manager/pkg/tosec"
)

func printUsage() {
	fmt.Println(`Usage: tosec <command> [<args>]

Available commands:

    show <path>		Show file tree of the specified path
    stats <path>	Show statistics about files in the specified path
    help		Show this help message`)
}

func main() {
	if len(os.Args) == 1 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "show":
		path := getPath()

		platform := flag.StringP("platform", "p", "", "Platform to filter by (optional)")
		flag.Parse()

		tosecFolder := tosec.Create(path, *platform)

		lines := tosecFolder.FormatTree()
		for line := range lines {
			fmt.Println(line)
		}
	case "stats":
		path := getPath()

		platform := flag.StringP("platform", "p", "", "Platform to filter by (optional)")
		flag.Parse()

		tosecFolder := tosec.Create(path, *platform)

		stats, err := tosecFolder.GetStats()

		if err != nil {
			fmt.Printf("Error retrieving stats: %v\n", err)
			return
		}
		fmt.Printf("Total files: %d\n", stats.TotalFiles)
		for _, key := range slices.Sorted(maps.Keys(stats.DirectoryCounts)) {
			fmt.Printf("%s (%d)\n", key, stats.DirectoryCounts[key])
		}
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
	return path
}

type Options struct {
	Platform string
}
