package recurssion

import (
	"fmt"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		arr    []int
		index  int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "combination sum",
			args: args{arr: []int{2, 2, 3}, target: 7},
			want: [][]int{{2, 2, 3}, {7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.args.arr, tt.args.index, tt.args.target)
			fmt.Println(got)
			//assert.ElementsMatch(t, tt.want, got)
		})
	}
}
