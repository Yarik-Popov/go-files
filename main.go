package main

import (
	"Yarik-Popov/go-files/src"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Go files")

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("You must use 2 arguments. src dst")
		os.Exit(1)
	}

	source := args[0]
	destination := args[1]

	fullDst, err := files.InitDestinationDir(destination)
	fmt.Printf("src: %s, dst: %s \n", source, fullDst)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
