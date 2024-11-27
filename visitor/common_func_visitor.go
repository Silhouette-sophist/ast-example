package visitor

import (
	"ast-example/dto"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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
		// 提取返回值信息
		var returnTypes []string
		if fn.Type.Results != nil {
			for _, result := range fn.Type.Results.List {
				returnTypes = append(returnTypes, formatType(result.Type))
			}
		}
		// 提取参数信息
		var params []dto.ParamInfo
		if fn.Type.Params != nil {
			for _, param := range fn.Type.Params.List {
				paramType := formatType(param.Type)
				for _, name := range param.Names {
					params = append(params, dto.ParamInfo{
						Name: name.Name,
						Type: paramType,
					})
				}
			}
		}
		// 将函数信息添加到列表
		v.Functions = append(v.Functions, &dto.FunctionInfo{
			AFile:       v.AFile,
			RFile:       v.RFile,
			Name:        funcName,
			StartLine:   startPos.Line,
			EndLine:     endPos.Line,
			Hash:        hash,
			ReturnTypes: returnTypes,
			Params:      params,
		})
	case *ast.FuncLit:
		// 提取匿名函数的信息
		startPos := v.Fset.Position(fn.Pos())
		endPos := v.Fset.Position(fn.End())
		// 生成函数体的哈希
		hash := generateAnonymousFuncHash(fn)
		// 提取返回值信息
		var returnTypes []string
		if fn.Type.Results != nil {
			for _, result := range fn.Type.Results.List {
				returnTypes = append(returnTypes, formatType(result.Type))
			}
		}
		// 提取参数信息
		var params []dto.ParamInfo
		if fn.Type.Params != nil {
			for _, param := range fn.Type.Params.List {
				paramType := formatType(param.Type)
				for _, name := range param.Names {
					params = append(params, dto.ParamInfo{
						Name: name.Name,
						Type: paramType,
					})
				}
			}
		}
		// 在匿名函数中没有名称，使用特定字符串表示
		v.Functions = append(v.Functions, &dto.FunctionInfo{
			AFile:       v.AFile,
			RFile:       v.RFile,
			Name:        "anonymous function", // 使用占位符
			StartLine:   startPos.Line,
			EndLine:     endPos.Line,
			Hash:        hash,
			ReturnTypes: returnTypes,
			Params:      params,
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

// formatType 用于格式化返回值类型
func formatType(t ast.Expr) string {
	// 这里可以根据类型表达式进行格式化
	return fmt.Sprintf("%s", t) // 你可以根据需要返回更详细的类型信息
}
