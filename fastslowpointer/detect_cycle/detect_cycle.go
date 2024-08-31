package fastslowpointer

import "github.com/Tushar456/go-leetcode/linklist"

func detectCycle(head *linklist.Node) bool {

	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {

			return true

		}
	}

	return false

}

func detectCycleNode(head *linklist.Node) int {

	if head == nil {
		return -1
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {

			lengthCycle := getLengthCycle(slow)

			index := getCyclicNode(lengthCycle, head)

			return index

		}
	}
	return -1
}

func getCyclicNode(length int, head *linklist.Node) int {
	p2 := head
	p1 := head
	for i := 0; i < length; i++ {
		p2 = p2.Next
	}

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1.Data

}

func getLengthCycle(node *linklist.Node) int {

	temp := node.Next

	size := 1

	for temp != node {

		temp = temp.Next
		size++

	}

	return size

}
