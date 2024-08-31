package linklistProblem

import (
	"testing"

	"github.com/Tushar456/go-leetcode/linklist"
	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers(t *testing.T) {

	ll1 := linklist.LinkedList{}
	ll1.AddToTail(2)
	ll1.AddToTail(4)
	ll1.AddToTail(3)

	ll2 := linklist.LinkedList{}
	ll2.AddToTail(5)
	ll2.AddToTail(6)
	ll2.AddToTail(4)

	lle := linklist.LinkedList{}
	lle.AddToTail(7)
	lle.AddToTail(0)
	lle.AddToTail(8)
	type args struct {
		l1 *linklist.Node
		l2 *linklist.Node
	}
	tests := []struct {
		name string
		args args
		want *linklist.Node
	}{
		{
			name: "add two linklist",
			args: args{l1: ll1.Head, l2: ll2.Head},
			want: lle.Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)

			for tt.want != nil && got != nil {

				if tt.want.Data != got.Data {
					assert.Fail(t, "linklist value not equal")
				}
				tt.want = tt.want.Next
				got = got.Next
			}

		})
	}
}
