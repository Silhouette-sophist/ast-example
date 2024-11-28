package visitor

import (
	"ast-example/dto"
	"fmt"
	"go/ast"
	"go/token"
)

type CommonBlockVisitor struct {
	RootDir    string
	Pkg        string
	RFile      string
	AFile      string
	Fset       *token.FileSet
	BlockInfos []*dto.BlockInfo
}

func (v *CommonBlockVisitor) Visit(node ast.Node) ast.Visitor {
	switch fn := node.(type) {
	case *ast.BlockStmt:
		startPos := v.Fset.Position(fn.Lbrace)
		endPos := v.Fset.Position(fn.Rbrace)
		fmt.Printf("file %s startPos: %d, endPos: %d\n", v.RFile, startPos.Line, endPos.Line)
		v.BlockInfos = append(v.BlockInfos, &dto.BlockInfo{
			AFile: v.AFile,
			RFile: v.RFile,
			Start: startPos,
			End:   endPos,
		})
	}
	return v
}
