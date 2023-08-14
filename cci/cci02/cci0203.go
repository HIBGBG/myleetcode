package cci02


func deleteNode(node *ListNode) {
	// 把下面节点的值给这个节点，删除下个节点

	node.Val=node.Next.Val
	node.Next=node.Next.Next


}

func deleteNode_old(node *ListNode) {
	//入参只有待删除的节点，没有头结点
	//找不到前驱结点，不能用cur.Next=cur.Next.Next进行
	//把下个节点值赋过来，删除下个节点


	node.Val=node.Next.Val
	node.Next=node.Next.Next

}