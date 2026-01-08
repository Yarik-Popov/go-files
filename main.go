package main

import (
	"Yarik-Popov/go-files/src"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Go files")
	config, err := src.ParseArgs()
	if err != nil {
		log.Fatal(err)
	}

	err = src.InitDestinationDir(config.Destination)
	fmt.Printf("src: %s, dst: %s, pattern: %s \n", config.Source, config.Destination, config.Pattern)
	if err != nil {
		log.Fatal(err)
	}

	err = src.OperateOnFiles(config)
	if err != nil {
		log.Fatal(err)
	}
}
