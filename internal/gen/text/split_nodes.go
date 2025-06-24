package text

import (
	"fmt"
	"strings"
)

func splitNodesDelimiter(oldNodes []Node, delimiter string, textType TextType) ([]Node, error) {
	newNodes := []Node{}
	for _, node := range oldNodes {
		if node.textType != TextPlain {
			newNodes = append(newNodes, node)
			continue
		}
		sections := strings.Split(node.value, delimiter)
		if len(sections)%2 == 0 {
			return nil, fmt.Errorf("invalid markdown: delimiter %s is not closed", delimiter)
		}
		for i, sec := range sections {
			if sec == "" {
				continue
			}
			if i%2 == 0 {
				// not inside delimiter
				newNodes = append(newNodes, Node{
					textType: TextPlain,
					value:    sec,
				})
			} else {
				// inside delimiter
				newNodes = append(newNodes, Node{
					textType: textType,
					value:    sec,
				})
			}
		}
	}
	return newNodes, nil
}

func splitNodesRef(oldNodes []Node) ([]Node, error) {
	newNodes := []Node{}
	for _, node := range oldNodes {
		if node.textType != TextPlain {
			newNodes = append(newNodes, node)
			continue
		}
		refs, err := extractMarkdownRefs(node.value)
		if err != nil {
			return nil, err
		}
		cur := 0
		for _, ref := range refs {
			textType := TextLink
			if ref.refType == refImage {
				textType = TextImage
			}
			if ref.start > cur {
				newNodes = append(newNodes, Node{
					textType: TextPlain,
					value:    node.value[cur:ref.start],
				})
			}
			newNodes = append(newNodes, Node{
				textType: textType,
				value:    ref.text,
				url:      ref.url,
			})
			cur = ref.end
		}
		if cur < len(node.value) {
			newNodes = append(newNodes, Node{
				textType: TextPlain,
				value:    node.value[cur:],
			})
		}
	}
	return newNodes, nil
}
