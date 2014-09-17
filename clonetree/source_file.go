package main

import "path"

type SourceFile struct {
	Filename   string
	ImportPath string
}

func (this SourceFile) Destination(target string) string {
	return path.Join(target, "src", this.ImportPath, path.Base(this.Filename))
}
