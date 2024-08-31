package checkparenthesis

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_checkBalancedParenthesisANdReturnHighest(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "check the current height",
			args: args{str: "(((X)(Y))((())))"},
			want: 4,
		},
		{
			name: "check the current height",
			args: args{str: "(((X)(Y))((()))"},
			want: -1,
		},
		{
			name: "check the current height",
			args: args{str: ")(((X)(Y))((())))("},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkBalancedParenthesisANdReturnHighest(tt.args.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_longestBalancedParenthesisInUnbalanced(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "check the current height",
			args: args{str: "((()())((())))"},
			want: 4,
		},
		// {
		// 	name: "check the current height",
		// 	args: args{str: "(((X)(Y))((()))"},
		// 	want: -1,
		// },
		// {
		// 	name: "check the current height",
		// 	args: args{str: ")(((X)(Y))((())))("},
		// 	want: -1,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestBalancedParenthesisInUnbalanced(tt.args.str)
			assert.Equal(t, tt.want, got)
		})
	}
}
