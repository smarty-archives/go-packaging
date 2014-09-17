package main

import "fmt"

func main() {
	options := ParseOptions()
	tree := BuildTree(options)

	if err := tree.CopySource(); err != nil {
		fmt.Println(err)
	} else if err := tree.GenerateMakefile(); err != nil {
		fmt.Println(err)
	}
}
