package main

import (
	"fmt"
	"os"

	"github.com/yanmoyy/go-markdown-html/internal/gen"
)

const (
	markdownDirPath = "./files/content"
	outputDirPath   = "./files/output"
	templatePath    = "./template.html"
)

func main() {
	fmt.Println("Markdown to HTML")
	basePath := ""
	if len(os.Args) == 2 {
		basePath = os.Args[1]
	}
	fmt.Printf("Enter Markdown directory path: (Default: %s)\n> ", markdownDirPath)
	mdPath := getInput()
	if mdPath == "" {
		mdPath = markdownDirPath
	}
	fmt.Println("path: ", mdPath)
	fmt.Printf("Enter output directory path: (Default: %s)\n> ", outputDirPath)
	outPath := getInput()
	if outPath == "" {
		outPath = outputDirPath
	}
	fmt.Println("path: ", outPath)
	err := gen.GeneratePagesRecursive(mdPath, templatePath, outPath, basePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Done")
}

func getInput() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return ""
	}
	return input
}
