package cci01

import "testing"

func Test_oneEditAway(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1",args{
			first:  "pale",
			second: "ple",
		},true},
		{"case2",args{
			first:  "ple",
			second: "pale",
		},true},
		{"case3",args{
			first:  "plaes",
			second: "plae",
		},true},
		{"case4",args{
			first:  "plaesd",
			second: "plae",
		},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneEditAway(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("oneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}