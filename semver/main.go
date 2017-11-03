// semver accepts a semantic version number (with optional commit
// distance) as produced by `git describe`. Example inputs:
//
// '1.2.3'
//
// or
//
// '1.2.3-1-ab5def'
//
// The output will either be:
//
// 1. The same version number in the absence of any commit distance.
// 2. An incremented version number.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()

	if input, err := ioutil.ReadAll(os.Stdin); err != nil {
		os.Exit(1)
	} else if parsed, err := ParseNative(string(input)); err != nil {
		os.Exit(1)
	} else {
		fmt.Println(parsed.Increment())
	}
}
