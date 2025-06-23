package text

import (
	"reflect"
	"testing"
)

func Test_extractMarkdownRefs(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []ref
	}{
		{
			name: "just link",
			args: args{
				text: "[alt text](url)",
			},
			want: []ref{
				{
					refType: refLink,
					text:    "alt text",
					url:     "url",
					start:   0,
					end:     15,
				},
			},
		},
		{
			name: "just image",
			args: args{
				text: "![alt text](url)",
			},
			want: []ref{
				{
					refType: refImage,
					text:    "alt text",
					url:     "url",
					start:   0,
					end:     16,
				},
			},
		},
		{
			name: "multiple links",
			args: args{
				text: "This is text with a [rick roll](https://i.imgur.com/aKaOqIh.gif) and [obi wan](https://i.imgur.com/fJRm4Vk.jpeg)",
			},
			want: []ref{
				{
					refType: refLink,
					text:    "rick roll",
					url:     "https://i.imgur.com/aKaOqIh.gif",
					start:   20,
					end:     64,
				},
				{
					refType: refLink,
					text:    "obi wan",
					url:     "https://i.imgur.com/fJRm4Vk.jpeg",
					start:   69,
					end:     112,
				},
			},
		},
		{
			name: "multiple images",
			args: args{
				text: "This is text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif) and ![obi wan](https://i.imgur.com/fJRm4Vk.jpeg)",
			},
			want: []ref{
				{
					refType: refImage,
					text:    "rick roll",
					url:     "https://i.imgur.com/aKaOqIh.gif",
					start:   20,
					end:     65,
				},
				{
					refType: refImage,
					text:    "obi wan",
					url:     "https://i.imgur.com/fJRm4Vk.jpeg",
					start:   70,
					end:     114,
				},
			},
		},
		{
			name: "combine link and image",
			args: args{
				text: "This is text with a [rick roll](https://i.imgur.com/aKaOqIh.gif) and ![obi wan](https://i.imgur.com/fJRm4Vk.jpeg)",
			},
			want: []ref{
				{
					refType: refLink,
					text:    "rick roll",
					url:     "https://i.imgur.com/aKaOqIh.gif",
					start:   20,
					end:     64,
				},
				{
					refType: refImage,
					text:    "obi wan",
					url:     "https://i.imgur.com/fJRm4Vk.jpeg",
					start:   69,
					end:     113,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractMarkdownRefs(tt.args.text)
			if err != nil {
				t.Errorf("error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
