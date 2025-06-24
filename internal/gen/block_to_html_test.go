package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanmoyy/go-markdown-html/internal/gen/html"
)

func TestBlockToHTML(t *testing.T) {
	type args struct {
		block string
	}
	tests := []struct {
		name    string
		args    args
		want    html.Node
		wantErr bool
	}{
		{
			name: "paragraph",
			args: args{
				block: "This is a paragraph",
			},
			want: html.NewParentNode("p", []html.Node{
				html.NewLeafNode("p", "This is a paragraph", nil),
			}, nil),
		},
		{
			name: "header",
			args: args{
				block: "# This is a header",
			},
			want: html.NewParentNode("h1", []html.Node{
				html.NewLeafNode("p", "This is a header", nil),
			}, nil),
		},
		{
			name: "code",
			args: args{
				block: "```\nThis is a code block\n```",
			},
			want: html.NewParentNode("pre", []html.Node{
				html.NewParentNode("code", []html.Node{
					html.NewLeafNode("p", "This is a code block", nil),
				}, nil),
			}, nil),
		},
		{
			name: "quote",
			args: args{
				block: "> This is a quote\n> This is another quote",
			},
			want: html.NewParentNode("blockquote", []html.Node{
				html.NewLeafNode("p", "This is a quote This is another quote", nil),
			}, nil),
		},
		{
			name: "ulist",
			args: args{
				block: "- This is a ulist\n- This is another ulist",
			},
			want: html.NewParentNode("ul", []html.Node{
				html.NewParentNode("li", []html.Node{
					html.NewLeafNode("p", "This is a ulist", nil),
				}, nil),
				html.NewParentNode("li", []html.Node{
					html.NewLeafNode("p", "This is another ulist", nil),
				}, nil),
			}, nil),
		},
		{
			name: "olist",
			args: args{
				block: "1. This is a olist\n2. This is another olist",
			},
			want: html.NewParentNode("ol", []html.Node{
				html.NewParentNode("li", []html.Node{
					html.NewLeafNode("p", "This is a olist", nil),
				}, nil),
				html.NewParentNode("li", []html.Node{
					html.NewLeafNode("p", "This is another olist", nil),
				}, nil),
			}, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := blockToHTML(tt.args.block)
			if (err != nil) != tt.wantErr {
				t.Errorf("blockToHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMarkdownToHTMLNode(t *testing.T) {
	type args struct {
		markdown string
	}
	tests := []struct {
		name    string
		args    args
		want    html.Node
		wantErr bool
	}{
		{
			name: "sample markdown",
			args: args{
				markdown: `
This is **bolded** paragraph

This is another paragraph with _italic_ text and ` + "`code`" + ` here
This is the same paragraph on a new line

- This is a list
- with items
`,
			},
			want: html.NewParentNode("div", []html.Node{
				html.NewParentNode("p", []html.Node{
					html.NewLeafNode("p", "This is ", nil),
					html.NewLeafNode("b", "bolded", nil),
					html.NewLeafNode("p", " paragraph", nil),
				}, nil),
				html.NewParentNode("p", []html.Node{
					html.NewLeafNode("p", "This is another paragraph with ", nil),
					html.NewLeafNode("i", "italic", nil),
					html.NewLeafNode("p", " text and ", nil),
					html.NewLeafNode("code", "code", nil),
					html.NewLeafNode("p", " here This is the same paragraph on a new line", nil),
				}, nil),
				html.NewParentNode("ul", []html.Node{
					html.NewParentNode("li", []html.Node{
						html.NewLeafNode("p", "This is a list", nil),
					}, nil),
					html.NewParentNode("li", []html.Node{
						html.NewLeafNode("p", "with items", nil),
					}, nil),
				}, nil),
			}, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := markdownToHTMLNode(tt.args.markdown)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
