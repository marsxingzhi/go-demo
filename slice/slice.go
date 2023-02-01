package main

import "fmt"

/**
方式一：通过下标的方式获得数组或者切片的一部分

使用下标初始化切片不会拷贝原数组或者原切片中的数据，只会创建一个指向原数组的切片结构体，因此修改新切片的数据也会修改原切片
*/
func newSlice() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[0:3]

	fmt.Println("arr: ", arr)
	fmt.Println("slice: ", slice)

	// 修改slice，导致原数组中的元素也变化了
	slice[0] = 100
	fmt.Println("---")

	fmt.Println("arr: ", arr)
	fmt.Println("slice: ", slice)
}

/**
方式二：使用字面量初始化新的切片
*/
func newSlice2() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice: ", slice)
}

/**
方式三：使用关键字make创建切片
*/
func newSlice3() {
	// len：3；cap：4
	slice := make([]int, 3, 4)
	fmt.Println("slice: ", slice)
}
