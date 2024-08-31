package trees

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) InserNode(data int) {
	newNode := &Node{Data: data}
	if t.Root == nil {
		t.Root = newNode
	}

	insertNode(t.Root, newNode)
}

func insertNode(root *Node, node *Node) {

	if node.Data < root.Data {

		if root.Left == nil {
			root.Left = node
		} else {
			insertNode(root.Left, node)
		}

	} else if node.Data > root.Data {
		if root.Right == nil {
			root.Right = node
		} else {
			insertNode(root.Right, node)
		}

	}

}

func (t *Tree) Search(key int) bool {

	return search(t.Root, key)
}

func search(node *Node, key int) bool {

	if node == nil {
		return false
	}

	if node.Data == key {
		return true
	}

	if key < node.Data {
		return search(node.Left, key)

	} else {
		return search(node.Right, key)
	}

}

func (t *Tree) PreOrderTraversal(n *Node) {

	if n == nil {
		return
	}

	fmt.Printf("%d ", n.Data)
	t.PreOrderTraversal(n.Left)
	t.PreOrderTraversal(n.Right)
}

func (t *Tree) InOrderTraversal(n *Node) {

	if n == nil {
		return
	}

	t.InOrderTraversal(n.Left)
	fmt.Printf("%d ", n.Data)
	t.InOrderTraversal(n.Right)
}

func (t *Tree) PostOrderTraversal(n *Node) {

	if n == nil {
		return
	}

	t.PostOrderTraversal(n.Left)
	t.PostOrderTraversal(n.Right)
	fmt.Printf("%d ", n.Data)
}

func (t *Tree) LevelOrderTraversal() {
	if t.Root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(t.Root)

	for queue.Len() > 0 {
		levelSize := queue.Len()

		// Process nodes at the current level
		for i := 0; i < levelSize; i++ {
			currentNode := queue.Front().Value.(*Node)
			queue.Remove(queue.Front())

			// Print the value of the current node
			fmt.Print(currentNode.Data, " ")

			// Enqueue the left and right children (if any)
			if currentNode.Left != nil {
				queue.PushBack(currentNode.Left)
			}
			if currentNode.Right != nil {
				queue.PushBack(currentNode.Right)
			}
		}

		// Move to the next line after processing all nodes at the current level
		fmt.Println()
	}
}
