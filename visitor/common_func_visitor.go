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

type CommonFuncVisitor struct {
	RootDir   string
	Pkg       string
	RFile     string
	AFile     string
	Fset      *token.FileSet
	Functions []*dto.FunctionInfo
}

// Visit https://poe.com/s/R6ak4jHDTYI6SZChp6Pk
func (v *CommonFuncVisitor) Visit(node ast.Node) ast.Visitor {
	switch fn := node.(type) {
	case *ast.FuncDecl:
		// 提取具名函数、起始行号和结束行号
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
	case *ast.FuncLit:
		// 提取匿名函数的信息
		startPos := v.Fset.Position(fn.Pos())
		endPos := v.Fset.Position(fn.End())
		// 生成函数体的哈希
		hash := generateAnonymousFuncHash(fn)
		// 在匿名函数中没有名称，使用特定字符串表示
		v.Functions = append(v.Functions, &dto.FunctionInfo{
			AFile:     v.AFile,
			RFile:     v.RFile,
			Name:      "anonymous function", // 使用占位符
			StartLine: startPos.Line,
			EndLine:   endPos.Line,
			Hash:      hash,
		})
	}
	return v
}

// generateBaseFuncHash 生成函数体的哈希
func generateAnonymousFuncHash(fn *ast.FuncLit) string {
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
