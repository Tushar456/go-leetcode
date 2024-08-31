package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestSubstringWithAtmost2DistanctCharacter(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find subarray with two distanct element",
			args: args{input: "eceba"},
			want: 3,
		},
		{
			name: "find subarray with two distanct element",
			args: args{input: "ccaabbb"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubstringWithAtmost2DistanctCharacter(tt.args.input)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_longestSubstringWithAtmostKDistanctCharacter(t *testing.T) {
	type args struct {
		input string
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find subarray with two distanct element",
			args: args{input: "eceba", k: 2},
			want: 3,
		},
		{
			name: "find subarray with two distanct element",
			args: args{input: "ccaabbbde", k: 3},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubstringWithAtmostKDistanctCharacter(tt.args.input, tt.args.k)

			assert.Equal(t, tt.want, got)
		})
	}
}
