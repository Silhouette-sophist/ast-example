package dto

// FunctionInfo 结构体，用于存储函数信息
type FunctionInfo struct {
	AFile     string
	RFile     string
	Name      string
	StartLine int
	EndLine   int
	Hash      string
}
