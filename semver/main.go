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
//
// In the second case the patch number will be incremented by default
// unless either the -major flag or -minor flag is set. Setting the
// -minor flag causes the minor (2nd) number to be incremented and
// the patch number to be reset. Likewise, setting the -major flag
// will cause the major (1st) number to be incremented and both the
// minor and patch numbers to be reset.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	minor := flag.Bool("minor", false, "Increment the minor number (instead of the patch number) in case of a commit distance.")
	major := flag.Bool("major", false, "Increment the major number (instead of the patch number) in case of a commit distance.")
	flag.Parse()

	if input, err := ioutil.ReadAll(os.Stdin); err != nil {
		os.Exit(1)
	} else if parsed, err := ParseVersion(string(input)); err != nil {
		os.Exit(1)
	} else if *major {
		fmt.Println(parsed.IncrementMajor())
	} else if *minor {
		fmt.Println(parsed.IncrementMinor())
	} else {
		fmt.Println(parsed.IncrementPatch())
	}
}
