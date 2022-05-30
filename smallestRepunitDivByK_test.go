package LeetCode

import "testing"

func Test_smallestRepunitDivByK(t *testing.T) {
	type args struct {
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "A",
			args: args{
				K: 1,
			},
			want: 1,
		},
		{name: "B",
			args: args{
				K: 2,
			},
			want: -1,
		},
		{name: "C",
			args: args{
				K: 3,
			},
			want: 3,
		},
		{name: "D",
			args: args{
				K: 7,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRepunitDivByK(tt.args.K); got != tt.want {
				t.Errorf("smallestRepunitDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
