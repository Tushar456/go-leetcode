package fastslowpointer

import (
	"github.com/Tushar456/go-leetcode/linklist"
)

/*

Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.

*/

func middleElemenet(head *linklist.Node) int {

	if head == nil {
		return -1
	}

	if head.Next == nil || head.Next.Next == nil {
		return -1
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

	}

	return slow.Data

}
