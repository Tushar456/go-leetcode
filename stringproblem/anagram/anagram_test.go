package stringproblem

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestIsAnagramMap(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "find anagram using map approach",
			args: args{str1: "anagram", str2: "naamarg"},
			want: true,
		},
		{
			name: "find anagram using map approach",
			args: args{str1: "anagram", str2: "naamsag"},
			want: false,
		},

		{
			name: "find anagram using map approach",
			args: args{str1: "anagr", str2: "naamsag"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAnagramMap(tt.args.str1, tt.args.str2)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_findAnagramInSlice(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "find anagram using map approach",
			args: args{input: []string{"ate", "tea", "bat", "tab", "cat"}},
			want: [][]string{{"ate", "tea"}, {"bat", "tab"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAnagramInSlice(tt.args.input)
			assert.Equal(t, tt.want, got)

		})
	}
}
