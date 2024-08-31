package arrays

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_pivot_index(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get the pivot index",
			args: args{numbers: []int{1, 7, 3, 6, 5, 6}},
			want: 3,
		},
		{
			name: "get the pivot index",
			args: args{numbers: []int{1, 2, 3}},
			want: -1,
		},
		{
			name: "get the pivot index",
			args: args{numbers: []int{2, 1, -1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivot_index(tt.args.numbers)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_pivotIndexOptimized(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get the pivot index",
			args: args{numbers: []int{1, 7, 3, 6, 5, 6}},
			want: 3,
		},
		{
			name: "get the pivot index",
			args: args{numbers: []int{1, 2, 3}},
			want: -1,
		},
		{
			name: "get the pivot index",
			args: args{numbers: []int{2, 1, -1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotIndexOptimized(tt.args.numbers)
			assert.Equal(t, tt.want, got)
		})
	}
}
