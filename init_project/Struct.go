package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// json结构体
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func main() {
	var p1 Person
	p1.Name = "Tom"
	p1.Age = 30
	fmt.Println("p1 = ", p1)

	var p2 = Person{Name: "Burke", Age: 31}
	fmt.Println("p2=", p2)

	p3 := Person{Name: "Aaron", Age: 32}
	fmt.Println("p1=", p3)

	// 匿名结构体
	p4 := struct {
		Name string
		Age  int
	}{Name: "匿名", Age: 33}

	fmt.Println("p4 = ", p4)

	// Result 结构体变量 res赋值
	var res Result
	res.Code = 200
	res.Message = "success"

	// 序列化（形成json），如果json序列化成功 err 等于 nil
	jsons, err := json.Marshal(res)
	if err != nil {
		fmt.Println("json marshal error:")
	}
	fmt.Println("json data", string(jsons))

	// 反序列化 Unmarshal
	var res2 Result
	err = json.Unmarshal(jsons, &res2)
	if err != nil {
		fmt.Println("json unmarshal err:", err)
	}
	fmt.Println("res2:", res2)
	// &res是res的地址
	toJson(&res)
	// 修改参数
	setData(&res)
	toJson(&res)
}

// 参数为指针类型
func setData(res *Result) {
	res.Code = 500
	res.Message = "fail"
}

// 将结构体数据，序列化为Json
func toJson(res *Result) {
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(jsons))
}
