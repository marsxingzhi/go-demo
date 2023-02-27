package main

import (
	"container/heap"
	"fmt"
)

// 利用container/heap实现堆
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

// 官方的example
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 决定是最小堆，还是最大堆
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 这里为什么要用指针，只有用指针，添加元素时，切片的长度和内容才会改变
// 可以看下：https://xiongliuhua.com/golang/144/
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// 永远都是最后一个出去，可以看下heap.go源码，就知道为啥了，其实方法内部做了交换，将堆顶元素与最后元素交换了
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
