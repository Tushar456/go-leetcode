package recurrsion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fibonaci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "fibonacii",
			args: args{n: 5},
			want: 5,
		},
		{
			name: "fibonacii",
			args: args{n: 6},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonaci(tt.args.n); got != tt.want {
				t.Errorf("fibonaci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "fibonacii",
			args: args{n: 5},
			want: 120,
		},
		{
			name: "fibonacii",
			args: args{n: 3},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.args.n); got != tt.want {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find element pos",
			args: args{arr: []int{2, 5, 7, 9, 15, 19}, target: 5},
			want: 1,
		},
		{
			name: "find element pos",
			args: args{arr: []int{2, 5, 7, 9, 15, 19}, target: 1},
			want: -1,
		},
		{
			name: "find element pos",
			args: args{arr: []int{2, 5, 7, 9, 15, 19}, target: 2},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum of digit",
			args: args{n: 1543},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfDigit(tt.args.n); got != tt.want {
				t.Errorf("sumOfDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverese(t *testing.T) {
	type args struct {
		arr   []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "reverese the number",
			args: args{arr: []int{2, 5, 7, 9, 15, 19}},
			want: []int{19, 15, 9, 7, 5, 2},
		},
		{
			name: "reverese the number",
			args: args{arr: []int{1, 2, 5, 7, 9, 15, 19}},
			want: []int{19, 15, 9, 7, 5, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverese(tt.args.arr, tt.args.index)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func Test_printSubsequences(t *testing.T) {
	type args struct {
		arr     []int
		index   int
		currSub []int
		result  [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "print subsequence",
			args: args{arr: []int{1, 3, 4}, index: 0, currSub: []int{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printSubsequences(tt.args.arr, tt.args.index, tt.args.currSub, &tt.args.result)
			fmt.Println(tt.args.result)
		})
	}
}

func TestSubsequencesSumEqualK(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "print subsequence",
			args: args{arr: []int{1, 3, 4, 5, 7}, target: 14},
			want: [][]int{{3, 4, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubsequencesSumEqualK(tt.args.arr, tt.args.target)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "merge sort",
			args: args{arr: []int{6, 1, 4, 3, 4, 5, 7}},
			want: []int{1, 3, 4, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeSort(tt.args.arr)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
