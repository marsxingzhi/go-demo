package main

import (
	"fmt"

	"github.com/mars/go-demo/cache"
)

func main() {
	fmt.Println("Hello World")

	// test cache
	cache := cache.New()

	cacheTable := cache.Create("test")
	res, exist := cacheTable.Get("k1")
	if !exist {
		fmt.Println("Not Found value of k1")
		return
	}
	fmt.Println("res: ", res)
}
