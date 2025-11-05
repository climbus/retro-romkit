package main

import (
	"fmt"
	"os"
	"github.com/climbus/retro-romkit/pkg/tosec"
)

func main() {
	tf, err := tosec.ParseFileName(os.Args[1])
	if err != nil {
		fmt.Println("Error parsing file name:", err)
		return
	}
	fmt.Println("Parsed Tosec File: ", *tf)
}
