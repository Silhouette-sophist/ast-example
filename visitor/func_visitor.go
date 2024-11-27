package visitor

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"strings"
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

// FunctionInfo 结构体，用于存储函数信息
type FunctionInfo struct {
	AFile     string
	RFile     string
	Name      string
	StartLine int
	EndLine   int
	Hash      string
}

// FunctionVisitor 结构体，用于收集函数信息
type FunctionVisitor struct {
	RootDir   string
	Pkg       string
	RFile     string
	AFile     string
	Fset      *token.FileSet
	Functions []*FunctionInfo
}

// Visit 方法用于访问 AST 节点
func (v *FunctionVisitor) Visit(node ast.Node) ast.Visitor {
	if fn, ok := node.(*ast.FuncDecl); ok {
		// 提取函数名、起始行号和结束行号
		funcName := fn.Name.Name
		startPos := v.Fset.Position(fn.Pos())
		endPos := v.Fset.Position(fn.End())

		// 生成函数体的哈希
		hash := generateHash(fn)

		// 将函数信息添加到列表
		v.Functions = append(v.Functions, &FunctionInfo{
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

// generateHash 生成函数体的哈希
func generateHash(fn *ast.FuncDecl) string {
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
