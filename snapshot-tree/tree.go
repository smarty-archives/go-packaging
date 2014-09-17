package main

import (
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var buildContext build.Context

type Tree map[string]SourceFile
type SourceFile struct {
	Filename   string
	ImportPath string
}

func init() {
	buildContext = build.Default
	buildContext.UseAllFiles = true
	buildContext.CgoEnabled = true
}

func BuildTree(directory string) Tree {
	tree := Tree(map[string]SourceFile{})
	for _, dir := range followDirectories(directory) {
		pkg, _ := buildContext.ImportDir(dir, 0)
		tree.appendPackage(pkg)
	}
	return tree
}
func followDirectories(root string) []string {
	found := []string{root}

	if strings.HasSuffix(root, ".git") || strings.HasSuffix(root, ".hg") {
		return []string{} // skip source directories
	}

	contents, _ := ioutil.ReadDir(root)
	for _, item := range contents {
		if item.IsDir() {
			directory := path.Join(root, item.Name())
			children := followDirectories(directory)
			found = append(found, children...)
		}
	}

	return found
}
func (this Tree) appendPackage(pkg *build.Package) {
	if pkg == nil {
		return // ignore packages that can't be loaded
	} else if pkg.Goroot && pkg.ImportPath != "" && !strings.Contains(pkg.ImportPath, ".") {
		return // ignore standard library package
	}

	this.appendPackageImports(pkg, pkg.Imports)
	// this.appendPackageImports(pkg, pkg.TestImports)
	// this.appendPackageImports(pkg, pkg.XTestImports) // causes a recursion problem

	this.appendPackageFiles(pkg, pkg.GoFiles)
	// this.appendPackageFiles(pkg, pkg.IgnoredGoFiles)
	// this.appendPackageFiles(pkg, pkg.CgoFiles)
	// this.appendPackageFiles(pkg, pkg.CFiles)
	// this.appendPackageFiles(pkg, pkg.CXXFiles)
	// this.appendPackageFiles(pkg, pkg.MFiles)
	// this.appendPackageFiles(pkg, pkg.HFiles)
	// this.appendPackageFiles(pkg, pkg.SFiles)
	// this.appendPackageFiles(pkg, pkg.SwigFiles)
	// this.appendPackageFiles(pkg, pkg.SwigCXXFiles)
	// this.appendPackageFiles(pkg, pkg.SysoFiles)
	// this.appendPackageFiles(pkg, pkg.TestGoFiles)
	// this.appendPackageFiles(pkg, pkg.XTestGoFiles)
}
func (this Tree) appendPackageImports(pkg *build.Package, imports []string) {
	for _, importPath := range imports {
		child, _ := buildContext.Import(importPath, "", 0)
		this.appendPackage(child)
	}
}
func (this Tree) appendPackageFiles(pkg *build.Package, items []string) {
	for _, item := range items {
		fullname := path.Join(pkg.Dir, item)
		this[fullname] = SourceFile{
			Filename:   fullname,
			ImportPath: pkg.ImportPath,
		}
	}
}

func (this Tree) Copy(outputDirectory string, dryRun bool) error {
	ensureDirectory(outputDirectory, dryRun)

	for _, file := range this {
		source := file.Filename
		destination := path.Join(outputDirectory, "src", file.ImportPath, path.Base(file.Filename))
		if err := copyFile(source, destination, dryRun); err != nil {
			return err
		}
	}
	return nil
}
func ensureDirectory(directory string, dryRun bool) error {
	if dryRun {
		fmt.Printf("Making directory [%s]\n", directory)
	} else if err := os.MkdirAll(directory, 0x777); err != nil {
		return err
	}
	return nil
}
func copyFile(source, destination string, dryRun bool) error {
	if err := ensureDirectory(path.Dir(destination), dryRun); err != nil {
		return err
	} else if dryRun {
		fmt.Printf("Copying [%s] to [%s]\n", source, destination)
		return nil
	}

	sourceHandle, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceHandle.Close()

	destinationHandle, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationHandle.Close()

	_, err = io.Copy(sourceHandle, destinationHandle)
	return err
}
