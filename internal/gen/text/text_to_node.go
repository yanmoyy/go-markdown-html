package text

import (
	"fmt"
	"strings"
)

const (
	// delimiters
	delimiterBold   = "**"
	delimiterItalic = "_"
	delimiterCode   = "`"
)

func splitNodesDelimiter(oldNodes []*Node, delimiter string, textType textType) ([]*Node, error) {
	newNodes := []*Node{}
	for _, node := range oldNodes {
		if node.textType != textPlain {
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
				newNodes = append(newNodes, &Node{
					textType: textPlain,
					value:    sec,
				})
			} else {
				// inside delimiter
				newNodes = append(newNodes, &Node{
					textType: textType,
					value:    sec,
				})
			}
		}
	}
	return newNodes, nil
}
