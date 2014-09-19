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

type Tree struct {
	files   map[string]SourceFile
	options *Options
}

func BuildTree(options *Options) *Tree {
	tree := &Tree{
		files:   map[string]SourceFile{},
		options: options,
	}

	for _, directory := range tree.followDirectories(0, options.SourceDirectory) {
		pkg, _ := options.Context.ImportDir(directory, 0)
		tree.appendPackage(pkg)
	}

	return tree
}
func (this *Tree) followDirectories(depth int, root string) []string {
	if strings.HasSuffix(root, ".git") || strings.HasSuffix(root, ".hg") {
		return []string{} // skip source directories
	} else if depth == 1 && strings.HasSuffix(root, path.Join("/", this.options.TargetDirectory)) {
		// note: if they have some qualifying prefix on the clone directory, e.g. ../ that qualifier
		// places the clone directory outside of the scope of our concern
		return []string{} // skip the clone directory
	}

	found := []string{root}
	contents, _ := ioutil.ReadDir(root)
	for _, item := range contents {
		if item.IsDir() {
			directory := path.Join(root, item.Name())
			children := this.followDirectories(depth+1, directory)
			found = append(found, children...)
		}
	}

	return found
}
func (this *Tree) appendPackage(pkg *build.Package) {
	if pkg == nil {
		return // ignore packages that can't be loaded
	} else if pkg.Goroot && pkg.ImportPath != "" && !strings.Contains(pkg.ImportPath, ".") {
		return // ignore standard library package
	}

	this.appendPackageImports(pkg, pkg.Imports)
	this.appendPackageImports(pkg, pkg.TestImports)
	// this.appendPackageImports(pkg, pkg.XTestImports) // TODO: investigate recursion problem

	this.appendPackageFiles(pkg, pkg.GoFiles)
	this.appendPackageFiles(pkg, pkg.IgnoredGoFiles)
	this.appendPackageFiles(pkg, pkg.CgoFiles)
	this.appendPackageFiles(pkg, pkg.CFiles)
	this.appendPackageFiles(pkg, pkg.CXXFiles)
	this.appendPackageFiles(pkg, pkg.MFiles)
	this.appendPackageFiles(pkg, pkg.HFiles)
	this.appendPackageFiles(pkg, pkg.SFiles)
	this.appendPackageFiles(pkg, pkg.SwigFiles)
	this.appendPackageFiles(pkg, pkg.SwigCXXFiles)
	this.appendPackageFiles(pkg, pkg.SysoFiles)
	this.appendPackageFiles(pkg, pkg.TestGoFiles)
	// this.appendPackageFiles(pkg, pkg.XTestGoFiles)
}
func (this *Tree) appendPackageImports(pkg *build.Package, imports []string) {
	for _, importPath := range imports {
		child, _ := this.options.Context.Import(importPath, "", 0)
		this.appendPackage(child)
	}
}
func (this *Tree) appendPackageFiles(pkg *build.Package, items []string) {
	for _, item := range items {
		fullname := path.Join(pkg.Dir, item)
		this.files[fullname] = SourceFile{
			Filename:   fullname,
			ImportPath: pkg.ImportPath,
		}
	}
}

func (this *Tree) CopySource() error {
	targetDir := this.options.TargetDirectory
	if err := this.ensureDirectory(targetDir); err != nil {
		return err
	}

	for _, file := range this.files {
		source := file.Filename
		destination := file.Destination(targetDir)
		if err := this.copyFile(source, destination); err != nil {
			return err
		}
	}
	return nil
}
func (this *Tree) ensureDirectory(directory string) error {
	if this.options.DryRun {
		fmt.Printf("Making directory [%s]\n", directory)
	} else if err := os.MkdirAll(directory, 0777); err != nil {
		return err
	}
	return nil
}
func (this *Tree) copyFile(source, destination string) error {
	if err := this.ensureDirectory(path.Dir(destination)); err != nil {
		return err
	} else if this.options.DryRun {
		fmt.Printf("Copying [%s] to [%s]\n", source, destination)
	} else if err := copyFileContents(source, destination); err != nil {
		return err
	}
	return nil
}
func copyFileContents(source, destination string) error {
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

	_, err = io.Copy(destinationHandle, sourceHandle)
	return err
}

func (this *Tree) GenerateMakefile() error {
	destinationFile := path.Join(this.options.TargetDirectory, "Makefile")
	if this.options.Makefile == "" {
		return nil
	} else if this.options.DryRun {
		fmt.Printf("Generating [Makefile] in [%s]\n", destinationFile)
	} else {
		makefileDir := path.Join("src", this.options.Makefile) // "go-ify" the directory
		contents := fmt.Sprintf(templateMakefile, makefileDir, makefileDir)
		return ioutil.WriteFile(destinationFile, []byte(contents), 0777)
	}
	return nil
}

const templateMakefile = `#!/usr/bin/make -f
%%:
	@export GOPATH="$(PWD)"; cd "%s"; make $@
clean:
	@export GOPATH="$(PWD)"; cd "%s"; make clean
	@rm -rf bin pkg
`
