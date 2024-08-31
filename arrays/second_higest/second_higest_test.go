package arrays

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestSecondHigeshtNumber(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			name: "find second largest",
			args: args{input: []int{9, 3, 5, 15, 8, 1, 0}},
			want: 9,
		},
		{
			name: "find second largest",
			args: args{input: []int{9, 3, 5, 15, 12, 8, 10, 8, 1, 0}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SecondHigeshtNumber(tt.args.input)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_findKHigeshtNumber(t *testing.T) {
	type args struct {
		input []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			name: "find second largest",
			args: args{input: []int{9, 3, 5, 15, 8, 1, 0}, k: 3},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKHigeshtNumber(tt.args.input, tt.args.k); got != tt.want {
				t.Errorf("findKHigeshtNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
