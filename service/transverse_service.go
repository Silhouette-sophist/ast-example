package service

import (
	"ast-example/visitor"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func TransversePkgMethods(dir string) ([]*visitor.FunctionInfo, error) {
	// 1.创建fileSet，通过parse解析指定目录
	fset := token.NewFileSet()
	node, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	// 2.创建一个 FunctionVisitor用于采集信息
	visitor := &visitor.FunctionVisitor{
		Fset: fset,
	}
	// 遍历每个包中的文件
	for _, pkg := range node {
		fmt.Printf("TransversePkgMethods pkg %s, files %d\n", pkg.Name, len(pkg.Files))
		for _, file := range pkg.Files {
			fmt.Printf("TransversePkgMethods file: %s\n", file.Name)
			ast.Walk(visitor, file)
		}
	}
	return visitor.Functions, nil
}
