package fset

import (
	"go/token"
	"golang.org/x/tools/go/packages"
)

func NewEmptyFileSet() *token.FileSet {
	return token.NewFileSet()
}

func SSAFileSet(rootDir string) *token.FileSet {
	config := &packages.Config{
		Mode:  packages.LoadAllSyntax,
		Tests: false,
		Dir:   rootDir,
	}
	loadPackages, err := packages.Load(config)
	if err != nil {
		return nil
	}
	if len(loadPackages) == 0 {
		return nil
	}
	return loadPackages[0].Fset
}
