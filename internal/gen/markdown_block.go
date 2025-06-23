package gen

import (
	"fmt"
	"strings"
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
	if len(lines) > 1 && strings.HasPrefix(lines[1], "```") && strings.HasSuffix(lines[len(lines)-1], "```") {
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
