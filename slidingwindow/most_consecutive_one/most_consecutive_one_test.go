package slidingwindow

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_findMaxConsecutiveOneWithoutFliping(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{1, 1, 0, 1, 1, 1}},
			want: 3,
		},

		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{1, 1, 1, 1, 0, 1}},
			want: 4,
		},

		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{1, 1, 1, 1, 1, 1}},
			want: 6,
		},
		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{0, 0, 0, 0, 0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxConsecutiveOneWithoutFliping(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findMaxConsecutiveOneWithOneFliping(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{1, 0, 1, 1, 0}},
			want: 4,
		},
		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{1, 1, 0, 1, 1, 1}},
			want: 6,
		},

		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{1, 0, 1, 1, 0, 1}},
			want: 4,
		},

		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{1, 1, 1, 1, 1, 1}},
			want: 6,
		},
		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{0, 0, 0, 0, 0}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxConsecutiveOneWithOneFliping(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findMaxConsecutiveOneWithKFliping(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{1, 0, 1, 1, 0}, k: 1},
			want: 4,
		},
		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{1, 0, 1, 1, 0}, k: 2},
			want: 5,
		},

		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{0, 0, 0, 0, 0, 1}, k: 3},
			want: 4,
		},

		{
			name: "max consecutive one without flipping",
			args: args{nums: []int{0, 0, 0, 0, 0}, k: 5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxConsecutiveOneWithKFliping(tt.args.nums, tt.args.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
