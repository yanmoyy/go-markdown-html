package html

import (
	"fmt"
	"testing"
)

func TestPropsToHTML(t *testing.T) {
	type args struct {
		props Props
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				props: Props{},
			},
			want: "",
		},
		{
			name: "single",
			args: args{
				props: Props{
					"class": "foo",
				},
			},
			want: ` class="foo"`,
		},
		{
			name: "multiple",
			args: args{
				props: Props{
					"class": "foo",
					"id":    "bar",
				},
			},
			want: ` class="foo" id="bar"`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d-%s", i, tt.name), func(t *testing.T) {
			if got := propsToHTML(tt.args.props); got != tt.want {
				t.Errorf("propsToHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToHTML(t *testing.T) {
	type args struct {
		node Node
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "leaf",
			args: args{
				node: NewLeafNode("p", "foo", Props{}),
			},
			want: `<p>foo</p>`,
		},
		{
			name: "leaf-empty-value",
			args: args{
				node: NewLeafNode("p", "", Props{}),
			},
			wantErr: true,
		},
		{
			name: "leaf-empty-tag",
			args: args{
				node: NewLeafNode("", "foo", Props{}),
			},
			want: `foo`,
		},
		{
			name: "parent",
			args: args{
				node: NewParentNode("div", []Node{
					NewLeafNode("p", "foo", Props{}),
					NewLeafNode("p", "bar", Props{}),
				}, Props{}),
			},
			want: `<div><p>foo</p><p>bar</p></div>`,
		},
		{
			name: "parent-empty-tag",
			args: args{
				node: NewParentNode("", []Node{
					NewLeafNode("p", "foo", Props{}),
					NewLeafNode("p", "bar", Props{}),
				}, Props{}),
			},
			wantErr: true,
		},
		{
			name: "parent-empty-children",
			args: args{
				node: NewParentNode("div", []Node{}, Props{}),
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d-%s", i, tt.name), func(t *testing.T) {
			got, err := tt.args.node.ToHTML()
			if (err != nil) != tt.wantErr {
				t.Errorf("toHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
