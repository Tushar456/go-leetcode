package trees

import (
	"testing"

	"github.com/Tushar456/go-leetcode/trees"
	"github.com/go-playground/assert"
)

func TestHeight(t *testing.T) {
	tr := trees.Tree{}
	tr.InserNode(10)
	tr.InserNode(8)
	tr.InserNode(12)
	tr.InserNode(14)
	tr.InserNode(6)
	tr.InserNode(4)
	tr.InserNode(2)
	tr.InserNode(1)
	tr.InserNode(13)
	tr.InserNode(15)
	type args struct {
		node *trees.Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Get Max Height",
			args: args{node: tr.Root},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Height(tt.args.node)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDiameter(t *testing.T) {
	tr := trees.Tree{}
	tr.InserNode(10)
	tr.InserNode(8)
	tr.InserNode(12)
	tr.InserNode(14)
	tr.InserNode(6)
	tr.InserNode(4)
	tr.InserNode(2)
	tr.InserNode(1)
	tr.InserNode(13)
	tr.InserNode(15)
	type args struct {
		node *trees.Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Get Diameter",
			args: args{node: tr.Root},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Diameter(tt.args.node)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsBalanced(t *testing.T) {
	tr := trees.Tree{}
	tr.InserNode(10)
	tr.InserNode(8)
	tr.InserNode(12)
	tr.InserNode(14)
	tr.InserNode(6)
	tr.InserNode(4)
	tr.InserNode(2)
	tr.InserNode(1)
	tr.InserNode(13)
	tr.InserNode(15)

	tr2 := trees.Tree{}
	tr2.InserNode(10)
	tr2.InserNode(8)
	tr2.InserNode(13)
	tr2.InserNode(12)
	tr2.InserNode(14)
	tr2.InserNode(6)
	tr2.InserNode(9)
	type args struct {
		node *trees.Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Is balanced",
			args: args{node: tr.Root},
			want: false,
		},
		{
			name: "Is balanced",
			args: args{node: tr2.Root},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsBalanced(tt.args.node)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsIdentical(t *testing.T) {
	tr := trees.Tree{}
	tr.InserNode(10)
	tr.InserNode(8)
	tr.InserNode(12)
	tr.InserNode(14)
	tr.InserNode(6)
	tr.InserNode(4)
	tr.InserNode(2)
	tr.InserNode(1)
	tr.InserNode(13)
	tr.InserNode(15)

	tr2 := trees.Tree{}
	tr2.InserNode(10)
	tr2.InserNode(8)
	tr2.InserNode(12)
	tr2.InserNode(14)
	tr2.InserNode(6)
	tr2.InserNode(4)
	tr2.InserNode(2)
	tr2.InserNode(1)
	tr2.InserNode(13)
	tr2.InserNode(15)

	tr3 := trees.Tree{}
	tr3.InserNode(20)
	tr3.InserNode(8)
	tr3.InserNode(12)
	tr3.InserNode(14)
	tr3.InserNode(6)
	tr3.InserNode(4)
	tr3.InserNode(2)
	tr3.InserNode(1)
	tr3.InserNode(13)
	tr3.InserNode(15)

	tr4 := trees.Tree{}
	tr4.InserNode(20)
	tr4.InserNode(8)
	tr4.InserNode(12)
	tr4.InserNode(14)
	tr4.InserNode(6)
	tr4.InserNode(2)
	tr4.InserNode(1)
	tr4.InserNode(13)
	tr4.InserNode(15)

	type args struct {
		node1 *trees.Node
		node2 *trees.Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Is Identical",
			args: args{node1: tr.Root, node2: tr2.Root},
			want: true,
		},
		{
			name: "Is Identical",
			args: args{node1: tr.Root, node2: tr3.Root},
			want: false,
		},
		{
			name: "Is Identical",
			args: args{node1: tr3.Root, node2: tr4.Root},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := IsIdentical(tt.args.node1, tt.args.node2)
			assert.Equal(t, tt.want, got)

		})
	}
}

func TestIsSumTree(t *testing.T) {

	tr := trees.Tree{}
	tr.InserNode(10)
	tr.InserNode(-8)
	tr.InserNode(18)
	type args struct {
		node *trees.Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Is Identical",
			args: args{node: tr.Root},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSumTree(tt.args.node)
			assert.Equal(t, tt.want, got)
		})
	}
}
