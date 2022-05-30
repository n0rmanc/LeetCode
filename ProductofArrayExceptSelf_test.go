package LeetCode

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{-1, 1, 0, -3, 3},
			},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{0, 1, 0, -3, 3},
			},
			want: []int{0, 0, 0, 0, 0},
		},
		{
			name: "case 4",
			args: args{
				nums: []int{0, 0, 0, 0, 0},
			},
			want: []int{0, 0, 0, 0, 0},
		},
		{
			name: "case 5",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: []int{1, 1, 1, 1, 1},
		},
		{
			name: "case 6",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
		{
			name: "case 7",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
