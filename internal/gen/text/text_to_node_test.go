package text

import (
	"reflect"
	"testing"
)

func TestTextToTextNodes(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    []Node
		wantErr bool
	}{
		{
			name: "all types",
			args: args{
				text: "plain text **bold** _italic_ `code` [link](https://example.com) and ![image](https://example.com/image.png)",
			},
			want: []Node{
				{textType: textPlain, value: "plain text "},
				{textType: textBold, value: "bold"},
				{textType: textPlain, value: " "},
				{textType: textItalic, value: "italic"},
				{textType: textPlain, value: " "},
				{textType: textCode, value: "code"},
				{textType: textPlain, value: " "},
				{textType: textLink, value: "link", url: "https://example.com"},
				{textType: textPlain, value: " and "},
				{textType: textImage, value: "image", url: "https://example.com/image.png"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := textToTextNodes(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
				return
			}
			for i, gotNode := range got {
				if reflect.DeepEqual(gotNode, tt.want[i]) {
					continue
				}
				t.Errorf("got[%d] = %v, want[%d] = %v", i, gotNode, i, tt.want[i])
			}
		})
	}
}
