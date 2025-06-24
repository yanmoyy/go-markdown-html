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
