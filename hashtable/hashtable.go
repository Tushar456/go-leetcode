package hashtable

import (
	"fmt"
)

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func initHashTable() *HashTable {
	hashTable := &HashTable{}

	for i := range hashTable.array {

		hashTable.array[i] = &bucket{}
	}

	return hashTable

}

func (ht *HashTable) Add(key string) {

	index := hash(key)

	ht.array[index].insert(key)

}

func (ht *HashTable) Search(key string) bool {
	index := hash(key)

	return ht.array[index].search(key)

}

func (ht *HashTable) Delete(key string) {

	index := hash(key)

	ht.array[index].delete(key)

}

func hash(key string) int {
	sumOfCh := 0
	for _, v := range key {

		sumOfCh += int(v)
	}

	return sumOfCh % ArraySize

}

func (b *bucket) insert(key string) {
	if !b.search(key) {

		newNode := &bucketNode{key: key}

		newNode.next = b.head
		b.head = newNode

	} else {
		fmt.Println("key already exist")
	}

}

func (b *bucket) search(key string) bool {

	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}

	return false

}

func (b *bucket) delete(key string) {

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	previouNode := b.head

	for previouNode.next != nil {
		if previouNode.next.key == key {
			previouNode.next = previouNode.next.next
			return
		}
		previouNode = previouNode.next
	}

}
