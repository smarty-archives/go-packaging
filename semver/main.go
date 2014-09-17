package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if input, err := ioutil.ReadAll(os.Stdin); err != nil {
		os.Exit(1)
	} else if parsed, err := ParseVersion(string(input)); err != nil {
		os.Exit(1)
	} else {
		fmt.Println(parsed.Increment())
	}
}
