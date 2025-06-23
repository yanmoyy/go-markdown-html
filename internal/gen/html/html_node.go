package html

import (
	"fmt"
	"sort"
)

type nodeType int
type Props map[string]string

// define to parent or leaf node
const (
	typeLeaf nodeType = iota
	typeParent
)

type Node struct {
	nodeType nodeType
	tag      string
	value    string
	props    Props
	children []*Node
}

func NewLeafNode(tag, value string, props Props) *Node {
	return &Node{
		nodeType: typeLeaf,
		tag:      tag,
		value:    value,
		props:    props,
		children: []*Node{},
	}
}

func NewParentNode(tag string, children []*Node, props Props) *Node {
	return &Node{
		nodeType: typeParent,
		tag:      tag,
		props:    props,
		children: children,
	}
}

func (n *Node) ToHTML() (string, error) {
	switch n.nodeType {
	case typeLeaf:
		if n.tag == "" {
			return n.value, nil
		}
		props := propsToHTML(n.props)
		if n.value == "" {
			if n.tag != "img" {
				return "", fmt.Errorf("invalid HTML: only img tag can has empty value")
			}
			return fmt.Sprintf("<%s%s>", n.tag, props), nil
		}
		return fmt.Sprintf("<%s%s>%s</%s>", n.tag, props, n.value, n.tag), nil
	case typeParent:
		if n.tag == "" {
			return "", fmt.Errorf("invalid HTML: tag is empty")
		}
		if len(n.children) == 0 {
			return "", fmt.Errorf("invalid HTML: children is empty")
		}
		html := fmt.Sprintf("<%s%s>", n.tag, propsToHTML(n.props))
		for _, child := range n.children {
			val, err := child.ToHTML()
			if err != nil {
				return "", err
			}
			html += val
		}
		return html + fmt.Sprintf("</%s>", n.tag), nil
	}
	return "", nil
}

func propsToHTML(p Props) string {
	keys := make([]string, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	props := ""
	for _, k := range keys {
		props += fmt.Sprintf(" %s=\"%s\"", k, p[k])
	}
	return props
}
