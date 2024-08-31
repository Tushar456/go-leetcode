package twopointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicate(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{
			name:  "move zeros to right",
			args:  args{numbers: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want:  5,
			want1: []int{0, 1, 2, 3, 4},
		},
		{
			name:  "move zeros to right",
			args:  args{numbers: []int{1, 1, 2}},
			want:  2,
			want1: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeDuplicate(tt.args.numbers)

			assert.Equal(t, tt.want, got)
			assert.ElementsMatch(t, tt.want1, got1)

		})
	}
}
