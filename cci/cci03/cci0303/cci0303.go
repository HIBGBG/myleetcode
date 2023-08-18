package cci0303


type StackOfPlates struct {
	stacks [][]int
	cap int
	stacksize int
}


func Constructor(cap int) StackOfPlates {
	return StackOfPlates{stacks: make([][]int,0),cap: cap}
}


func (this *StackOfPlates) Push(val int)  {
	if this.cap==0{
		return
	}
	if this.stacksize==0{
		this.stacks=append(this.stacks,[]int{val})
		this.stacksize++
		return
	}

	if len(this.stacks[this.stacksize-1])==this.cap{
		this.stacks=append(this.stacks,[]int{val})
		this.stacksize++
	}else{
		this.stacks[this.stacksize-1]=append(this.stacks[this.stacksize-1],val)
	}
}


func (this *StackOfPlates) Pop() int {  //其实就是调用 PopAt(this.stacksize-1)
	if this.stacksize==0{
		return -1
	}
	if l:=len(this.stacks[this.stacksize-1]);l==1{//只有一个，就删整个切片
		res:=this.stacks[this.stacksize-1][0]
		this.stacks=this.stacks[:this.stacksize-1]
		this.stacksize--
		return res
	}else{
		res:=this.stacks[this.stacksize-1][l-1]
		this.stacks[this.stacksize-1]=this.stacks[this.stacksize-1][:l-1]
		return res
	}

}


func (this *StackOfPlates) PopAt(index int) int {
	if this.stacksize==0{
		return -1
	}

	if index<0 || index > this.stacksize-1{
		return -1
	}else{
		if l:=len(this.stacks[index]);l==1{//只有一个，就删整个切片,把后面的接起来
			res:=this.stacks[index][0]
			this.stacks=append(this.stacks[:index],this.stacks[index+1:]...)//前面一半不是[:index-1]，后一半不是[index:]
			this.stacksize--
			return res
		}else{
			res:=this.stacks[index][l-1]
			this.stacks[index]=this.stacks[index][:l-1]
			return res
		}
	}
}


/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */