package main

import "fmt"

/*
结题：

B,1,2,3
D,3,2,5
E,3,2,7
C,3,5,8
A,1,3,4
*/
func main() {
	x := 1
	y := 2
	defer calc("A", x, calc("B", x, y))
	x = 3
	defer calc("C", x, calc("D", x, y))
	defer calc("E", x, y)

	defer func() {
		fmt.Println(x, y, x+y)
	}()

	y = 4
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
