package text

import (
	"fmt"

	"github.com/yanmoyy/go-markdown-html/internal/gen/html"
)

type TextType int

const (
	TextPlain TextType = iota
	TextBold
	TextItalic
	TextCode
	TextLink
	TextImage
)

func (t TextType) String() string {
	switch t {
	case TextPlain:
		return "Plain"
	case TextBold:
		return "Bold"
	case TextItalic:
		return "Italic"
	case TextCode:
		return "Code"
	case TextLink:
		return "Link"
	case TextImage:
		return "Image"
	default:
		return "Unknown"
	}
}

type Node struct {
	textType TextType
	value    string
	url      string
}

func NewTextNode(textType TextType, value string, url string) Node {
	return Node{
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
	case TextPlain:
		return html.NewLeafNode("p", n.value, nil), nil
	case TextBold:
		return html.NewLeafNode("b", n.value, nil), nil
	case TextItalic:
		return html.NewLeafNode("i", n.value, nil), nil
	case TextCode:
		return html.NewLeafNode("code", n.value, nil), nil
	case TextLink:
		return html.NewLeafNode("a", n.value, html.Props{
			"href": n.url,
		}), nil
	case TextImage:
		return html.NewLeafNode("img", "", html.Props{
			"src": n.url,
			"alt": n.value,
		}), nil
	default:
		return html.Node{}, fmt.Errorf("unknown text type: %d", n.textType)
	}
}
