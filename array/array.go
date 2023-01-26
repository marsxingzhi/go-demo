package array

import "fmt"

/**
golang中二维数组的创建
*/

type Array struct {
	Matrix [][]int
}

func CreateMatrix() *Array {
	// 固定长度
	matrx := [3][4]int{}
	fmt.Println(matrx)

	// 以下报错
	// m, n := 3, 4
	// matrx1 := [m][n]int{}
	// fmt.Println(matrx1)

	// 如果要使用变量，则必须使用如下方式初始化
	m, n := 3, 4
	matrx2 := make([][]int, m)

	for index := range matrx2 {
		matrx2[index] = make([]int, n)
	}
	return &Array{
		Matrix: matrx2,
	}
}
