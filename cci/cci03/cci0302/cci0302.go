package cci0302



type MinStack struct {
	stack []int
	currentMin []int
	size int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{currentMin: []int{},stack: []int{}}
}


func (this *MinStack) Push(x int)  {

	l:=len(this.stack)
	// fmt.Println("====",l)
	if l==0{
		this.stack=[]int{x}
		this.currentMin=[]int{x}
		this.size++
		return
	}
	this.stack=append(this.stack,x)
	if this.currentMin[l-1]>=x{
		this.currentMin=append(this.currentMin,x)
	}else {
		this.currentMin=append(this.currentMin,this.currentMin[this.size-1])
	}
	this.size++

}


func (this *MinStack) Pop()  {
	if this.size!=0{
		this.stack=this.stack[:this.size-1]
		this.currentMin=this.currentMin[:this.size-1]
		this.size--
	}
	return

}


func (this *MinStack) Top() int {
	if this.size==0{
		return -1
	}
	return this.stack[this.size-1]
}


func (this *MinStack) GetMin() int {
	if this.size==0{
		return -1
	}
	return this.currentMin[this.size-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */