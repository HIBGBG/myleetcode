package cci01

import "testing"

func Test_compressString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	{"case1",args{"aabcccccaaa"},"a2b1c5a3"},
		{"case2",args{"abbccd"},"abbccd"},
		{"case3",args{"bb"},"bb"},
		{"case4",args{"TTTTTUUUUUTTTTTTTTBBBBBBBPPPPPPPPDDDDDDllllllhhhhhhvvvvvvvUQQQQvvvvvvH"},"T5U5T8B7P8D6l6h6v7U1Q4v6H1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.args.S); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compressString2(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1",args{"aabcccccaaa"},"a2b1c5a3"},
		{"case2",args{"abbccd"},"abbccd"},
		{"case3",args{"bb"},"bb"},
		{"case4",args{"TTTTTUUUUUTTTTTTTTBBBBBBBPPPPPPPPDDDDDDllllllhhhhhhvvvvvvvUQQQQvvvvvvH"},"T5U5T8B7P8D6l6h6v7U1Q4v6H1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString2(tt.args.S); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}