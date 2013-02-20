// Package deepimports provides helpers to recursively find packages used by
// other packages.
package deepimports

import (
	"fmt"
	"go/build"
)

type recursiveImports struct {
	paths    []string
	srcDir   string
	done     map[string]bool
	imported []*build.Package
	all      []*build.Package
}

func (r *recursiveImports) FindImportsOnly() ([]*build.Package, error) {
	if err := r.analyzeAll(); err != nil {
		return nil, err
	}
	return r.imported, nil
}

func (r *recursiveImports) Find() ([]*build.Package, error) {
	if err := r.analyzeAll(); err != nil {
		return nil, err
	}
	return r.all, nil
}

func (r *recursiveImports) analyzeAll() error {
	r.done = make(map[string]bool)
	for _, pkg := range r.paths {
		if err := r.analyzePackage(pkg, true); err != nil {
			return err
		}
	}
	return nil
}

func (r *recursiveImports) analyzePackage(importPath string, given bool) error {
	if importPath == "C" {
		return nil
	}
	if r.done[importPath] {
		return nil
	}
	r.done[importPath] = true

	pkg, err := build.Import(importPath, r.srcDir, build.AllowBinary)
	if err != nil {
		return fmt.Errorf(
			"Failed to find import path %s with error %s", importPath, err)
	}

	if !given {
		r.imported = append(r.imported, pkg)
	}
	r.all = append(r.all, pkg)

	for _, importPath := range pkg.Imports {
		if err := r.analyzePackage(importPath, false); err != nil {
			return err
		}
	}
	return nil
}

// Finds imports and includes input packages as well.
func Find(paths []string, srcDir string) ([]*build.Package, error) {
	r := &recursiveImports{
		paths:  paths,
		srcDir: srcDir,
	}
	return r.Find()
}

// Finds imports only, does not include self. If multiple input packages are
// specified and they import each other, they will be included.
func FindImportsOnly(paths []string, srcDir string) ([]*build.Package, error) {
	r := &recursiveImports{
		paths:  paths,
		srcDir: srcDir,
	}
	return r.FindImportsOnly()
}
