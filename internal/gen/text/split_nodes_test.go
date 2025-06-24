package text

import (
	"reflect"
	"testing"
)

func Test_splitNodesDelimiter(t *testing.T) {
	type args struct {
		oldNodes  []Node
		delimiter string
		textType  TextType
	}
	tests := []struct {
		name string
		args args
		want []Node
	}{
		{
			name: "delimiter bold",
			args: args{
				oldNodes: []Node{
					{
						textType: TextPlain,
						value:    "This is **bold** text",
					},
				},
				delimiter: delimiterBold,
				textType:  TextBold,
			},
			want: []Node{
				{textType: TextPlain, value: "This is "},
				{textType: TextBold, value: "bold"},
				{textType: TextPlain, value: " text"},
			},
		},
		{
			name: "delimiter italic",
			args: args{
				oldNodes: []Node{
					{
						textType: TextPlain,
						value:    "This is _italic_ text",
					},
				},
				delimiter: delimiterItalic,
				textType:  TextItalic,
			},
			want: []Node{
				{textType: TextPlain, value: "This is "},
				{textType: TextItalic, value: "italic"},
				{textType: TextPlain, value: " text"},
			},
		},
		{
			name: "delimiter code",
			args: args{
				oldNodes: []Node{
					{
						textType: TextPlain,
						value:    "This is `code` text",
					},
				},
				delimiter: delimiterCode,
				textType:  TextCode,
			},
			want: []Node{
				{textType: TextPlain, value: "This is "},
				{textType: TextCode, value: "code"},
				{textType: TextPlain, value: " text"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitNodesDelimiter(tt.args.oldNodes, tt.args.delimiter, tt.args.textType)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitNodesRef(t *testing.T) {
	type args struct {
		oldNodes []Node
	}
	tests := []struct {
		name string
		args args
		want []Node
	}{
		{
			name: "link",
			args: args{
				oldNodes: []Node{
					{
						textType: TextPlain,
						value:    "This is [link](https://example.com) text",
					},
				},
			},
			want: []Node{
				{textType: TextPlain, value: "This is "},
				{textType: TextLink, value: "link", url: "https://example.com"},
				{textType: TextPlain, value: " text"},
			},
		},
		{
			name: "image",
			args: args{
				oldNodes: []Node{
					{
						textType: TextPlain,
						value:    "This is ![image](https://example.com/image.png) text",
					},
				},
			},
			want: []Node{
				{textType: TextPlain, value: "This is "},
				{textType: TextImage, value: "image", url: "https://example.com/image.png"},
				{textType: TextPlain, value: " text"},
			},
		},
		{
			name: "link and image",
			args: args{
				oldNodes: []Node{
					{
						textType: TextPlain,
						value:    "This is [link](https://example.com) and ![image](https://example.com/image.png) text",
					},
				},
			},
			want: []Node{
				{textType: TextPlain, value: "This is "},
				{textType: TextLink, value: "link", url: "https://example.com"},
				{textType: TextPlain, value: " and "},
				{textType: TextImage, value: "image", url: "https://example.com/image.png"},
				{textType: TextPlain, value: " text"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitNodesRef(tt.args.oldNodes)
			if err != nil {
				t.Errorf("error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v, want = %+v", got, tt.want)
			}
		})
	}
}
