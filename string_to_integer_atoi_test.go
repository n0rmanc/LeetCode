package LeetCode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"42", args{"42"}, 42},
		{"   -42", args{"   -42"}, -42},
		{"4193 with words", args{"4193 with words"}, 4193},
		{"words and 987", args{"words and 987"}, 0},
		{"-91283472332", args{"-91283472332"}, -2147483648},
		{"3.14159", args{"3.14159"}, 3},
		{"+1", args{"+1"}, 1},
		{"+-2", args{"+-2"}, 0},
		{"  -0012a42", args{"  -0012a42"}, -12},
		{"   +0 123", args{"   +0 123"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
