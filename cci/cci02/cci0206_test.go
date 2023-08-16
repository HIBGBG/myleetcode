package cci02

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1",args{head: SliceToListNode([]int{1,2,2,1})},true},
		//{"case2",args{head: SliceToListNode([]int{})},true},
		//{"case3",args{head: SliceToListNode([]int{1})},true},
		//{"case4",args{head: SliceToListNode([]int{1,2})},false},
		//{"case5",args{head: SliceToListNode([]int{1,1})},true},
		//{"case6",args{head: SliceToListNode([]int{1,2,2})},false},
		//{"case7",args{head: SliceToListNode([]int{1,2,1})},true},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome_old(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1",args{head: SliceToListNode([]int{1,2,2,1})},true},
		{"case2",args{head: SliceToListNode([]int{})},true},
		{"case3",args{head: SliceToListNode([]int{1})},true},
		{"case4",args{head: SliceToListNode([]int{1,2})},false},
		{"case5",args{head: SliceToListNode([]int{1,1})},true},
		{"case6",args{head: SliceToListNode([]int{1,2,2})},false},
		{"case7",args{head: SliceToListNode([]int{1,2,1})},true},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome_old(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}