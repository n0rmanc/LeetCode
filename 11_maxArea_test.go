package LeetCode

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			args: args{
				height: []int{1, 1},
			},
			want: 1,
		},
		{
			args: args{
				height: []int{1, 2, 1},
			},
			want: 2,
		},

		// {
		// 	args: args{
		// 		height: []int{2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2},
		// 	},
		// 	want: 22,
		// },
		// {
		// 	args: args{
		// 		height: []int{1, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 1},
		// 	},
		// 	want: 18,
		// },
		// {
		// 	args: args{
		// 		height: []int{1, 2, 1, 2},
		// 	},
		// 	want: 2,
		// },
		// {
		// 	args: args{
		// 		height: []int{1, 18, 2},
		// 	},
		// 	want: 18,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
