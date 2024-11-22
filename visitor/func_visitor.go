package visitor

import (
	"fmt"
	"go/ast"
	"log/slog"
)

type FuncVisitor struct {
}

func (receiver FuncVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		// 收集函数名称和签名
		funcInfo := fmt.Sprintf("Function: %s, Signature: %s", n.Name.Name, n.Type)
		slog.Info("funInfo: %s", funcInfo)
	}
	return receiver
}
