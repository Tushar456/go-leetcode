package twopointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumUnorderedListBruteForce(t *testing.T) {
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
			args: args{input: []int{2, 7, 11, 15}, target: 9},
			want: []int{0, 1},
		},
		{
			name: "find target in unsorted list",
			args: args{input: []int{3, 2, 4}, target: 6},
			want: []int{1, 2},
		},
		{
			name: "find target in sorted list",
			args: args{input: []int{3, 3}, target: 6},
			want: []int{0, 1},
		},
		{
			name: "if not present return -1,-1",
			args: args{input: []int{3, 3}, target: 7},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSumUnorderedListBruteForce(tt.args.input, tt.args.target)
			assert.ElementsMatch(t, tt.want, got)

		})
	}
}

func TestTwoSumUnorderedList(t *testing.T) {
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
			args: args{input: []int{2, 7, 11, 15}, target: 9},
			want: []int{0, 1},
		},
		{
			name: "find target in unsorted list",
			args: args{input: []int{3, 2, 4}, target: 6},
			want: []int{1, 2},
		},
		{
			name: "find target in sorted list",
			args: args{input: []int{3, 3}, target: 6},
			want: []int{0, 1},
		},
		{
			name: "if not present return -1,-1",
			args: args{input: []int{3, 3}, target: 7},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSumUnorderedList(tt.args.input, tt.args.target)
			assert.ElementsMatch(t, tt.want, got)

		})
	}
}

func TestTwoSumOrderedList(t *testing.T) {
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
			args: args{input: []int{2, 7, 11, 15}, target: 9},
			want: []int{0, 1},
		},
		{
			name: "find target in unsorted list",
			args: args{input: []int{2, 3, 8, 11}, target: 11},
			want: []int{1, 2},
		},
		{
			name: "if not present return -1,-1",
			args: args{input: []int{3, 3}, target: 7},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSumOrderedList(tt.args.input, tt.args.target)
			assert.ElementsMatch(t, tt.want, got)

		})
	}
}

func TestTwoSumUnorderedListMultiple(t *testing.T) {
	type args struct {
		input  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "find target in sorted list",
			args: args{input: []int{1, 2, 7, 8, 11, 15}, target: 9},
			want: [][]int{{0, 3}, {1, 2}},
		},
		{
			name: "if not present return -1,-1",
			args: args{input: []int{3, 3, 3, 3}, target: 7},
			want: [][]int{{-1, -1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSumUnorderedListMultiple(tt.args.input, tt.args.target)
			assert.ElementsMatch(t, tt.want, got)

		})
	}
}

func TestTwoSumOrderedListMultiple(t *testing.T) {
	type args struct {
		input  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "find target in sorted list",
			args: args{input: []int{1, 2, 7, 8, 11, 15}, target: 9},
			want: [][]int{{0, 3}, {1, 2}},
		},
		{
			name: "if not present return -1,-1",
			args: args{input: []int{3, 3, 3, 3}, target: 7},
			want: [][]int{{-1, -1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSumOrderedListMultiple(tt.args.input, tt.args.target)
			assert.ElementsMatch(t, tt.want, got)

		})
	}
}

func TestTwoSumWqualThirdElement(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find target in sorted list",
			args: args{input: []int{1, 2, 7, 8, 11, 20}},
			want: []int{0, 2, 3},
		},
		{
			name: "if not present return -1,-1",
			args: args{input: []int{3, 3, 3, 3}},
			want: []int{-1, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSumWqualThirdElement(tt.args.input)
			assert.ElementsMatch(t, tt.want, got)

		})
	}

}
