package main

import (
	"fmt"
	"os"
)

func main() {
	workingDirectory, _ := os.Getwd()

	files := BuildTree(workingDirectory)
	dryRun := false             // TODO: from command line
	outputDirectory := "output" // TODO: complain if the target exists? (or write to a temp dir?)

	err := files.Copy(outputDirectory, dryRun)
	if err != nil {
		fmt.Println(err)
	}
}
