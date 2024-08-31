package fastslowpointer

import "github.com/Tushar456/go-leetcode/linklist"

func deleteNthNodeFromLast(head *linklist.Node, pos int) *linklist.Node { // 1->2->3->4->5->6->7

	if head == nil && head.Next == nil {
		return nil
	}

	fast := head
	slow := head

	for i := 0; i < pos; i++ {

		fast = fast.Next

	}

	for fast != nil && fast.Next != nil {

		fast = fast.Next
		slow = slow.Next

	}

	delete := slow.Next
	slow.Next = delete.Next
	delete = nil

	return head

}
