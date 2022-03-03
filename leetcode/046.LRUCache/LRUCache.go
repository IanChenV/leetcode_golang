package leetcode

type Node struct {
	Key, Val  int
	Pre, Next *Node
}
type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}

func Constructor(capacity int) LRUCache{
	return LRUCache{Keys: make(map[int]*Node), Cap: capacity)}
}

func (this *LRUCache) Get(key int) int{
	if node, ok := this.Keys[key]; ok {
		// 删除节点后，移动到头部
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}



func (this *LRUCache) Remove(node *Node){
	// 节点为head
     if node == this.head {
		 this.head = node.Next
		 node.Next = nil
		 return 
	 }
	 if node == this.tail {
		 this.tail.Pre = node.Pre
		 node.Pre.Next = nil
		 node.Pre = nil
		 return 
	 }
	// 节点为tail
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) Add(node *Node){
	// 头插入
	node.Pre = nil
	node.Next = this.head

	if this.head != nil {
		this.head.Pre = node
	}

	this.head = node
	this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}


func (this *LRUCache) Put(key int, value int){
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	}else{
		node = &Node{Key: key, Val: value}
		this.Keys[key] = node
		this.Add(node)
	}
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}