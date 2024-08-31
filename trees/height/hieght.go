package trees

import (
	"github.com/Tushar456/go-leetcode/trees"
)

func Height(node *trees.Node) int {

	if node == nil {
		return 0
	}

	lHeight := Height(node.Left)
	rHeight := Height(node.Right)

	return max(lHeight, rHeight) + 1

}

func Diameter(node *trees.Node) int {

	diameter, _ := diameter(node)
	return diameter
}

func diameter(node *trees.Node) (int, int) {

	if node == nil {
		return 0, 0
	}

	lDiameter, lHeight := diameter(node.Left)
	rDiameter, rHeight := diameter(node.Right)
	diameter := max(lDiameter, max(rDiameter, lHeight+rHeight+1))

	return diameter, max(lHeight, rHeight) + 1

}

func IsBalanced(node *trees.Node) bool {
	balanced, _ := balanced(node)
	return balanced
}

func balanced(node *trees.Node) (bool, int) {

	if node == nil {
		return true, 0
	}

	lBalanced, lHeight := balanced(node.Left)
	rBalanced, rHeight := balanced(node.Right)
	diff := abs(lHeight - rHeight)

	if lBalanced && rBalanced && diff <= 1 {
		return true, max(lHeight, rHeight) + 1
	} else {
		return false, max(lHeight, rHeight)
	}

}

func IsSumTree(node *trees.Node) bool {
	isSum, _ := sumTree(node)
	return isSum
}

func sumTree(node *trees.Node) (bool, int) {

	if node == nil {
		return true, 0
	}

	if node.Left == nil && node.Right == nil {
		return true, node.Data
	}

	isLeftSum, lSum := sumTree(node.Left)
	isRightSum, rSum := sumTree(node.Right)

	cond := node.Data == lSum+rSum

	if isLeftSum && isRightSum && cond {
		return true, node.Data + lSum + rSum
	} else {
		return false, node.Data + lSum + rSum
	}

}

func IsIdentical(node1, node2 *trees.Node) bool {

	if node1 == nil && node2 == nil {
		return true
	}

	if node1 != nil && node2 == nil {
		return false
	}

	if node1 == nil && node2 != nil {
		return false
	}

	left := IsIdentical(node1.Left, node2.Left)
	right := IsIdentical(node1.Right, node2.Right)

	if node1.Data == node2.Data && left && right {
		return true
	} else {
		return false
	}

}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
