FileSet 是 Go 语言 go/token 包中的一个重要数据结构，它用于管理和表示源代码文件及其位置。理解 FileSet 的概念和用途对于处理 Go 代码的解析和分析非常有帮助。以下是对 FileSet 的详细解释。

什么是 FileSet？
FileSet 是一个包含多个文件的集合，每个文件都可以在该集合中被唯一标识。它主要用于：

位置管理：FileSet 提供了一种方式来跟踪源代码中的位置，允许你在代码解析和抽象语法树（AST）处理时获取具体的行号和列号信息。
多文件支持：在处理多个源文件时，FileSet 可以有效地管理所有文件的位置，而无需单独处理每个文件。
主要功能
位置映射：FileSet 可以将源代码中的位置（如行号和列号）映射到具体的文件。这使得在分析或错误提示时，可以准确地指向源代码的具体位置。
文件管理：可以添加多个源文件到 FileSet 中，并对它们进行统一管理。
典型用法
FileSet 通常在以下场景中使用：

- 代码解析：在使用 go/parser 解析 Go 代码时，FileSet 可用于记录每个语法节点的位置信息。
- AST 处理：在遍历和处理抽象语法树时，FileSet 提供位置信息，以便在需要时能够引用原始源代码。
- 错误处理：在报告编译错误或分析错误时，FileSet 可用于提供准确的错误位置，帮助开发者快速定位问题。


```go
type FileSet struct {
	mutex sync.RWMutex         // protects the file set
	base  int                  // base offset for the next file
	files []*File              // list of files in the order added to the set
	last  atomic.Pointer[File] // cache of last file looked up
}
```

### 需求演进
#### 采集函数信息
- 采集所有非匿名函数信息，参数及其类型信息
- 采集匿名函数信息
- 采集函数含有receiver的信息

https://poe.com/s/R6ak4jHDTYI6SZChp6Pk

#### 采集代码块信息
- 采集普通代码块信息
- 采集分支代码块信息，switch、select、if嵌套等
- 采集代码块所属函数

```go
// 按照显示或者隐式括号划分代码块，作为统一的执行单元
func(v *InnerBlockVisitor) Visit(node ast.Node) ast.Visitor {
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
```

#### 代码唯一标识（非直接依赖版本）
- 代码块hash（文件相对路径、函数名、内容）
- 临近hash（前代码块hash+代码块hash+后代码块hash）
