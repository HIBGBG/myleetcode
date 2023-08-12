package cci01

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1",args{matrix: [][]int{{1,2,3},{4,5,6},{7,8,9}}}},
		{"case2",args{matrix: [][]int{{5, 1, 9,11},{2, 4, 8,10},{13, 3, 6, 7},{15,14,12,16}}}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}