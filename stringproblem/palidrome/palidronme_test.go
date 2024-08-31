package stringproblem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPalindromeString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "reverese string with basic approach",
			args: args{str: "cacac"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindPalindromeString(tt.args.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFindPalindromeTwoPointer(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "reverese string with basic approach",
			args: args{str: "cacac"},
			want: true,
		},
		{
			name: "reverese string with basic approach",
			args: args{str: "cacat"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindPalindromeString(tt.args.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFindPalindromeInt(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "reverese string with basic approach",
			args: args{num: 3443},
			want: true,
		},
		{
			name: "reverese string with basic approach",
			args: args{num: 34411},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindPalindromeInt(tt.args.num)
			assert.Equal(t, tt.want, got)
		})
	}
}

// func TestFindPalindromeTwoPointerOpt(t *testing.T) {
// 	type args struct {
// 		s string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := FindPalindromeTwoPointerOpt(tt.args.s); got != tt.want {
// 				t.Errorf("FindPalindromeTwoPointerOpt() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func BenchmarkFindPalindromeTwoPointerOpt(b *testing.B) {

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test benchmak",
			args: args{s: "caaaaaaaaaaaacbcaaaaaaaaaaaac"},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindPalindromeTwoPointerOpt(tt.args.s)
			}

		})
	}
}

func BenchmarkFindPalindromeString(b *testing.B) {

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test benchmak",
			args: args{s: "caaaaaaaaaaaacbcaaaaaaaaaaaac"},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindPalindromeString(tt.args.s)
			}

		})
	}
}

// go test -bench=Bench -benchtime=1s -benchmem -count 10
func BenchmarkFindPalindromeTwoPointer(b *testing.B) {

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test benchmak",
			args: args{s: "caaaaaaaaaaaacbcaaaaaaaaaaaac"},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindPalindromeTwoPointer(tt.args.s)
			}

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
