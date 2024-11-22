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

