package cci02


type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToListNode(s []int) *ListNode {
	head:=&ListNode{
		Val:  0,
		Next: nil,
	}
	cur:=head
	for _,v:=range s{
		cur.Next=&ListNode{
			Val:  v,
			Next: nil,
		}
		cur=cur.Next
	}
	return head.Next
}