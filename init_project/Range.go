package main

import (
	"fmt"
)

func main() {
	// 遍历数组
	nums := []int{2, 3, 4}
	sum := 0
	// for ... range  ;  返回第一位：索引，返回第二位：值
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sim:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}
	// 遍历map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// 输出a、b
	for k := range kvs {
		fmt.Println("Key:", k)
	}
	// 遍历字符串
	for o, l := range "go" {
		fmt.Println(o, l)
	}
}
