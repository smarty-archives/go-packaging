package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	workingDirectory, _ := os.Getwd()
	packages := collectRootPackages(build.Default, []*build.Package{}, workingDirectory)

	for _, item := range packages {
		fmt.Println(item.ImportPath)
	}
}
func collectRootPackages(context build.Context, packages []*build.Package, directory string) []*build.Package {
	if pkg, err := context.ImportDir(directory, 0); err == nil {
		packages = aggregatePackages(context, packages, pkg)
	}

	files, _ := ioutil.ReadDir(directory)
	for _, file := range files {
		if file.IsDir() {
			packages = collectRootPackages(context, packages, path.Join(directory, file.Name()))
		}
	}

	return packages
}
func aggregatePackages(context build.Context, packages []*build.Package, pkg *build.Package) []*build.Package {
	packages = append(packages, pkg)
	packages = aggregateImports(context, packages, pkg.Imports)
	packages = aggregateImports(context, packages, pkg.TestImports)
	packages = aggregateImports(context, packages, pkg.XTestImports)
	return packages
}
func aggregateImports(context build.Context, packages []*build.Package, imports []string) []*build.Package {
	for _, path := range imports {
		packages = aggregateImport(context, packages, path)
	}
	return packages
}
func aggregateImport(context build.Context, packages []*build.Package, importPath string) []*build.Package {
	if contains(packages, importPath) {
		return packages // already contains package
	} else if pkg, err := context.Import(importPath, "", 0); err != nil {
		return packages // couldn't import package
	} else if pkg.Goroot && pkg.ImportPath != "" && !strings.Contains(pkg.ImportPath, ".") {
		return packages // standard library
	} else {
		return aggregatePackages(context, packages, pkg)
	}
}
func contains(packages []*build.Package, path string) bool {
	for _, pkg := range packages {
		if pkg.ImportPath == path {
			return true
		}
	}
	return false
}
