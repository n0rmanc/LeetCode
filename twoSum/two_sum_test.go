package twoSum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{`first`, args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{`second`, args{[]int{-3, 4, 3, 90}, 0}, []int{0, 2}},
		{`third`, args{[]int{0, 3, -3, 4, -1}, -1}, []int{0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
