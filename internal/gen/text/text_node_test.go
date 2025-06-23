package text

import "testing"

func TestTextNodeToHTMLNode(t *testing.T) {
	tests := []struct {
		name    string
		node    *Node
		want    string
		wantErr bool
	}{
		{
			name: "plain text",
			node: &Node{
				textType: textPlain,
				value:    "plain text",
			},
			want:    "<p>plain text</p>",
			wantErr: false,
		},
		{
			name: "bold text",
			node: &Node{
				textType: textBold,
				value:    "bold text",
			},
			want:    "<b>bold text</b>",
			wantErr: false,
		},
		{
			name: "italic text",
			node: &Node{
				textType: textItalic,
				value:    "italic text",
			},
			want:    "<i>italic text</i>",
			wantErr: false,
		},
		{
			name: "code text",
			node: &Node{
				textType: textCode,
				value:    "code text",
			},
			want:    "<code>code text</code>",
			wantErr: false,
		},
		{
			name: "link text",
			node: &Node{
				textType: textLink,
				value:    "link text",
				url:      "https://example.com",
			},
			want:    "<a href=\"https://example.com\">link text</a>",
			wantErr: false,
		},
		{
			name: "image text",
			node: &Node{
				textType: textImage,
				value:    "image text",
				url:      "https://example.com/image.png",
			},
			want:    "<img alt=\"image text\" src=\"https://example.com/image.png\">",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.node.toHTMLNode()
			if (err != nil) != tt.wantErr {
				t.Errorf("toHTMLNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			html, err := got.ToHTML()
			if err != nil {
				t.Errorf("ToHTML() error = %v", err)
				return
			}
			if html != tt.want {
				t.Errorf("ToHTML() got = %v, want %v", html, tt.want)
			}
		})
	}
}
