package string

import (
	"strconv"
)

/**
string相关的操作

1. string转化成int
2. string转换成int64

3. int转化成string
4. int64转换成string
*/

// 1. string转化成int
func str2Int(str string) (int, error) {
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// 2. string转换成int64
func str2Int64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// 3. int转化成string
func int2String(val int) string {
	return strconv.Itoa(val)
}

// 4. int64转换成string
func int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}
