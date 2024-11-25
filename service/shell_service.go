package service

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// ExecGoCommandWithDir 在指定目录下执行go命令
func ExecGoCommandWithDir(execDir string, args ...string) (string, error) {
	command := exec.Command("go", args...)
	command.Dir = execDir
	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("ExecGoCommandWithDir error %v output %s\n", err, string(output))
		return "", err
	}
	s := strings.TrimSuffix(string(output), "\n")
	return s, nil
}

// QueryProjectDeps 查询项目依赖的包及其路径
func QueryProjectDeps(rootDir string) (map[string]string, error) {
	outputStr, err := ExecGoCommandWithDir(rootDir, "list", "-e", "-test=false", "-deps=true", "-f", "{{.Dir}};{{.ImportPath}}")
	if err != nil {
		return nil, err
	}
	split := strings.Split(outputStr, "\n")
	if len(split) == 0 {
		return nil, errors.New("deps empty")
	}
	depMap := make(map[string]string)
	for _, dep := range split {
		pair := strings.Split(dep, ";")
		if len(pair) == 2 {
			dir := pair[0]
			pkg := pair[1]
			if dir == "" || pkg == "" {
				fmt.Printf("invalid dir %s pkg %s\n", dir, pkg)
				continue
			}
			depMap[dir] = pkg
		}
	}
	return depMap, nil
}
