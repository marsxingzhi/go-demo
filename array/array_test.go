package array

import (
	"fmt"
	"testing"
)

func TestCreateMatrix(t *testing.T) {
	arr := CreateMatrix()
	fmt.Printf("rows: %v, cols: %v\n", len(arr.Matrix), len(arr.Matrix[0]))
}
