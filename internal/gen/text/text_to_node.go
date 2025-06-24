package text

import (
	"fmt"
)

const (
	// delimiters
	delimiterBold   = "**"
	delimiterItalic = "_"
	delimiterCode   = "`"
)

func TextToTextNodes(text string) ([]Node, error) {
	nodes := []Node{
		{textType: TextPlain, value: text},
	}
	nodes, err := splitNodesDelimiter(nodes, delimiterBold, TextBold)
	if err != nil {
		return nil, fmt.Errorf("error (bold): ")
	}
	nodes, err = splitNodesDelimiter(nodes, delimiterItalic, TextItalic)
	if err != nil {
		return nil, fmt.Errorf("error (italic): ")
	}
	nodes, err = splitNodesDelimiter(nodes, delimiterCode, TextCode)
	if err != nil {
		return nil, fmt.Errorf("error (code): ")
	}
	nodes, err = splitNodesRef(nodes)
	if err != nil {
		return nil, fmt.Errorf("error (ref): ")
	}
	return nodes, nil
}
