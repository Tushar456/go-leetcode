package fastslowpointer

import (
	"testing"

	"github.com/Tushar456/go-leetcode/linklist"
	"github.com/stretchr/testify/assert"
)

func Test_deleteNthNodeFromLast(t *testing.T) {

	ll := linklist.LinkedList{}
	ll.AddToTail(10)
	ll.AddToTail(20)
	ll.AddToTail(30)
	ll.AddToTail(40)
	ll.AddToTail(50)
	ll.AddToTail(60)
	ll.AddToTail(70)

	type args struct {
		head *linklist.Node
		pos  int
	}
	tests := []struct {
		name string
		args args
		want *linklist.Node
	}{
		{
			name: "delete 4 element",
			args: args{head: ll.Head, pos: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteNthNodeFromLast(tt.args.head, tt.args.pos)

			curr := got
			for curr != nil {

				if curr.Data == 40 {
					assert.Fail(t, "element found")
				}

				curr = curr.Next
			}

		})
	}
}
