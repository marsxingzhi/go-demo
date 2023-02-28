package main

import (
	"container/list"
	"fmt"
)

// container/list标准库，双向链表

func main() {

	l := list.New()

	l.PushBack(1)
	l.PushBack(3)
	l.PushFront(9)

	// 遍历 l.Font() 返回第一个节点
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("value: ", e.Value)
	}
}
