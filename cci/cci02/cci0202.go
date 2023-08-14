package cci02



func kthToLast(head *ListNode, k int) int {
	//前后指针
	f:=head
	for i := 1; i < k; i++ {//倒数第二个，先走一步。倒数第三个，先走两步
		f=f.Next
	}

	for f.Next!=nil{
		f=f.Next
		head=head.Next
	}

	return head.Val


}


func kthToLast2(head *ListNode, k int) int {
	r:=head
	for k!=0{
		r=r.Next
		k--
	}

	for r!=nil{
		r=r.Next
		head=head.Next
	}
	return head.Val
}