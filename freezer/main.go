package main

import (
	"fmt"
	"go/build"
	"os"
)

func main() {
	workingDirectory, _ := os.Getwd()
	fmt.Println("Working Directory:", workingDirectory)

	imported, err := build.ImportDir(workingDirectory, 0)
	if err != nil {
		panic(err)
	}

	for _, file := range imported.GoFiles {
		fmt.Println("Source File:", file)
	}
}
