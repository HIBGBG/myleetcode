package cci01

import "testing"

func Test_isUnique(t *testing.T) {
	type args struct {
		astr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1",args{ "abdefgg"}, false},
		{"case2",args{ "abcdefghijk"}, true},


	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique(tt.args.astr); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

