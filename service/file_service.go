package service

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJsonToFile(filePath string, obj any) {
	marshal, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("marshal Err %v", err)
		return
	}
	writeBytesToFile(filePath, marshal)
}

func writeBytesToFile(filePath string, bytes []byte) {
	create, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create file Err %v", err)
		return
	}
	n, err := create.Write(bytes)
	if err != nil {
		fmt.Printf("write file Err %v", err)
		return
	}
	fmt.Printf("write file bytes %d", n)
}
