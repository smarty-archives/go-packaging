package main

import (
	"flag"
	"go/build"
	"os"
)

type Options struct {
	Makefile        bool // copy root makefile + generate makefile
	Debian          bool // copy root debian to root of generated source tree
	DryRun          bool
	Context         build.Context
	TargetDirectory string
	SourceDirectory string
}

func ParseOptions() *Options {
	flag.Parse()

	context := build.Default
	context.UseAllFiles = true
	context.CgoEnabled = true
	workingDirectory, _ := os.Getwd()

	// should we simply have an "include"?, e.g. --include=debian --include=Makefile?
	// unfortunately both of those need special treatment

	// TODO
	return &Options{
		Makefile:        true,
		Debian:          true,
		DryRun:          false,
		Context:         context,
		SourceDirectory: workingDirectory,
		TargetDirectory: "output",
	}
}

const templateMakefile = ""
