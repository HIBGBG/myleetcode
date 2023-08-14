package cci02

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	//把后面的节点放到前面来
	flag:=head//flag标记为小于x的，
	cur:=head
	//前面节点可能连续都是大于等于x
	for cur!=nil&&cur.Val>= x{//找到第一个小于x的节点
		cur=cur.Next
	}

	if cur==nil{//如果没找到
		return head
	}else{
		//交换下cur和flag的val
		tmp_val:=cur.Val
		cur.Val=flag.Val
		flag.Val=tmp_val
	}

	for  cur.Next!=nil{
		//fmt.Println(cur.Next.Val)
		if cur.Next.Val<x {
			//[1,1] 规避  不然死循环
			if cur==flag{
				cur=cur.Next
				flag=flag.Next
				continue
			}

			//把cur 放到flag后面
			tmp:=cur.Next
			cur.Next=cur.Next.Next
			tmp.Next=flag.Next
			flag.Next=tmp
			//或是把后面的值换到前面来？

		}else{
			cur=cur.Next
		}

	}

	return head


}