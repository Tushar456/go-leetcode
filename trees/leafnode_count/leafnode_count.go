package trees

import "github.com/Tushar456/go-leetcode/trees"

func GetLeafNodeCount(n *trees.Node, count *int) {

	if n == nil {
		return
	}

	GetLeafNodeCount(n.Left, count)
	if n.Left == nil && n.Right == nil {
		*count = *count + 1
	}
	GetLeafNodeCount(n.Right, count)

}
