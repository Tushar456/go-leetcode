package trees

import (
	"testing"

	"github.com/Tushar456/go-leetcode/trees"
	"github.com/go-playground/assert"
)

func TestGetLeafNodeCount(t *testing.T) {

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
	count := 0
	type args struct {
		n     *trees.Node
		count *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "pre traversal",
			args: args{n: tr.Root, count: &count},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetLeafNodeCount(tt.args.n, tt.args.count)
			assert.Equal(t, tt.want, *tt.args.count)
		})
	}
}
