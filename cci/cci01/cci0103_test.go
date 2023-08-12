package cci01

import "testing"

func Test_replaceSpaces(t *testing.T) {
	type args struct {
		S      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1",args{
			S:      "Mr John Smith    ",
			length: 13,
		},"Mr%20John%20Smith"},
		{
			"case2",args{
			S:      "               ",
			length: 5,
		},"%20%20%20%20%20",
		},
		{
			"case3",args{
			S:      "ds sdfs afs sdfa dfssf asdf             ",
			length: 27,
		},"ds%20sdfs%20afs%20sdfa%20dfssf%20asdf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpaces(tt.args.S, tt.args.length); got != tt.want {
				t.Errorf("replaceSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}