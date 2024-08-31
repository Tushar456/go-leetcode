package fastslowpointer

import (
	"testing"

	"github.com/Tushar456/go-leetcode/linklist"
	"github.com/go-playground/assert"
)

func Test_middleElemenet(t *testing.T) {

	ll := linklist.LinkedList{}
	ll.AddToTail(10)
	ll.AddToTail(20)
	ll.AddToTail(30)
	ll.AddToTail(40)
	ll.AddToTail(50)
	ll.AddToTail(60)
	ll.AddToTail(70)

	ll2 := linklist.LinkedList{}
	ll2.AddToTail(10)
	ll2.AddToTail(20)
	ll2.AddToTail(30)
	ll2.AddToTail(40)
	ll2.AddToTail(50)
	ll2.AddToTail(60)
	ll2.AddToTail(70)
	ll2.AddToTail(80)

	ll3 := linklist.LinkedList{}

	ll4 := linklist.LinkedList{}
	ll4.AddToTail(10)

	ll5 := linklist.LinkedList{}
	ll5.AddToTail(10)
	ll5.AddToTail(20)

	type args struct {
		head *linklist.Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find middle element",
			args: args{head: ll.Head},
			want: 40,
		},
		{
			name: "find middle element",
			args: args{head: ll2.Head},
			want: 50,
		},
		{
			name: "find middle element",
			args: args{head: ll3.Head},
			want: -1,
		},
		{
			name: "find middle element",
			args: args{head: ll4.Head},
			want: -1,
		},
		{
			name: "find middle element",
			args: args{head: ll5.Head},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := middleElemenet(tt.args.head)
			assert.Equal(t, tt.want, got)

		})
	}
}
