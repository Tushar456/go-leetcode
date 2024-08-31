package fastslowpointer

import (
	"testing"

	"github.com/Tushar456/go-leetcode/linklist"
	"github.com/go-playground/assert"
)

func Test_detectCycle(t *testing.T) {
	ll := linklist.LinkedList{}
	ll.AddToTail(10)
	ll.AddToTail(20)
	ll.AddToTail(30)
	ll.AddToTail(40)
	ll.AddToTail(50)
	ll.AddToTail(60)
	ll.AddToTail(70)

	ll.Tail.Next = ll.Head.Next.Next

	ll2 := linklist.LinkedList{}
	ll2.AddToTail(10)
	ll2.AddToTail(20)
	ll2.AddToTail(30)
	ll2.AddToTail(40)
	ll2.AddToTail(50)

	ll3 := linklist.LinkedList{}

	ll4 := linklist.LinkedList{}
	ll4.AddToTail(10)

	ll5 := linklist.LinkedList{}
	ll5.AddToTail(10)

	ll5.Tail.Next = ll5.Head

	ll6 := linklist.LinkedList{}
	ll6.AddToTail(10)
	ll6.AddToTail(20)
	ll6.AddToTail(30)
	ll6.AddToTail(40)
	ll6.AddToTail(50)
	ll6.Tail.Next = ll6.Head

	type args struct {
		head *linklist.Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "detect cycle in linklist",
			args: args{head: ll.Head},
			want: true,
		},
		{
			name: "detect cycle in linklist not found",
			args: args{head: ll2.Head},
			want: false,
		},
		{
			name: "detect cycle in linklist not found",
			args: args{head: ll3.Head},
			want: false,
		},
		{
			name: "detect cycle in linklist not found",
			args: args{head: ll4.Head},
			want: false,
		},
		{
			name: "detect cycle in linklist not found",
			args: args{head: ll5.Head},
			want: true,
		},
		{
			name: "detect cycle in linklist not found",
			args: args{head: ll6.Head},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectCycle(tt.args.head)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_detectCycleNode(t *testing.T) {
	ll := linklist.LinkedList{}
	ll.AddToTail(10)
	ll.AddToTail(20)
	ll.AddToTail(30)
	ll.AddToTail(40)
	ll.AddToTail(50)
	ll.AddToTail(60)
	ll.AddToTail(70)

	ll.Tail.Next = ll.Head.Next.Next

	ll2 := linklist.LinkedList{}
	ll2.AddToTail(10)
	ll2.AddToTail(20)
	ll2.AddToTail(30)
	ll2.AddToTail(40)
	ll2.AddToTail(50)

	ll3 := linklist.LinkedList{}

	ll4 := linklist.LinkedList{}
	ll4.AddToTail(10)

	ll5 := linklist.LinkedList{}
	ll5.AddToTail(10)

	ll5.Tail.Next = ll5.Head

	ll6 := linklist.LinkedList{}
	ll6.AddToTail(10)
	ll6.AddToTail(20)
	ll6.AddToTail(30)
	ll6.AddToTail(40)
	ll6.AddToTail(50)
	ll6.Tail.Next = ll6.Head

	type args struct {
		head *linklist.Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "detect cycle in linklist",
			args: args{head: ll.Head},
			want: 30,
		},
		{
			name: "detect cycle in linklist not found",
			args: args{head: ll2.Head},
			want: -1,
		},
		{
			name: "detect cycle in linklist not found",
			args: args{head: ll3.Head},
			want: -1,
		},
		{
			name: "detect cycle in linklist not found",
			args: args{head: ll4.Head},
			want: -1,
		},
		{
			name: "detect cycle in linklist not found",
			args: args{head: ll5.Head},
			want: 10,
		},
		{
			name: "detect cycle in linklist not found",
			args: args{head: ll6.Head},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectCycleNode(tt.args.head)
			assert.Equal(t, tt.want, got)

		})
	}
}
