package visitor

import (
	"fmt"
	"go/ast"
)

type FuncVisitor struct {
}

func (receiver FuncVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		// 收集函数名称和签名
		funcInfo := fmt.Sprintf("Function: %s, Signature: %T", n.Name.Name, n.Type)
		fmt.Printf("funInfo: %s\n", funcInfo)
	}
	return receiver
}
