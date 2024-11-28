package service

import (
	"ast-example/dto"
	"ast-example/visitor"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func TransversePkgMethods(rootDir, pkgDir string) ([]*dto.FunctionInfo, error) {
	// 1.创建fileSet，通过parse解析指定目录
	fset := token.NewFileSet()
	node, err := parser.ParseDir(fset, pkgDir, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	// 2.遍历每个包中的文件
	functionInfos := make([]*dto.FunctionInfo, 0)
	for _, pkg := range node {
		fmt.Printf("TransversePkgMethods pkg %s, files %d\n", pkg.Name, len(pkg.Files))
		for _, file := range pkg.Files {
			fmt.Printf("TransversePkgMethods file: %s\n", file.Name)
			v := &visitor.FunctionVisitor{
				RootDir: rootDir,
				Pkg:     pkg.Name,
				RFile:   file.Name.Name,
				Fset:    fset,
			}
			ast.Walk(v, file)
			functionInfos = append(functionInfos, v.Functions...)
		}
	}
	return functionInfos, nil
}

// TransverseDir 遍历目录下指定包路径的函数
func TransverseDir(rootDir string, relatedPkgs ...string) ([]*dto.FunctionInfo, error) {
	depsMap, err := QueryProjectDeps(rootDir)
	if err != nil {
		return nil, err
	}
	infos := make([]*dto.FunctionInfo, 0)
	for dir, pkg := range depsMap {
		for _, relatedPkg := range relatedPkgs {
			if !strings.HasPrefix(relatedPkg, pkg) {
				fmt.Printf("not related pkg %s\n", pkg)
				continue
			}
			files, err := os.ReadDir(dir)
			if err != nil {
				fmt.Printf("read dir%s err %v\n", pkg, err)
				continue
			}
			for _, file := range files {
				if file.IsDir() {
					fmt.Printf("skip dir %v\n", file.Name())
					continue
				}
				absFile := fmt.Sprintf("%s/%s", dir, file.Name())
				rel, err := filepath.Rel(dir, absFile)
				if err != nil {
					fmt.Printf("not root dir file %v\n", err)
					continue
				}
				RFile := rel
				fileSet := token.NewFileSet()
				if contentBytes, err := os.ReadFile(absFile); err == nil {
					if parseFile, err := parser.ParseFile(fileSet, absFile, contentBytes, parser.ParseComments); err == nil {
						v := &visitor.FunctionVisitor{
							RootDir: rootDir,
							Pkg:     pkg,
							RFile:   RFile,
							AFile:   absFile,
							Fset:    fileSet,
						}
						ast.Walk(v, parseFile)
						infos = append(infos, v.Functions...)
					}
				}
			}
		}
	}
	return infos, nil
}

// CommonTransverseDir 遍历目录下指定包路径的函数
func CommonTransverseDir(rootDir string, relatedPkgs ...string) ([]*dto.FunctionInfo, error) {
	depsMap, err := QueryProjectDeps(rootDir)
	if err != nil {
		return nil, err
	}
	infos := make([]*dto.FunctionInfo, 0)
	for dir, pkg := range depsMap {
		for _, relatedPkg := range relatedPkgs {
			fmt.Printf("xxx %s yyy %s\n", relatedPkg, pkg)
			if !strings.HasPrefix(pkg, relatedPkg) {
				fmt.Printf("not related pkg %s\n", pkg)
				continue
			}
			files, err := os.ReadDir(dir)
			if err != nil {
				fmt.Printf("read dir%s err %v\n", pkg, err)
				continue
			}
			for _, file := range files {
				// TODO 注意，如果这里跳过文件夹，那么要求当前包依赖了文件夹中的文件，否则文件夹中的文件就不会遍历出来
				if file.IsDir() {
					fmt.Printf("skip dir %v\n", file.Name())
					continue
				}
				absFile := fmt.Sprintf("%s/%s", dir, file.Name())
				rel, err := filepath.Rel(rootDir, absFile)
				if err != nil {
					fmt.Printf("not root dir file %v\n", err)
					continue
				}
				RFile := rel
				fileSet := token.NewFileSet()
				if contentBytes, err := os.ReadFile(absFile); err == nil {
					if parseFile, err := parser.ParseFile(fileSet, absFile, contentBytes, parser.ParseComments); err == nil {
						v := &visitor.CommonFuncVisitor{
							RootDir: rootDir,
							Pkg:     pkg,
							RFile:   RFile,
							AFile:   absFile,
							Fset:    fileSet,
						}
						ast.Walk(v, parseFile)
						infos = append(infos, v.Functions...)
					}
				}
			}
		}
	}
	return infos, nil
}

// CommonCodeBlockTransverseDir 遍历目录下指定包路径的函数
func CommonCodeBlockTransverseDir(rootDir string, relatedPkgs ...string) ([]*dto.BlockInfo, error) {
	depsMap, err := QueryProjectDeps(rootDir)
	if err != nil {
		return nil, err
	}
	infos := make([]*dto.BlockInfo, 0)
	for dir, pkg := range depsMap {
		for _, relatedPkg := range relatedPkgs {
			fmt.Printf("xxx %s yyy %s\n", relatedPkg, pkg)
			if !strings.HasPrefix(pkg, relatedPkg) {
				fmt.Printf("not related pkg %s\n", pkg)
				continue
			}
			files, err := os.ReadDir(dir)
			if err != nil {
				fmt.Printf("read dir%s err %v\n", pkg, err)
				continue
			}
			for _, file := range files {
				// TODO 注意，如果这里跳过文件夹，那么要求当前包依赖了文件夹中的文件，否则文件夹中的文件就不会遍历出来
				if file.IsDir() {
					fmt.Printf("skip dir %v\n", file.Name())
					continue
				}
				absFile := fmt.Sprintf("%s/%s", dir, file.Name())
				rel, err := filepath.Rel(rootDir, absFile)
				if err != nil {
					fmt.Printf("not root dir file %v\n", err)
					continue
				}
				RFile := rel
				fileSet := token.NewFileSet()
				if contentBytes, err := os.ReadFile(absFile); err == nil {
					if parseFile, err := parser.ParseFile(fileSet, absFile, contentBytes, parser.ParseComments); err == nil {
						v := &visitor.CommonBlockVisitor{
							RootDir: rootDir,
							Pkg:     pkg,
							RFile:   RFile,
							AFile:   absFile,
							Fset:    fileSet,
						}
						ast.Walk(v, parseFile)
						infos = append(infos, v.BlockInfos...)
					}
				}
			}
		}
	}
	return infos, nil
}
