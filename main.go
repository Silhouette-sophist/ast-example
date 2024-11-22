package main

import (
	"ast-example/visitor"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	// 解析 Go 源文件
	fset := token.NewFileSet()
	node, err := parser.ParseDir(fset, "/Users/silhouette/work-practice/gin-example", nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个 FunctionVisitor 实例
	visitor := &visitor.FuncVisitor{}
	// 遍历每个包中的文件
	for _, pkg := range node {
		for _, file := range pkg.Files {
			fmt.Printf("file: %s\n", file.Name)
			ast.Walk(visitor, file)
		}
	}
}
