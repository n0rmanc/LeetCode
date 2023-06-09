package LeetCode

import "testing"

func Test_digitSum(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "11111222223",
				k: 3,
			},
			want: "135",
		},
		{
			args: args{
				s: "00000000",
				k: 3,
			},
			want: "000",
		},
		{
			args: args{
				s: "233",
				k: 3,
			},
			want: "233",
		},
		{
			args: args{
				s: "01234567890",
				k: 2,
			},
			want: "27",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitSum(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("digitSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
