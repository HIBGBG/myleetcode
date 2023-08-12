package cci02

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head==nil{
		return head
	}

	m := map[int]bool{}
	h := head
	m[h.Val]=true
	for h.Next != nil { //h!=nil，解决不了末尾指针。。
		if m[h.Next.Val] == true {
			//存在，进行删除
			h.Next=h.Next.Next
		} else {
			//不存在，加入map
			m[h.Next.Val] = true
			//往后遍历
			h=h.Next
		}

	}
	return head
}

func removeDuplicateNodes_before(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	cnt:=map[int]int{head.Val:1,}
	cur:=head
	for cur.Next!=nil{
		if cnt[cur.Next.Val]!=0{//已出现跳过
			cur.Next=cur.Next.Next
		}else{//未出现
			cnt[cur.Next.Val]=1
			cur=cur.Next
		}
	}
	// head.Next=nil
	cur.Next=nil

	return head
}


func removeDuplicateNodes_Wrong(head *ListNode) *ListNode {
	//链表去重
	//map 记录


	m:=map[int]bool{}
	h:=head
	// m[h.Val]=true
	for h!=nil{//h!=nil，解决不了末尾指针。。
		if m[h.Val]==true{
			//存在，进行删除
			if h.Next!=nil{
				h.Val=h.Next.Val//下个值往前移，
				h.Next=h.Next.Next//继续检查当前值
			}else{
				//是最后一个。
				h=nil  //     !!!!!!!!  解决不了删除最后一个值 ~!~!!!!!!!!!!!!!
			}
		}else{
			//不存在，加入map
			m[h.Val]=true
			//往后遍历
			h=h.Next
		}

	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
