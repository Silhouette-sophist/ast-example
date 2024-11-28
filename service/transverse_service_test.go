package service

import "testing"

// TestTransverseDir 遍历目录获取所有函数
func TestTransverseDir(t *testing.T) {
	// 1.获取所有依赖包及其路径
	rootDir := "/Users/silhouette/work-practice/gin-example"
	// relatedPkgs := []string{"gin-example", "github.com/gin-gonic/gin"}
	relatedPkgs := []string{"gin-example"}
	functionInfos, err := TransverseDir(rootDir, relatedPkgs...)
	if err != nil {
		t.Errorf("TransverseDir err %v", err)
		return
	}
	t.Logf("TransverseDir size %d", len(functionInfos))
}

func TestCommonTransverseDir(t *testing.T) {
	// 1.获取所有依赖包及其路径
	rootDir := "/Users/silhouette/work-practice/gin-example"
	// relatedPkgs := []string{"gin-example", "github.com/gin-gonic/gin"}
	relatedPkgs := []string{"gin-example"}
	functionInfos, err := CommonTransverseDir(rootDir, relatedPkgs...)
	if err != nil {
		t.Errorf("TransverseDir err %v", err)
		return
	}
	t.Logf("TransverseDir size %d", len(functionInfos))
}

func TestCommonCodeBlockTransverseDir(t *testing.T) {
	// 1.获取所有依赖包及其路径
	rootDir := "/Users/silhouette/work-practice/gin-example"
	// relatedPkgs := []string{"gin-example", "github.com/gin-gonic/gin"}
	relatedPkgs := []string{"gin-example"}
	blockInfos, err := CommonCodeBlockTransverseDir(rootDir, relatedPkgs...)
	if err != nil {
		t.Errorf("CommonCodeBlockTransverseDir err %v", err)
		return
	}
	t.Logf("CommonCodeBlockTransverseDir size %d", len(blockInfos))
}
