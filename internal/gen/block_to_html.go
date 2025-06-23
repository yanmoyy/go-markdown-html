package gen

import (
	"fmt"
	"strings"

	"github.com/yanmoyy/go-markdown-html/internal/gen/html"
	"github.com/yanmoyy/go-markdown-html/internal/gen/text"
)

func blockToHTML(block string) html.Node {
	blockType := getBlockType(block)
	switch blockType {
	case blockParagraph:
	case blockHeader:
	case blockCode:
	case blockQuote:
	case blockOList:
	case blockUList:
	}
	return html.Node{}
}

func paragraphToHTML(content string) html.Node {
	paragraph := strings.ReplaceAll(content, "\n", " ")
	children := textToChildren(paragraph)
	return html.NewParentNode("p", children, nil)
}

func headerToHTML(content string) (html.Node, error) {
	level := 0
	for _, c := range content {
		if c == '#' {
			level++
		} else {
			break
		}
	}
	if level+1 == len(content) {
		return html.Node{}, fmt.Errorf("invalid header level")
	}
	children := textToChildren(content[level+1:])
	return html.NewParentNode(fmt.Sprintf("h%d", level+1), children, nil), nil
}

func textToChildren(content string) []html.Node {
	nodes, err := text.TextToTextNodes(content)
	if err != nil {
		return nil
	}
	res := []html.Node{}
	for _, node := range nodes {
		htmlNode, err := node.ToHTMLNode()
		if err != nil {
			return nil
		}
		res = append(res, htmlNode)
	}
	return res
}
