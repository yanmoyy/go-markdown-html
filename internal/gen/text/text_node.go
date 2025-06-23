package text

import (
	"fmt"

	"github.com/yanmoyy/go-markdown-html/internal/gen/html"
)

type textType int

const (
	textPlain textType = iota
	textBold
	textItalic
	textCode
	textLink
	textImage
)

func (t textType) String() string {
	switch t {
	case textPlain:
		return "Plain"
	case textBold:
		return "Bold"
	case textItalic:
		return "Italic"
	case textCode:
		return "Code"
	case textLink:
		return "Link"
	case textImage:
		return "Image"
	default:
		return "Unknown"
	}
}

type Node struct {
	textType textType
	value    string
	url      string
}

func NewTextNode(textType textType, value string, url string) *Node {
	return &Node{
		textType: textType,
		value:    value,
		url:      url,
	}
}

func (n *Node) String() string {
	if n.url == "" {
		return fmt.Sprintf("Node{type: %v, value: %s}", n.textType, n.value)
	}
	return fmt.Sprintf("Node{type: %v, value: %s, url: %s}", n.textType, n.value, n.url)
}

func (n *Node) ToHTMLNode() (html.Node, error) {
	switch n.textType {
	case textPlain:
		return html.NewLeafNode("p", n.value, nil), nil
	case textBold:
		return html.NewLeafNode("b", n.value, nil), nil
	case textItalic:
		return html.NewLeafNode("i", n.value, nil), nil
	case textCode:
		return html.NewLeafNode("code", n.value, nil), nil
	case textLink:
		return html.NewLeafNode("a", n.value, html.Props{
			"href": n.url,
		}), nil
	case textImage:
		return html.NewLeafNode("img", "", html.Props{
			"src": n.url,
			"alt": n.value,
		}), nil
	default:
		return html.Node{}, fmt.Errorf("unknown text type: %d", n.textType)
	}
}
