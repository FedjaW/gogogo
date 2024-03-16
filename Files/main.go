package main

// import block
import (
	"fmt"
	"os"

	"leckomio.dev/go/files/data"
	"leckomio.dev/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := "/data/text.txt"
	result, err := fileutils.ReadTextFile(rootPath + filePath)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", result, result, result)
	writeError := fileutils.WriteToFile(rootPath+"/data/out.txt", newContent)

	fmt.Println(writeError)
	fmt.Println(result)

	data.Test()
}
