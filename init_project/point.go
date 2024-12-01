package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// break 跳出循环
	for {
		fmt.Println("loop")
		break
	}
	// continue 跳出本次循环，输出小于等于5的奇数
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know  type %T\n", t)
		}
	}

	// Call the function to fix the "declared and not used" error
	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {

			twoD[i][j] = i + j
		}
		fmt.Println("2d: ", twoD)
	}

	s := make([]string, 3)
	fmt.Print("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "test")
	s = append(s, "x", "y")
	fmt.Println("app:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy", c)

	twoD1 := make([][]int, 3)
	fmt.Println("cpy", twoD1)
	/*
			for i := 0; i < 3; i++ {
		        innerLen := i + 1
		        twoD[i] = make([]int, innerLen)
		        for j := 0; j < innerLen; j++ {
		            twoD[i][j] = i + j
		        }
		    }
		    fmt.Println("2d: ", twoD)
	*/

}
