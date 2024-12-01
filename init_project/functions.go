package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

// 变参函数 不定义入参，支持多个入参

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	// 遍历
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 闭包函数 func() int, 返回闭包函数 func() int
func intSeq() func() int {
	i := 0
	// fun() int作为一个整体
	return func() int {
		i++
		return i
	}
}

type Person struct {
	Name string
	Age  int
}

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	a, b := vals()
	println(a)
	println(b)
	_, c := vals()
	println(c)
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	str := "李铂言"
	fmt.Printf(MD5(str))

}

// MD5
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	// s.Sum(nil)返回MD5数据哈希值
	return hex.EncodeToString(s.Sum(nil))
}
