package datastructure

import "fmt"

/**
链表
*/

type Any interface{}

// 链表的定义
type Node struct {
	Data Any
	Next *Node
}

type List struct {
	headNode *Node // 头节点
}

// 1. 判断是否为空
func (this *List) IsEmpty() bool {
	return this.headNode == nil
}

// 2. 列表长度
func (this *List) Len() int {
	len := 0
	p := this.headNode

	for p != nil {
		len++
		p = p.Next
	}
	return len
}

// 3. 从链表头部添加元素
func (this *List) Add(data Any) *Node {
	node := &Node{
		Data: data,
	}
	node.Next = this.headNode
	this.headNode = node
	return node
}

// 4. 从链表尾部添加元素
func (this *List) Append(data Any) {
	node := &Node{Data: data}

	if this.IsEmpty() {
		this.headNode = node
	} else {
		p := this.headNode

		for p.Next != nil {
			p = p.Next
		}
		p.Next = node
	}
}

// 5. 在链表指定位置添加元素
func (this *List) Insert(index int, data Any) {
	if index < 0 {
		this.Add(data)
	} else if index > this.Len() {
		this.Append(data)
	} else {
		len := 0
		p := this.headNode

		for len < (index - 1) {
			len++
			p = p.Next
		}
		node := &Node{Data: data}

		// 注意顺序
		node.Next = p.Next
		p.Next = node
	}
}

// 6. 删除链表指定值的元素
func (this *List) Remove(data Any) {
	pre := this.headNode

	if pre.Data == data {
		this.headNode = pre.Next
		return
	}
	for pre != nil {
		if pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		} else {
			if pre.Data == data {
				pre = nil
			}
		}
	}
}

// 7. 删除指定位置的元素
func (this *List) RemoveIndex(index int) {
	pre := this.headNode
	if index <= 0 {
		this.headNode = pre.Next
	} else if index > this.Len() {
		panic("待删除元素的下标超出链表长度")
	} else {
		len := 0

		for len != (index-1) && pre.Next != nil {
			len++
			pre = pre.Next
		}
		// 此时的pre是待删除元素的前一个元素
		if pre.Next != nil {
			pre.Next = pre.Next.Next
		}
	}
}

// 8. 查看链表是否包含某个元素
func (this *List) Contains(data Any) bool {
	p := this.headNode

	for p != nil {
		if p.Data == data {
			return true
		}
		p = p.Next
	}
	return false
}

// 9. 遍历链表所有节点
func (this *List) LoopList() {
	if !this.IsEmpty() {
		p := this.headNode
		for p != nil {
			fmt.Printf("current.Data: %v\n", p.Data)
			p = p.Next
		}
	}
}
