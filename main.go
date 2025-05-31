package main

import (
	"fmt"
	"os"

	"tosec-manager/tosec"
)

func printUsage() {
	fmt.Println(`Usage: tosec <command> [<args>]

Available commands:

    show <path>		Show file tree of the specified path
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
		err := tosec.ShowFileTree(path)
		if err != nil {
			fmt.Printf("Error showing file tree: %v\n", err)
			return
		}
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printUsage()
	}
}
