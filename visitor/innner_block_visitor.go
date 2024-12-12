package visitor

import (
	"ast-example/dto"
	"fmt"
	"go/ast"
	"go/token"
)

type InnerBlockVisitor struct {
	RootDir    string
	Pkg        string
	RFile      string
	AFile      string
	Fset       *token.FileSet
	BlockInfos []*dto.BlockInfo
}

func (v *InnerBlockVisitor) Visit(node ast.Node) ast.Visitor {
	switch fn := node.(type) {
	case *ast.BlockStmt:
		if len(fn.List) > 0 {
			for _, stmt := range fn.List {
				switch stmt.(type) {
				case *ast.CaseClause:
					pos := v.Fset.Position(stmt.Pos())
					fmt.Printf("case clause...%v\n", pos)
				case *ast.CommClause:
					pos := v.Fset.Position(stmt.Pos())
					fmt.Printf("common clause...%v\n", pos)
				}
			}
		}
	}
	return v
}
