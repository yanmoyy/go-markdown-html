package gen

import (
	"reflect"
	"testing"
)

func TestMarkdownToBlocks(t *testing.T) {
	type args struct {
		markdown string
	}
	tests := []struct {
		name string
		args args
		want []string
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
			want: []string{
				"This is **bolded** paragraph",
				"This is another paragraph with _italic_ text and `code` here\nThis is the same paragraph on a new line",
				"- This is a list\n- with items",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := markdownToBlocks(tt.args.markdown); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestGetBlockType(t *testing.T) {
	type args struct {
		block string
	}
	tests := []struct {
		name string
		args args
		want blockType
	}{
		{
			name: "paragraph",
			args: args{
				block: "This is a paragraph",
			},
			want: blockParagraph,
		},
		{
			name: "header",
			args: args{
				block: "# This is a header",
			},
			want: blockHeader,
		},
		{
			name: "code",
			args: args{
				block: "```\nThis is a code block\n```",
			},
			want: blockCode,
		},
		{
			name: "quote",
			args: args{
				block: "> This is a quote",
			},
			want: blockQuote,
		},
		{
			name: "ulist",
			args: args{
				block: "- This is a ulist",
			},
			want: blockUList,
		},
		{
			name: "olist",
			args: args{
				block: "1. This is a olist",
			},
			want: blockOList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBlockType(tt.args.block); got != tt.want {
				t.Errorf("got = %+v, want = %+v", got, tt.want)
			}
		})
	}
}
