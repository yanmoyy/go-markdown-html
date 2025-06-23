package text

import (
	"reflect"
	"testing"
)

func Test_splitNodesDelimiter(t *testing.T) {
	type args struct {
		oldNodes  []Node
		delimiter string
		textType  textType
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
						textType: textPlain,
						value:    "This is **bold** text",
					},
				},
				delimiter: delimiterBold,
				textType:  textBold,
			},
			want: []Node{
				{textType: textPlain, value: "This is "},
				{textType: textBold, value: "bold"},
				{textType: textPlain, value: " text"},
			},
		},
		{
			name: "delimiter italic",
			args: args{
				oldNodes: []Node{
					{
						textType: textPlain,
						value:    "This is _italic_ text",
					},
				},
				delimiter: delimiterItalic,
				textType:  textItalic,
			},
			want: []Node{
				{textType: textPlain, value: "This is "},
				{textType: textItalic, value: "italic"},
				{textType: textPlain, value: " text"},
			},
		},
		{
			name: "delimiter code",
			args: args{
				oldNodes: []Node{
					{
						textType: textPlain,
						value:    "This is `code` text",
					},
				},
				delimiter: delimiterCode,
				textType:  textCode,
			},
			want: []Node{
				{textType: textPlain, value: "This is "},
				{textType: textCode, value: "code"},
				{textType: textPlain, value: " text"},
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
						textType: textPlain,
						value:    "This is [link](https://example.com) text",
					},
				},
			},
			want: []Node{
				{textType: textPlain, value: "This is "},
				{textType: textLink, value: "link", url: "https://example.com"},
				{textType: textPlain, value: " text"},
			},
		},
		{
			name: "image",
			args: args{
				oldNodes: []Node{
					{
						textType: textPlain,
						value:    "This is ![image](https://example.com/image.png) text",
					},
				},
			},
			want: []Node{
				{textType: textPlain, value: "This is "},
				{textType: textImage, value: "image", url: "https://example.com/image.png"},
				{textType: textPlain, value: " text"},
			},
		},
		{
			name: "link and image",
			args: args{
				oldNodes: []Node{
					{
						textType: textPlain,
						value:    "This is [link](https://example.com) and ![image](https://example.com/image.png) text",
					},
				},
			},
			want: []Node{
				{textType: textPlain, value: "This is "},
				{textType: textLink, value: "link", url: "https://example.com"},
				{textType: textPlain, value: " and "},
				{textType: textImage, value: "image", url: "https://example.com/image.png"},
				{textType: textPlain, value: " text"},
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
