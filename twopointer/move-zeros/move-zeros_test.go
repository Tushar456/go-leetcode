package twopointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeros(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "move zeros to right",
			args: args{input: []int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "move zeros to right",
			args: args{input: []int{1, 1, 0, 3, 12, 0}},
			want: []int{1, 1, 3, 12, 0, 0},
		},
		{
			name: "move zeros to right",
			args: args{input: []int{0, 1}},
			want: []int{1, 0},
		},
		{
			name: "move zeros to right",
			args: args{input: []int{0}},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := moveZeros(tt.args.input)

			assert.ElementsMatch(t, tt.want, got)

		})
	}
}

func Test_moveZeroSimple(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "move zeros to right",
			args: args{input: []int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "move zeros to right",
			args: args{input: []int{1, 2, 0, 3, 12, 0}},
			want: []int{1, 2, 3, 12, 0, 0},
		},
		{
			name: "move zeros to right",
			args: args{input: []int{0, 1}},
			want: []int{1, 0},
		},
		{
			name: "move zeros to right",
			args: args{input: []int{0}},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := moveZeroSimple(tt.args.input)

			assert.ElementsMatch(t, tt.want, got)

		})
	}
}
