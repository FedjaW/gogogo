package main

import (
	"fmt"
	"os"

	"leckomio.dev/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	filename := "/data/text.txt"
	result, err := fileutils.ReadTextFile(rootPath + filename)

	if err != nil {
		fmt.Printf("ERROR: %v" , err)
	}

	fmt.Println(result)
}