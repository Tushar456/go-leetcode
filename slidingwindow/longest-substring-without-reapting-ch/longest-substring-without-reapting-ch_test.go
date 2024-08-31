package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		{
			name:  "the longest substring length is",
			args:  args{s: "abcdababaef"},
			want:  4,
			want1: "baef",
		},
		{
			name:  "the longest substring length is",
			args:  args{s: "aaaaaa"},
			want:  1,
			want1: "a",
		},
		{
			name:  "the longest substring length is",
			args:  args{s: "pwwkew"},
			want:  3,
			want1: "kew",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := lengthOfLongestSubstring(tt.args.s)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)

		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "the longest palidrome substring",
			args: args{s: "abcababaef"},
			want: "ababa",
		},
		{
			name: "the longest palidrome substring if all character are same",
			args: args{s: "aaaaaa"},
			want: "aaaaaa",
		},
		{
			name: "the longest substring if palidrome substring has only 2 character",
			args: args{s: "pwwkew"},
			want: "",
		},
		{
			name: "the longest substring if it is not palidrome",
			args: args{s: "abcdef"},
			want: "",
		},
		{
			name: "the longest substring length is",
			args: args{s: "ababef"},
			want: "aba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.args.s)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_lengthOfLongestSubstringHavingSameLetter(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "the longest substring length is",
			args: args{s: "aabccbb", k: 2},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstringHavingSameLetter(tt.args.s, tt.args.k)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_lengthOfLongestSubstringBruteForce(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "the longest substring length is",
			args: args{s: "abcdababaef"},
			want: 4,
		},
		{
			name: "the longest substring length is",
			args: args{s: "aaaaaa"},
			want: 1,
		},
		{
			name: "the longest substring length is",
			args: args{s: "pwwkew"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstringBruteForce(tt.args.s)
			assert.Equal(t, tt.want, got)

		})
	}
}
