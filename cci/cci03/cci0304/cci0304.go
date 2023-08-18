package cci0304




type MyQueue struct {
	inStack,outStack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.inStack=append(this.inStack,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.outStack)==0{
		this.in2out()
	}
	l:=len(this.outStack)
	res:=this.outStack[l-1]
	this.outStack=this.outStack[:l-1]
	return res

}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty(){
		return -1
	}
	if len(this.outStack)==0{
		this.in2out()
	}
	return this.outStack[len(this.outStack)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack)==0 && len(this.outStack)==0
}

func (this *MyQueue) in2out() {

	for l:=len(this.inStack);l!=0;l--{//到0
		this.outStack=append(this.outStack,this.inStack[l-1])
		// fmt.Println(this.outStack)
	}
	// this.inStack=this.inStack[:]
	//直接重置为空切片,没有长度
	this.inStack=[]int{}

	// for len(this.inStack)>0{
	// 	this.outStack=append(this.outStack,this.inStack[len(this.inStack)-1])
	// 	this.inStack=this.inStack[:len(this.inStack)-1]
	// }

}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */