package main

import (
	"ast-example/dto"
	"ast-example/service"
	"fmt"
	"strings"
)

func main() {
	// 1.获取所有依赖包及其路径
	rootDir := "/Users/silhouette/work-practice/gin-example"
	relatedPkgs := []string{"gin-example", "github.com/gin-gonic/gin"}
	depsMap, err := service.QueryProjectDeps(rootDir)
	if err != nil {
		return
	}
	// 2.对所有关联依赖包收集所有函数
	infos := make([]*dto.FunctionInfo, 0)
	for dir, pkg := range depsMap {
		for _, relatedPkg := range relatedPkgs {
			if strings.Contains(relatedPkg, pkg) {
				fmt.Printf("deps dir %s pkg %s\n", dir, pkg)
				methods, err := service.TransversePkgMethods(rootDir, dir)
				if err != nil {
					return
				}
				for _, method := range methods {
					infos = append(infos, method)
				}
			}
		}
	}
	service.WriteJsonToFile("method.json", infos)
}
