package linklistProblem

import "github.com/Tushar456/go-leetcode/linklist"

func addTwoNumbers(l1 *linklist.Node, l2 *linklist.Node) *linklist.Node {

	ptr, curry := new(linklist.Node), 0
	dummy := ptr
	for l1 != nil || l2 != nil {
		sum := 0 + curry
		if l1 != nil {
			sum += l1.Data
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Data
			l2 = l2.Next
		}

		curry = sum / 10
		sum = sum % 10

		ptr.Next = &linklist.Node{Data: sum}
		ptr = ptr.Next

	}

	if curry > 0 {
		ptr.Next = &linklist.Node{Data: curry}
	}
	return dummy.Next
}
