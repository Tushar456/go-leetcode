package stringproblem

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestReverseString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverese string with basic approach",
			args: args{str: "tushar"},
			want: "rahsut",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseString(tt.args.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReverseStringAppened(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverese string with brange and appened",
			args: args{str: "tushar"},
			want: "rahsut",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseString(tt.args.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReverseStringTwoPointer(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverese string with two pointer approach",
			args: args{str: "tushar"},
			want: "rahsut",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseStringTwoPointer(tt.args.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReverseInt(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "reverese string with basic approach",
			args: args{num: 3421},
			want: 1243,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseInt(tt.args.num)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReverseWordofSentence(t *testing.T) {
	type args struct {
		senetence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverse each word of a string",
			args: args{senetence: "I am Indian"},
			want: "I ma naidnI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseWordofSentence(tt.args.senetence)
			assert.Equal(t, tt.want, got)
		})
	}
}
