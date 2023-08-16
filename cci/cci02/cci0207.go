package cci02

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for headA==nil||headB==nil{
		return nil
	}

	h1,h2:=headA,headB
	// h11,h22:=headA,headB
	for h1.Next!=nil&&h2.Next!=nil{
		h1=h1.Next
		h2=h2.Next
	}
	// fmt.Println(h1.Val,h2.Val)
	for h1.Next!=nil&&h2.Next==nil{//不是if
		//h1长。
		h1=h1.Next
		headA=headA.Next//  hA 先走几步
		// fmt.Println("==",h1.Val,headA.Val)
	}
	for h2.Next!=nil&&h1.Next==nil{//不是if
		//h2长。
		h2=h2.Next
		headB=headB.Next//hB 先走几步
	}

	for headA!=nil{
		// fmt.Println(headA.Val,headB.Val)
		if headA==headB{
			return headB
		}
		headA=headA.Next
		headB=headB.Next
	}

	return nil

}
