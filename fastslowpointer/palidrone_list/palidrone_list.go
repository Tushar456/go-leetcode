package fastslowpointer

import (
	"github.com/Tushar456/go-leetcode/linklist"
)

func IsPalidrone(head *linklist.Node) bool {

	slow := head
	fast := head

	// 12321
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	first := head
	second := reverse(slow)

	for second != nil {
		if first.Data == second.Data {

			first = first.Next
			second = second.Next

		} else {
			return false
		}

	}
	return true
}

func reverse(head *linklist.Node) *linklist.Node {

	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	var pre *linklist.Node

	for curr != nil { //1->2->3

		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}

	return pre

}
