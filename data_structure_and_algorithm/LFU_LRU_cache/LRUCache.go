package LFU_LRU_cache
type LRUCache struct {
	cap int

	// l  []int
	// m  map[int]int
	// 如果用数组，将会影响更新。

	head, tail *DLinkedNode
	m  map[int]*DLinkedNode
	size int//不是数组切片，就没有len()计算长度了，需要手动保存
}


func Constructor(capacity int) LRUCache {

	l:=LRUCache{
		cap:capacity,
		head: initDLinkedNode(0, 0),
		tail: initDLinkedNode(0, 0),
		m:map[int]*DLinkedNode{},
	}
	l.head.next=l.tail
	l.tail.prev=l.head
	return l
}


func (this *LRUCache) Get(key int) int {
	//读取到更新位置，
	//读取不到返回-1
	if _,ok:=this.m[key];!ok{
		return -1
	}
	node:=this.m[key]
	// this.addToHead(node)
	this.moveToHead(node)
	return node.value

}


func (this *LRUCache) Put(key int, value int)  {

	//存在:更新值，往前移位置，上面也更新位置了，可以提取出来。
	//不存在，构造node，判断还有没有容量，
	// 没有空间先老化末尾
	// 把node加到头部
	if _,ok:=this.m[key];!ok{
		if this.size==this.cap{
			last:=this.removeTail()
			delete(this.m,last.key)
			this.size--
		}
		node:=initDLinkedNode(key,value)
		this.m[key]=node
		// this.moveToHead(node)
		this.addToHead(node)
		this.size++

	}else{
		node:=this.m[key]
		node.value=value
		this.moveToHead(node)
	}

}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		value: value,
	}
}


func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */