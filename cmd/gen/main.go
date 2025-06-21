package main

import "fmt"

func main() {
	fmt.Println("Markdown to HTML")
	fmt.Println("Enter Markdown directory path: (Default: ./markdown)")
	var path string
	_, err := fmt.Scanln(&path)
	if err.Error() == "unexpected newline" {
		path = "./markdown"
	} else if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Entered path: ", path)
}
