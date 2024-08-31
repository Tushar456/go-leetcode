package slidingwindo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxSubarraySumEqualTarget(t *testing.T) {
	type args struct {
		numbers   []int
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 3, 5, -3, 2, 3, 6, 7}, targetSum: 8},
			want: []int{-3, 2, 3, 6},
		},
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 3, -1, 0, 1, 5, -3, 2, 3, 6, 7}, targetSum: 8},
			want: []int{3, -1, 0, 1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxSubarraySumEqualTarget(tt.args.numbers, tt.args.targetSum)
			assert.ElementsMatch(t, tt.want, got)

		})
	}
}

func Test_findMinSubarraySumEqualTarget(t *testing.T) {
	type args struct {
		numbers   []int
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{2, 3, 1, 2, 4, 3}, targetSum: 7},
			want: []int{4, 3},
		},
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 4, 4}, targetSum: 4},
			want: []int{4},
		},
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 1, 1, 1, 1, 1, 1}, targetSum: 11},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMinSubarraySumEqualTarget(tt.args.numbers, tt.args.targetSum)
			assert.ElementsMatch(t, tt.want, got)

		})
	}
}

func Test_findMinSubarraySumEqualGreaterTarget(t *testing.T) {
	type args struct {
		numbers   []int
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{2, 3, 1, 2, 4, 3}, targetSum: 7},
			want: 2,
		},
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 4, 4}, targetSum: 4},
			want: 1,
		},
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 1, 1, 1, 1, 1, 1}, targetSum: 11},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMinSubarraySumEqualGreaterTarget(tt.args.numbers, tt.args.targetSum)
			assert.Equal(t, tt.want, got)

		})
	}
}
