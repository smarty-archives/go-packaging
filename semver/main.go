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
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if result, err := parse(readInput(), readUpstream()); err == nil {
		fmt.Println(result)
	} else {
		os.Exit(1)
	}
}

func readInput() string {
	input, _ := ioutil.ReadAll(os.Stdin)
	return string(input)
}
func readUpstream() string {
	args := os.Args[1:]
	if len(args) > 0 {
		return strings.TrimSpace(args[0])
	} else {
		return ""
	}
}

func parse(input, upstream string) (interface{}, error) {
	if len(upstream) > 0 {
		return parseUpstreamVersion(input, upstream)
	} else {
		return parseNativeVersion(input)
	}
}

func parseUpstreamVersion(input, upstream string) (interface{}, error) {
	if parsed, err := ParseUpstream(upstream, input); err == nil {
		return parsed, nil
	} else {
		return "", err
	}
}

func parseNativeVersion(input string) (interface{}, error) {
	if parsed, err := ParseNative(input); err == nil {
		return parsed.Increment(), nil
	} else {
		return "", err
	}
}
