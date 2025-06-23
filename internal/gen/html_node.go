package gen

import "fmt"

type nodeType int
type props map[string]string

// define to parent or leaf node
const (
	typeLeaf nodeType = iota
	typeParent
)

type htmlNode struct {
	nodeType nodeType
	tag      string
	value    string
	props    props
	children []*htmlNode
}

func NewLeafNode(tag, value string, props props) *htmlNode {
	return &htmlNode{
		nodeType: typeLeaf,
		tag:      tag,
		value:    value,
		props:    props,
		children: []*htmlNode{},
	}
}

func NewParentNode(tag string, children []*htmlNode, props props) *htmlNode {
	return &htmlNode{
		nodeType: typeParent,
		tag:      tag,
		props:    props,
		children: children,
	}
}

func (n *htmlNode) toHTML() (string, error) {
	switch n.nodeType {
	case typeLeaf:
		if n.value == "" {
			return "", fmt.Errorf("invalid HTML: value is empty")
		}
		if n.tag == "" {
			return n.value, nil
		}
		return fmt.Sprintf("<%s>%s</%s>", n.tag, n.value, n.tag), nil
	case typeParent:
		if n.tag == "" {
			return "", fmt.Errorf("invalid HTML: tag is empty")
		}
		if len(n.children) == 0 {
			return "", fmt.Errorf("invalid HTML: children is empty")
		}
		html := fmt.Sprintf("<%s%s>", n.tag, n.propsToHTML())
		for _, child := range n.children {
			val, err := child.toHTML()
			if err != nil {
				return "", err
			}
			html += val
		}
		return html + fmt.Sprintf("</%s>", n.tag), nil
	}
	return "", nil
}

func (n *htmlNode) propsToHTML() string {
	props := ""
	for k, v := range n.props {
		props += fmt.Sprintf(" %s=\"%s\"", k, v)
	}
	return props
}
