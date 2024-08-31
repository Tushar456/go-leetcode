package twopointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumUnorderedListBruteForce(t *testing.T) {
	type args struct {
		input  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find target in sorted list",
			args: args{input: []int{2, 5, 7, 11, 15, 19}, target: 27},
			want: []int{1, 2, 4},
		},
		{
			name: "find target in unsorted list",
			args: args{input: []int{3, 2, 6, 8, 7, 23}, target: 21},
			want: []int{2, 3, 4},
		},
		{
			name: "if not present return -1,-1",
			args: args{input: []int{3, 3, 3}, target: 7},
			want: []int{-1, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSumUnorderedListBruteForce(tt.args.input, tt.args.target)
			assert.ElementsMatch(t, tt.want, got)

		})
	}
}

func TestThreeSumOrderedList(t *testing.T) {
	type args struct {
		input  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find target in sorted list",
			args: args{input: []int{2, 5, 7, 11, 15, 19}, target: 27},
			want: []int{1, 2, 4},
		},
		{
			name: "find target in unsorted list",
			args: args{input: []int{1, 2, 6, 8, 7, 23}, target: 21},
			want: []int{2, 3, 4},
		},
		{
			name: "if not present return -1,-1",
			args: args{input: []int{3, 3, 3}, target: 7},
			want: []int{-1, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSumOrderedList(tt.args.input, tt.args.target)
			assert.ElementsMatch(t, tt.want, got)

		})
	}
}

func TestThreeSumUnorderedList(t *testing.T) {
	type args struct {
		input  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find target in sorted list",
			args: args{input: []int{5, 2, 11, 19, 15}, target: 31},
			want: []int{5, 11, 15},
		},
		{
			name: "if not present return -1,-1",
			args: args{input: []int{3, 3, 3}, target: 7},
			want: []int{-1, -1, -1},
		},
		{
			name: "if not present return -1,-1",
			args: args{input: []int{3, 1, 8, 5, 9, 4}, target: 7},
			want: []int{3, 1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSumUnorderedList(tt.args.input, tt.args.target)
			assert.ElementsMatch(t, tt.want, got)

		})
	}
}

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		input  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find target in sorted list",
			args: args{input: []int{2, 5, 7, 11, 15, 19}, target: 25},
			want: 26,
		},
		{
			name: "find target in sorted list",
			args: args{input: []int{2, 5, 7, 11, 15, 19}, target: 16},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumClosest(tt.args.input, tt.args.target)
			assert.Equal(t, tt.want, got)

		})
	}
}
