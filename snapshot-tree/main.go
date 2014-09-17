package main

import (
	"fmt"
	"os"
)

func main() {
	workingDirectory, _ := os.Getwd()

	// TODO: flags from command line:
	// --makefile (generate a "forwarding makefile")
	// --dryrun (indicates if we're not actually copying the files)
	// --output=dir (the location of the directory in which to place the files)

	files := BuildTree(workingDirectory)
	dryRun := false             // TODO: from command line
	outputDirectory := "output" // TODO: complain if the target exists? (or write to a temp dir?)

	err := files.Copy(outputDirectory, dryRun)
	if err != nil {
		fmt.Println(err)
	}
}
