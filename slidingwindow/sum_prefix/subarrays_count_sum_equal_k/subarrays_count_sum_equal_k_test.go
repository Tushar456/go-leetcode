package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_findSubArrayWhoseSumEqualtoTarget(t *testing.T) {
	type args struct {
		numbers []int
		targert int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name:       "find sum",
			args:       args{numbers: []int{1, 4, -4, 5, -5, 0}, targert: 1},
			wantResult: 5,
		},
		{
			name:       "find sum",
			args:       args{numbers: []int{1, 4, -4, 5, -5, 0}, targert: 0},
			wantResult: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := findSubArrayWhoseSumEqualtoTarget(tt.args.numbers, tt.args.targert); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("findSubArrayWhoseSumEqualtoTarget() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
