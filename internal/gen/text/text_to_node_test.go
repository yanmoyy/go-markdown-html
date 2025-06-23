package text

import (
	"reflect"
	"testing"
)

func TestSplitNodesDelimiter(t *testing.T) {
	type args struct {
		oldNodes  []*Node
		delimiter string
		textType  textType
	}
	tests := []struct {
		name    string
		args    args
		want    []*Node
		wantErr bool
	}{
		{
			name: "delimiter bold",
			args: args{
				oldNodes: []*Node{
					{
						textType: textPlain,
						value:    "This is **bold** text",
					},
				},
				delimiter: delimiterBold,
				textType:  textBold,
			},
			want: []*Node{
				{textType: textPlain, value: "This is "},
				{textType: textBold, value: "bold"},
				{textType: textPlain, value: " text"},
			},
		},
		{
			name: "delimiter italic",
			args: args{
				oldNodes: []*Node{
					{
						textType: textPlain,
						value:    "This is _italic_ text",
					},
				},
				delimiter: delimiterItalic,
				textType:  textItalic,
			},
			want: []*Node{
				{textType: textPlain, value: "This is "},
				{textType: textItalic, value: "italic"},
				{textType: textPlain, value: " text"},
			},
		},
		{
			name: "delimiter code",
			args: args{
				oldNodes: []*Node{
					{
						textType: textPlain,
						value:    "This is `code` text",
					},
				},
				delimiter: delimiterCode,
				textType:  textCode,
			},
			want: []*Node{
				{textType: textPlain, value: "This is "},
				{textType: textCode, value: "code"},
				{textType: textPlain, value: " text"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitNodesDelimiter(tt.args.oldNodes, tt.args.delimiter, tt.args.textType)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitNodesDelimiter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitNodesDelimiter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
