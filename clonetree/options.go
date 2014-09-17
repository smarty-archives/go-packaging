package main

import (
	"flag"
	"go/build"
	"os"
)

type Options struct {
	DryRun          bool
	Makefile        string
	TargetDirectory string
	SourceDirectory string
	Context         build.Context
}

func ParseOptions() *Options {
	targetDirectory, makefile, dryRun := "", "", false
	flag.StringVar(&targetDirectory, "target", "clone", "")
	flag.StringVar(&makefile, "makefile", "", "For example: github.com/smartystreets/goconvey")
	flag.BoolVar(&dryRun, "dry-run", false, "")
	flag.Parse()

	context := build.Default
	context.UseAllFiles = true
	context.CgoEnabled = true
	workingDirectory, _ := os.Getwd()

	return &Options{
		Makefile:        makefile,
		DryRun:          dryRun,
		Context:         context,
		SourceDirectory: workingDirectory,
		TargetDirectory: targetDirectory,
	}
}
