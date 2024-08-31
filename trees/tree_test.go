package trees

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Traversal(t *testing.T) {

	tr := Tree{}
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
		n *Node
	}
	tests := []struct {
		name string
		tr   *Tree
		args args
	}{
		{
			name: "pre traversal",
			tr:   &tr,
			args: args{n: tr.Root},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("*******Level Tree Traversal ********")
			tt.tr.LevelOrderTraversal()
			fmt.Println("\n*******Pre Traversal********")
			tt.tr.PreOrderTraversal(tt.args.n)
			fmt.Println("\n*******Inorder Traversal********")
			tt.tr.InOrderTraversal(tt.args.n)
			fmt.Println("\n********Post Traversal********")
			tt.tr.PostOrderTraversal(tt.args.n)
			fmt.Println("\n********Tree********")
		})
	}
}

func TestTree_Search(t *testing.T) {
	tr := Tree{}
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
		key int
	}
	tests := []struct {
		name string
		tr   *Tree
		args args
		want bool
	}{
		{
			name: "search a node",
			tr:   &tr,
			args: args{key: 13},
			want: true,
		},
		{
			name: "search a node",
			tr:   &tr,
			args: args{key: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.Search(tt.args.key)

			assert.Equal(t, tt.want, got)
		})
	}
}
