package arrays

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{

			name: "k top element",
			args: args{nums: []int{9, 3, 3, 3, 15, 15, 15, 1, 8, 1, 0}, k: 3},
			want: []int{3, 15, 1},
		},
		{

			name: "k top element",
			args: args{nums: []int{9, 3, 3, 15, 15, 1, 8, 1, 0}, k: 2},
			want: []int{3, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
