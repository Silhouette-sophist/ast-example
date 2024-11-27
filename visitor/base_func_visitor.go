package visitor

import (
	"ast-example/dto"
	"crypto/sha256"
	"encoding/hex"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"strings"
)

// FunctionVisitor 结构体，用于收集函数信息
type FunctionVisitor struct {
	RootDir   string
	Pkg       string
	RFile     string
	AFile     string
	Fset      *token.FileSet
	Functions []*dto.FunctionInfo
}

// Visit 方法用于访问 AST 节点
func (v *FunctionVisitor) Visit(node ast.Node) ast.Visitor {
	if fn, ok := node.(*ast.FuncDecl); ok {
		// 提取函数名、起始行号和结束行号
		funcName := fn.Name.Name
		startPos := v.Fset.Position(fn.Pos())
		endPos := v.Fset.Position(fn.End())

		// 生成函数体的哈希
		hash := generateBaseFuncHash(fn)

		// 将函数信息添加到列表
		v.Functions = append(v.Functions, &dto.FunctionInfo{
			AFile:     v.AFile,
			RFile:     v.RFile,
			Name:      funcName,
			StartLine: startPos.Line,
			EndLine:   endPos.Line,
			Hash:      hash,
		})
	}
	return v
}

// generateBaseFuncHash 生成函数体的哈希
func generateBaseFuncHash(fn *ast.FuncDecl) string {
	var src strings.Builder
	// 使用 printer.Fprint 将节点转换为字符串
	if err := printer.Fprint(&src, token.NewFileSet(), fn.Body); err != nil {
		log.Fatal(err)
	}
	// 计算哈希
	h := sha256.New()
	h.Write([]byte(src.String()))
	return hex.EncodeToString(h.Sum(nil))
}
