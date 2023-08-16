package cci02

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	//链表转切片
	tmp:=[]int{}
	for head!=nil{
		tmp=append(tmp,head.Val)
		head=head.Next
	}
	l:=len(tmp)
	for i:=0;i<l/2;i++{
		if tmp[i]!=tmp[l-1-i]{
			return false
		}
	}

	return true



	//fail  原来都会写。。。现在反而不会了。。。。 ！！！！
	//if head == nil || head.Next == nil { //空或是一个元素
	//	return true
	//}
	////先找中间节点
	//fast := head
	//slow := head
	//for fast != nil && fast.Next != nil {
	//	fast = fast.Next.Next
	//	slow = slow.Next
	//}
	////print(slow.Val,slow.Next.Val)
	////if fast!=nil{
	////	print(fast.Val,slow.Val)
	////}
	////奇偶影响分析
	////head  slow fast
	////1     2  3
	////1    2 3  4
	////1  3  5
	//
	//// 偶数需要翻转slow 到最后一个
	//if fast == nil {
	//	//找到slow上一个节点
	//	tmp := head
	//	for tmp.Next != slow {
	//		tmp = tmp.Next
	//	}
	//	slow = tmp
	//}
	//fmt.Println(slow.Val,slow.Next.Val)
	////把偶数个也转换成翻转带头指针的，和下面奇数个一起翻转
	////	奇数  翻转slow后面的  带头结点 后面链表翻转
	//tail := slow.Next
	//if tail == nil { //如果只有两个元素
	//	return head.Val == head.Next.Val
	//}
	//cur := tail.Next
	//if cur == nil { //如果只有三个元素
	//	return head.Val == head.Next.Next.Val
	//
	//}
	//fmt.Println(head.Val, slow.Val, tail.Val, cur.Val)
	//for cur.Next != nil {
	//	//应该是头插法  在slow和slow.Next中间插入后面的节点
	//	tail = slow.Next
	//	slow.Next = cur
	//	cur = cur.Next
	//	slow.Next.Next = tail
	//}
	//for slow != nil {
	//	fmt.Println(head.Val, slow.Val,slow.Next.Val)
	//	if head.Val != slow.Val {
	//		return false
	//	} else {
	//		head = head.Next
	//		slow = slow.Next
	//	}
	//}
	//return true

}


func isPalindrome_old(head *ListNode) bool {
	//翻转整个链表，再逐个比较
	//略

	if head==nil||head.Next==nil{
		return true

	}
	//到中间节点，翻转后面节点（头插法），比较前一半和后一半
	fast,slow:=head,head
	for fast!=nil&&fast.Next!=nil{
		fast=fast.Next.Next
		if fast==nil{//slow.Next为偶数个的后半段第一个，要先break
			break
		}
		slow=slow.Next
	}
	// fmt.Println(slow.Next.Val)//要把偶数个的后半段第一个翻转到最后
	// 翻转
	// head  。。。。  slow  ....  fast
	// slow ,3,2,1
	head1:=slow.Next
	cur:=head1.Next
	head1.Next=nil
	for cur!=nil{
		n:=cur.Next
		tmp:=slow.Next
		slow.Next=cur
		cur.Next=tmp
		cur=n

	}
	//验证翻转
	// for head!=nil{
	//     fmt.Print(head.Val)
	//     head=head.Next
	// }
	// fmt.Println(slow)

	//前半段和翻转的后半段比较
	slow=slow.Next
	for slow!=nil{
		if slow.Val!=head.Val{
			return false
		}
		slow=slow.Next
		head=head.Next
	}

	return true
}