package linklist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	length int
}

func (ll *LinkedList) AddtoHead(data int) { // 2 3 4
	newMode := &Node{Data: data}

	if ll.length == 0 {
		ll.Tail = newMode
	} else {
		newMode.Next = ll.Head
	}

	ll.Head = newMode
	ll.length++

}

func (ll *LinkedList) AddToTail(data int) {
	newMode := &Node{Data: data}

	if ll.length == 0 {
		ll.Head = newMode
	} else {
		ll.Tail.Next = newMode
	}

	ll.Tail = newMode
	ll.length++
}

func (ll *LinkedList) Insert(data int, pos int) {
	newMode := &Node{Data: data}

	if ll.length <= pos {
		ll.AddToTail(data)
		return
	} else if pos == 0 {
		ll.AddtoHead(data)
		return
	}

	currentNode := ll.Head
	for i := 1; i <= ll.length; i++ {
		if i == pos-1 {
			newMode.Next = currentNode.Next
			currentNode.Next = newMode
			break

		}
		currentNode = currentNode.Next
	}
	ll.length++
}

func (ll *LinkedList) Delete(pos int) {
	if pos > ll.length {
		fmt.Println("cannot delete it check the pos of item")
		return
	}
	if ll.length == 1 {
		ll.Head = nil
		ll.Tail = nil
		ll.length--
		return

	}

	if pos == 1 {
		head := ll.Head
		ll.Head = head.Next
		head = nil
		ll.length--
		return

	}

	currentNode := ll.Head
	for i := 1; i < pos; i++ {
		if i == pos-1 {
			deletenode := currentNode.Next
			currentNode.Next = deletenode.Next
			deletenode = nil

			break
		}
		currentNode = currentNode.Next
	}
	ll.length--
}

// // move fastpointer to pos+1 steps and then move slow and faspointer together until fastpointer pointer is nil
// func (ll *LinkedList) DeleteNthFromLast(pos int) {
// 	if pos > ll.length {
// 		fmt.Println("cannot delete it check the pos of item")
// 		return
// 	}
// 	if ll.length == 1 {
// 		ll.Head = nil
// 		ll.Tail = nil
// 		ll.length--
// 		return

// 	}

// 	if pos == ll.length {
// 		head := ll.Head
// 		ll.Head = head.Next
// 		head = nil
// 		ll.length--
// 		return

// 	}

// 	slowPointer := ll.Head
// 	fastPointer := ll.Head

// 	for i := 1; i <= pos+1; i++ {
// 		fastPointer = fastPointer.Next
// 	}

// 	for fastPointer != nil {
// 		fastPointer = fastPointer.Next

// 		slowPointer = slowPointer.Next
// 	}

// 	slowPointer.Next = slowPointer.Next.Next
// 	ll.length--
// 	if pos == 1 {
// 		ll.Tail = slowPointer
// 	}

// }

// func (ll *LinkedList) SortList() {
// 	if ll.Head == nil || ll.Head.Next == nil {
// 		return
// 	}

// 	current := ll.Head

// 	for current != nil {

// 		index := current.Next

// 		for index != nil {
// 			if current.Data > index.Data {

// 				temp := current.Data
// 				current.Data = index.Data
// 				index.Data = temp
// 			}
// 			index = index.Next
// 		}
// 		current = current.Next

// 	}

// }

// func (ll *LinkedList) IterateList() {

// 	currentNode := ll.Head
// 	for currentNode != nil {
// 		fmt.Printf("%p|%d|%p->", currentNode, currentNode.Data, currentNode.Next)
// 		currentNode = currentNode.Next
// 	}
// }

// func (ll *LinkedList) RemoveDuplicate() {
// 	if ll.Head == nil {
// 		return
// 	}
// 	visited := make(map[int]struct{})
// 	current := ll.Head
// 	pre := ll.Head

// 	for current != nil {

// 		if _, ok := visited[current.Data]; ok {

// 			pre.Next = current.Next

// 		} else {
// 			visited[current.Data] = struct{}{}
// 			pre = current
// 		}

// 		current = current.Next
// 	}
// }

// func (ll *LinkedList) Reverese() {

// 	var preNode *Node
// 	var nextNode *Node

// 	ll.Tail = ll.Head
// 	currentNode := ll.Head

// 	for currentNode != nil {
// 		nextNode, currentNode.Next = currentNode.Next, preNode
// 		preNode, currentNode = currentNode, nextNode

// 	}

// 	ll.Head = preNode

// }

// func (ll *LinkedList) FindMiddleElement() *Node {

// 	if ll.Head == nil {
// 		return nil
// 	}

// 	slowPointer := ll.Head
// 	fastPointer := ll.Head

// 	for fastPointer != nil && fastPointer.Next != nil {

// 		slowPointer = slowPointer.Next
// 		fastPointer = fastPointer.Next.Next
// 	}

// 	return slowPointer
// }

// func (ll *LinkedList) IsPalidrone() bool {

// 	if ll.Head == nil {
// 		return false
// 	}

// 	middle := ll.FindMiddleElement()

// 	secondHalf := reverse(middle.Next)

// 	firstHalf := ll.Head
// 	for secondHalf != nil {
// 		if firstHalf.Data != secondHalf.Data {
// 			return false
// 		}
// 		firstHalf = firstHalf.Next
// 		secondHalf = secondHalf.Next
// 	}

// 	return true

// }

// func reverse(head *Node) *Node {
// 	var prev *Node
// 	current := head

// 	for current != nil {
// 		next := current.Next
// 		current.Next = prev
// 		prev = current
// 		current = next
// 	}

// 	return prev
// }

// func (ll *LinkedList) HasCyscle() bool {

// 	if ll.Head == nil {
// 		return false
// 	}

// 	slowPointer := ll.Head
// 	fastPointer := ll.Head

// 	for fastPointer != nil && fastPointer.Next != nil {
// 		slowPointer = slowPointer.Next
// 		fastPointer = fastPointer.Next.Next

// 		if slowPointer == fastPointer {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (ll *LinkedList) RemoveCyscle() {

// 	if ll.Head == nil {
// 		return
// 	}

// 	slowPointer := ll.Head
// 	fastPointer := ll.Head
// 	var pre *Node

// 	for fastPointer != nil && fastPointer.Next != nil {
// 		pre = slowPointer
// 		slowPointer = slowPointer.Next
// 		fastPointer = fastPointer.Next.Next

// 		if slowPointer == fastPointer {
// 			break
// 		}
// 	}

// 	if slowPointer != fastPointer {
// 		return
// 	}

// 	slowPointer = ll.Head
// 	for slowPointer != fastPointer {
// 		pre = fastPointer
// 		slowPointer = slowPointer.Next
// 		fastPointer = fastPointer.Next

// 	}

// 	pre.Next = nil

// }
