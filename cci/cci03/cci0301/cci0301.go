package cci0301

//吐槽下题目，
// 三合一。描述如何只用一个数组来实现三个栈。
// 你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum)方法。stackNum表示栈下标，value表示压入的值。
// 构造函数会传入一个stackSize参数，代表每个栈的大小。


// 等价于有三个栈，stackNum表示栈下标，就是对应于不同的栈，

type TripleInOne struct {
	l []int
	stackUsed [3]int
	stackSize int
}


func Constructor(stackSize int) TripleInOne {
	return TripleInOne{stackSize:stackSize,l:make([]int,3*stackSize)}
}


func (this *TripleInOne) Push(stackNum int, value int)  {
	//检查满没满，
	if this.stackUsed[stackNum] >= this.stackSize{
		return
	}
	this.l[this.stackSize*stackNum+this.stackUsed[stackNum]]=value
	this.stackUsed[stackNum]++
}


func (this *TripleInOne) Pop(stackNum int) int {
	if this.stackUsed[stackNum]==0{
		return -1
	}
	val:=this.l[this.stackSize*stackNum+this.stackUsed[stackNum]-1]
	this.stackUsed[stackNum]--
	return val
}


func (this *TripleInOne) Peek(stackNum int) int {
	if this.stackUsed[stackNum]==0{
		return -1
	}
	val:=this.l[this.stackSize*stackNum+this.stackUsed[stackNum]-1]
	// this.stackUsed[stackNum]--
	return val
}


func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.stackUsed[stackNum]==0
}


/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
