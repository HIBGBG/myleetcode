package cci01

import "testing"

func Test_canPermutePalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1",args{"code"},false},
		{"case2",args{"tactcoa"},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPermutePalindrome(tt.args.s); got != tt.want {
				t.Errorf("canPermutePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}