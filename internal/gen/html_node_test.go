package gen

import (
	"fmt"
	"testing"
)

func TestPropsToHTML(t *testing.T) {
	type args struct {
		props props
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				props: props{},
			},
			want: "",
		},
		{
			name: "single",
			args: args{
				props: props{
					"class": "foo",
				},
			},
			want: ` class="foo"`,
		},
		{
			name: "multiple",
			args: args{
				props: props{
					"class": "foo",
					"id":    "bar",
				},
			},
			want: ` class="foo" id="bar"`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d-%s", i, tt.name), func(t *testing.T) {
			n := &htmlNode{
				props: tt.args.props,
			}
			if got := n.propsToHTML(); got != tt.want {
				t.Errorf("propsToHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToHTML(t *testing.T) {
	type args struct {
		node *htmlNode
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
				node: NewLeafNode("p", "foo", props{}),
			},
			want: `<p>foo</p>`,
		},
		{
			name: "leaf-empty-value",
			args: args{
				node: NewLeafNode("p", "", props{}),
			},
			wantErr: true,
		},
		{
			name: "leaf-empty-tag",
			args: args{
				node: NewLeafNode("", "foo", props{}),
			},
			want: `foo`,
		},
		{
			name: "parent",
			args: args{
				node: NewParentNode("div", []*htmlNode{
					NewLeafNode("p", "foo", props{}),
					NewLeafNode("p", "bar", props{}),
				}, props{}),
			},
			want: `<div><p>foo</p><p>bar</p></div>`,
		},
		{
			name: "parent-empty-tag",
			args: args{
				node: NewParentNode("", []*htmlNode{
					NewLeafNode("p", "foo", props{}),
					NewLeafNode("p", "bar", props{}),
				}, props{}),
			},
			wantErr: true,
		},
		{
			name: "parent-empty-children",
			args: args{
				node: NewParentNode("div", []*htmlNode{}, props{}),
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d-%s", i, tt.name), func(t *testing.T) {
			got, err := tt.args.node.toHTML()
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
