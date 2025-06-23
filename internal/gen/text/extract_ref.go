package text

import (
	"fmt"
	"regexp"
)

type refType int

const (
	refLink refType = iota
	refImage
)

type ref struct {
	refType refType
	text    string
	url     string
	start   int
	end     int
}

func extractMarkdownRefs(text string) ([]ref, error) {
	// link: [alt text](url)
	// image: ![alt text](url)
	pattern := `\[([^\[\]]*)\]\(([^\(\)]*)\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("error (compile regex): ")
	}
	matches := re.FindAllStringSubmatchIndex(text, -1)
	res := []ref{}
	// append to result (link or image)
	for _, match := range matches {
		// match: [st(pattern), en(pattern), st(alt), en(alt), st(url), en(url)]
		alt := text[match[2]:match[3]]
		url := text[match[4]:match[5]]
		if match[0] == 0 {
			res = append(res, ref{refType: refLink, text: alt, url: url, start: match[0], end: match[1]})
			continue
		}
		if text[match[0]-1:match[0]] == "!" {
			res = append(res, ref{refType: refImage, text: alt, url: url, start: match[0] - 1, end: match[1]})
		} else {
			res = append(res, ref{refType: refLink, text: alt, url: url, start: match[0], end: match[1]})
		}
	}
	return res, nil
}
