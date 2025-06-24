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
				{textType: TextPlain, value: "plain text "},
				{textType: TextBold, value: "bold"},
				{textType: TextPlain, value: " "},
				{textType: TextItalic, value: "italic"},
				{textType: TextPlain, value: " "},
				{textType: TextCode, value: "code"},
				{textType: TextPlain, value: " "},
				{textType: TextLink, value: "link", url: "https://example.com"},
				{textType: TextPlain, value: " and "},
				{textType: TextImage, value: "image", url: "https://example.com/image.png"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TextToTextNodes(tt.args.text)
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
