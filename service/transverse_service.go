package service

import (
	"ast-example/visitor"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func TransversePkgMethods(rootDir, pkgDir string) ([]*visitor.FunctionInfo, error) {
	// 1.创建fileSet，通过parse解析指定目录
	fset := token.NewFileSet()
	node, err := parser.ParseDir(fset, pkgDir, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	// 2.遍历每个包中的文件
	functionInfos := make([]*visitor.FunctionInfo, 0)
	for _, pkg := range node {
		fmt.Printf("TransversePkgMethods pkg %s, files %d\n", pkg.Name, len(pkg.Files))
		for _, file := range pkg.Files {
			fmt.Printf("TransversePkgMethods file: %s\n", file.Name)
			v := &visitor.FunctionVisitor{
				RootDir: rootDir,
				Pkg:     pkg.Name,
				File:    file.Name.Name,
				Fset:    fset,
			}
			ast.Walk(v, file)
			functionInfos = append(functionInfos, v.Functions...)
		}
	}
	return functionInfos, nil
}
