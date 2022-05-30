package LeetCode

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "A",
			args: args{
				s: "1*2*3",
			},
			want: 6,
		},
		{name: "B",
			args: args{
				s: "1+2+3",
			},
			want: 6,
		},
		{name: "C",
			args: args{
				s: " 3+5 / 2 ",
			},
			want: 5,
		},
		{name: "D",
			args: args{
				s: " 3*5 + 2 * 4 ",
			},
			want: 23,
		},
		{name: "E",
			args: args{
				s: "0-2147483647",
			},
			want: -2147483647,
		},
		{name: "F",
			args: args{
				s: "1-1+1",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
