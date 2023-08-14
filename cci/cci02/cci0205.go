package cci02

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	summod10 := 0
	//补0补齐
	//rsp:=l1//哪个长不一定,还可能最后一位进位，比l1和l2都长
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := head
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 == nil {
			a = 0
		} else {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			b = 0
		} else {
			b = l2.Val
			l2 = l2.Next
		}

		flag, summod10 = (a+b+flag)/10, (a+b+flag)%10

		cur.Next = &ListNode{
			Val:  summod10,
			Next: nil,
		}
		cur = cur.Next //要前进，不然就在这个节点上修改

	}

	// 还可能最后一位进位，比l1和l2都长
	//处理进位
	if flag == 1 {
		cur.Next = &ListNode{Val: 1,
			Next: nil}
	}

	return head.Next

}
