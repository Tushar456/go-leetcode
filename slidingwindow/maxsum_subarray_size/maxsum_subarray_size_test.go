package slidingwindow

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_findMaxSumOfSubArrayOfSizeK(t *testing.T) {
	type args struct {
		numbers []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 2},
			want: 13,
		},
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxSumOfSubArrayOfSizeK(tt.args.numbers, tt.args.k)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_findAverageSumofSubArrayOfSizeK(t *testing.T) {
	type args struct {
		numbers []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 2},
			want: []float64{2, 1, -2, 1, 4, 4.5, 6.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAverageSumofSubArrayOfSizeK(tt.args.numbers, tt.args.k)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_findMaxSumOfSubArrayOfSizeKBruteForce(t *testing.T) {
	type args struct {
		numbers []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 2},
			want: 13,
		},
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxSumOfSubArrayOfSizeKBruteForce(tt.args.numbers, tt.args.k)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_findAverageSumofSubArrayOfSizeKKBruteForce(t *testing.T) {
	type args struct {
		numbers []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "find mas sum in subarray of size 2",
			args: args{numbers: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 2},
			want: []float64{2, 1, -2, 1, 4, 4.5, 6.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAverageSumofSubArrayOfSizeKKBruteForce(tt.args.numbers, tt.args.k)
			assert.Equal(t, tt.want, got)

		})
	}
}
