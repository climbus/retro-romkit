package main

import (
	"fmt"
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
		if len(os.Args) < 3 {
			fmt.Println("Error: 'show' command requires a path argument.")
			return
		}
		path := os.Args[2]

		lines := tosec.FormatTree(path)
		for line := range lines {
			fmt.Println(line)
		}
	case "stats":
		if len(os.Args) < 3 {
			fmt.Println("Error: 'stats' command requires a path argument.")
			return
		}
		path := os.Args[2]
		stats, err := tosec.GetStats(path)

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
