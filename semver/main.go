package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var currentVersion string
	flag.StringVar(&currentVersion, "current", "", "")
	flag.Parse()

	if parsed, err := ParseVersion(currentVersion); err != nil {
		os.Exit(1)
	} else {
		fmt.Println(parsed.Increment())
	}
}
