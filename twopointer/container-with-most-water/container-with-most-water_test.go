package twopointer

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_findMaxArea(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "move zeros to right",
			args: args{numbers: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxArea(tt.args.numbers)
			assert.Equal(t, tt.want, got)
		})
	}
}
