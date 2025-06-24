package gen

import (
	"fmt"
	"os"
	"strings"
)

// #nosec: G304
func generatePage(from, to, templatePath, basePath string) error {
	fmt.Println("Generating Page from", from, "to", to, "with", templatePath)
	md, err := os.ReadFile(from)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	node, err := markdownToHTMLNode(string(md))
	if err != nil {
		return fmt.Errorf("failed to convert markdown to html: %w", err)
	}
	html, err := node.ToHTML()
	if err != nil {
		return fmt.Errorf("failed to convert html node to html: %w", err)
	}
	title, err := extractTitle(string(md))
	if err != nil {
		return fmt.Errorf("failed to extract title: %w", err)
	}
	template, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}
	newHTML := strings.Replace(string(template), "{{ Title }}", title, 1)
	newHTML = strings.Replace(newHTML, "{{ Content }}", html, 1)
	newHTML = strings.ReplaceAll(newHTML, `href="/`, `href="`+basePath)
	newHTML = strings.ReplaceAll(newHTML, `src="/`, `src="`+basePath)
	err = os.WriteFile(to, []byte(newHTML), 0600)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func GeneratePagesRecursive(contentDirPath, templatePath, destDirPath, basePath string) error {
	// check if destDirPath exists
	if _, err := os.Stat(destDirPath); os.IsNotExist(err) {
		err = os.MkdirAll(destDirPath, 0750)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	dir, err := os.ReadDir(contentDirPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}
	for _, file := range dir {
		from := contentDirPath + "/" + file.Name()
		htmlFileName := strings.Replace(file.Name(), ".md", ".html", 1)
		to := destDirPath + "/" + htmlFileName
		fmt.Println("*", from, "->", to)
		if file.IsDir() {
			err := GeneratePagesRecursive(from, templatePath, to, basePath)
			if err != nil {
				return err
			}
		} else {
			err := generatePage(from, to, templatePath, basePath)
			if err != nil {
				return fmt.Errorf("failed to generate page: %w", err)
			}
		}
	}
	return nil
}
