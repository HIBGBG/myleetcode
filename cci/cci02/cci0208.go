package cci02

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast,slow:=head,head

	for fast!=nil&& fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow{
			break
		}
	}
	if fast==nil||fast.Next==nil{
		return nil
	}

	fast=head
	for fast!=slow{
		slow=slow.Next
		fast=fast.Next
	}
	return slow



}

