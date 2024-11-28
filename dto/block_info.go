package dto

import "go/token"

// BlockInfo 结构体，用于存储代码块信息
type BlockInfo struct {
	AFile string
	RFile string
	Hash  string
	Start token.Position
	End   token.Position
}
