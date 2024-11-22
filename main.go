package main

import (
	"ast-example/visitor"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	// 解析 Go 源文件
	fset := token.NewFileSet()
	node, err := parser.ParseDir(fset, "./", nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个 FunctionVisitor 实例
	visitor := &visitor.FuncVisitor{}
	// 遍历每个包中的文件
	for _, pkg := range node {
		for _, file := range pkg.Files {
			ast.Walk(visitor, file)
		}
	}
}
