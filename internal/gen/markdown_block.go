package gen

import (
	"fmt"
	"strings"

	"github.com/yanmoyy/go-markdown-html/internal/gen/html"
)

type blockType int

const (
	blockParagraph blockType = iota
	blockHeader
	blockCode
	blockQuote
	blockOList // Ordered List
	blockUList // Unordered List
)

func (t blockType) String() string {
	switch t {
	case blockParagraph:
		return "paragraph"
	case blockHeader:
		return "header"
	case blockCode:
		return "code"
	case blockQuote:
		return "quote"
	case blockOList:
		return "olist"
	case blockUList:
		return "ulist"
	default:
		return "unknown"
	}
}

func markdownToHTMLNode(markdown string) (html.Node, error) {
	blocks := markdownToBlocks(markdown)
	children := []html.Node{}
	for _, block := range blocks {
		node, err := blockToHTML(block)
		if err != nil {
			return html.Node{}, err
		}
		children = append(children, node)
	}
	return html.NewParentNode("div", children, nil), nil
}

func markdownToBlocks(markdown string) []string {
	blocks := strings.Split(markdown, "\n\n")
	filtered := []string{}
	for _, block := range blocks {
		if block == "" {
			continue
		}
		block = strings.TrimSpace(block)
		filtered = append(filtered, block)
	}
	return filtered
}

func getBlockType(block string) blockType {
	lines := strings.Split(block, "\n")
	headerPrefix := []string{"# ", "## ", "### ", "#### ", "##### ", "###### ", "###### "}
	for _, prefix := range headerPrefix {
		if strings.HasPrefix(block, prefix) {
			return blockHeader
		}
	}
	if len(lines) > 1 && strings.HasPrefix(lines[0], "```") && strings.HasSuffix(lines[len(lines)-1], "```") {
		return blockCode
	}
	if strings.HasPrefix(block, ">") {
		for _, line := range lines {
			if !strings.HasPrefix(line, ">") {
				return blockParagraph
			}
		}
		return blockQuote
	}
	if strings.HasPrefix(block, "- ") {
		for _, line := range lines {
			if !strings.HasPrefix(line, "- ") {
				return blockParagraph
			}
		}
		return blockUList
	}
	if strings.HasPrefix(block, "1. ") {
		for i, line := range lines {
			if !strings.HasPrefix(line, fmt.Sprintf("%d. ", i+1)) {
				return blockParagraph
			}
		}
		return blockOList
	}
	return blockParagraph
}
