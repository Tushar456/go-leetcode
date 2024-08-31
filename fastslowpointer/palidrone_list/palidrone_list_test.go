package fastslowpointer

import (
	"testing"

	"github.com/Tushar456/go-leetcode/linklist"
	"github.com/go-playground/assert"
)

func TestIsPalidrone(t *testing.T) {
	ll := linklist.LinkedList{}
	ll.AddToTail(1)
	ll.AddToTail(2)
	ll.AddToTail(3)
	ll.AddToTail(2)
	ll.AddToTail(1)

	ll2 := linklist.LinkedList{}
	ll2.AddToTail(1)
	ll2.AddToTail(2)
	ll2.AddToTail(2)
	ll2.AddToTail(1)

	ll3 := linklist.LinkedList{}
	ll3.AddToTail(1)
	ll3.AddToTail(2)
	ll3.AddToTail(1)
	ll3.AddToTail(1)

	type args struct {
		head *linklist.Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "find middle element",
			args: args{head: ll.Head},
			want: true,
		},
		{
			name: "find middle element",
			args: args{head: ll2.Head},
			want: true,
		},
		{
			name: "find middle element",
			args: args{head: ll3.Head},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalidrone(tt.args.head)
			assert.Equal(t, tt.want, got)

		})
	}
}
