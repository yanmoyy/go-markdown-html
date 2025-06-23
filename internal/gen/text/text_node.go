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

type textNode struct {
	textType textType
	value    string
	url      string
}

func NewTextNode(textType textType, value string, url string) *textNode {
	return &textNode{
		textType: textType,
		value:    value,
		url:      url,
	}
}

func (n *textNode) toHTMLNode() (*html.Node, error) {
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
		return nil, fmt.Errorf("unknown text type: %d", n.textType)
	}
}
