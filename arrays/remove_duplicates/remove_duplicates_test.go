package arrays

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicateSlice(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "remove duplicate",
			args: args{input: []int{15, 5, 7, 5, 15, 9, 0}},
			want: []int{15, 5, 7, 9, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicateSlice(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicateSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "remove duplicate",
			args: args{nums: []int{1, 1, 1, 2, 2, 3}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsDuplicateSlice(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "remove duplicate",
			args: args{input: []int{1, 2, 3, 1}},
			want: true,
		},
		{
			name: "contains duplicate",
			args: args{input: []int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "contains duplicate",
			args: args{input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicateSlice(tt.args.input); got != tt.want {
				t.Errorf("ContainsDuplicateSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
