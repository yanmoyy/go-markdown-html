package gen

import (
	"fmt"
	"strings"

	"github.com/yanmoyy/go-markdown-html/internal/gen/html"
	"github.com/yanmoyy/go-markdown-html/internal/gen/text"
)

func blockToHTML(block string) (html.Node, error) {
	blockType := getBlockType(block)
	switch blockType {
	case blockParagraph:
		return paragraphToHTML(block), nil
	case blockHeader:
		return headerToHTML(block)
	case blockCode:
		return codeToHTML(block)
	case blockQuote:
		return quoteToHTML(block)
	case blockOList:
		return olistToHTML(block)
	case blockUList:
		return ulistToHTML(block)
	}
	return html.Node{}, fmt.Errorf("unknown block type: %s", blockType)
}

func paragraphToHTML(block string) html.Node {
	paragraph := strings.ReplaceAll(block, "\n", " ")
	children := textToChildren(paragraph)
	return html.NewParentNode("p", children, nil)
}

func headerToHTML(block string) (html.Node, error) {
	level := 0
	for _, c := range block {
		if c == '#' {
			level++
		} else {
			break
		}
	}
	if level+1 == len(block) {
		return html.Node{}, fmt.Errorf("invalid header level")
	}
	children := textToChildren(block[level+1:])
	return html.NewParentNode(fmt.Sprintf("h%d", level), children, nil), nil
}

func codeToHTML(block string) (html.Node, error) {
	if !strings.HasPrefix(block, "```\n") || !strings.HasSuffix(block, "\n```") {
		return html.Node{}, fmt.Errorf("invalid code block")
	}
	code := block[4 : len(block)-4]
	rawTextNode := text.NewTextNode(text.TextPlain, code, "")
	node, err := rawTextNode.ToHTMLNode()
	if err != nil {
		return html.Node{}, err
	}
	children := []html.Node{node}
	codeNode := html.NewParentNode("code", children, nil)
	return html.NewParentNode("pre", []html.Node{codeNode}, nil), nil
}

func quoteToHTML(block string) (html.Node, error) {
	lines := strings.Split(block, "\n")
	new_lines := []string{}
	for _, line := range lines {
		if !strings.HasPrefix(line, ">") {
			return html.Node{}, fmt.Errorf("invalid quote block")
		}
		line = strings.TrimSpace(line[1:])
		new_lines = append(new_lines, line)
	}
	children := textToChildren(strings.Join(new_lines, " "))
	return html.NewParentNode("blockquote", children, nil), nil
}

func ulistToHTML(block string) (html.Node, error) {
	lines := strings.Split(block, "\n")
	children := []html.Node{}
	for _, line := range lines {
		if !strings.HasPrefix(line, "- ") {
			return html.Node{}, fmt.Errorf("invalid ulist block")
		}
		line = strings.TrimSpace(line[2:])
		children = append(children, textToListItem(line))
	}
	return html.NewParentNode("ul", children, nil), nil
}

func olistToHTML(block string) (html.Node, error) {
	lines := strings.Split(block, "\n")
	children := []html.Node{}
	for i, line := range lines {
		if !strings.HasPrefix(line, fmt.Sprintf("%d. ", i+1)) {
			return html.Node{}, fmt.Errorf("invalid olist block")
		}
		line = strings.TrimSpace(line[2:])
		children = append(children, textToListItem(line))
	}
	return html.NewParentNode("ol", children, nil), nil
}

func textToListItem(line string) html.Node {
	return html.NewParentNode("li", textToChildren(line), nil)
}

func textToChildren(block string) []html.Node {
	nodes, err := text.TextToTextNodes(block)
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
