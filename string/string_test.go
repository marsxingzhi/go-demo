package string

import (
	"fmt"
	"testing"
)

func TestStr2Int(t *testing.T) {
	testStr := "12"
	// testStr1 := "Hello World"

	val, err := str2Int(testStr)
	if err != nil {
		fmt.Println("err: ", err)
		panic(err)
	}
	fmt.Println("val: ", val)
}

func TestStr2Int64(t *testing.T) {
	testStr := "120000"
	val, err := str2Int64(testStr)
	if err != nil {
		fmt.Println("err: ", err)
		panic(err)
	}
	fmt.Println("val: ", val)
}

func TestInt2String(t *testing.T) {
	s := int2String(10)
	fmt.Println("s: ", s)
}

func TestInt64ToString(t *testing.T) {
	s := int64ToString(100)
	fmt.Println("s: ", s)
}
