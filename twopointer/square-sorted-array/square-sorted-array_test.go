package twopointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_squareSortedArray(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "move zeros to right",
			args: args{numbers: []int{-4, -1, 0, 3, 10}},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "move zeros to right",
			args: args{numbers: []int{7, -3, 2, 3, 11}},
			want: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := squareSortedArray(tt.args.numbers)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
