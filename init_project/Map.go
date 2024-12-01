package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var p1 map[int]string
	p1 = make(map[int]string)
	p1[1] = "Tom"
	fmt.Println("p1:", p1)

	p2 := map[int]string{}
	p2[2] = "Lob"
	fmt.Println("p2:", p2)
	// go语言中建议的初始化 映射的方法
	// var p3 map[int]string = make(map[int]string)

	// 通过"映射"生成Json
	res := make(map[string]interface{})
	res["cade"] = 200
	res["msg"] = "success"
	res["data"] = map[string]interface{}{
		"username": "Tom",
		"age":      "30",
		"hobby":    []string{"读书", "爬山"},
	}
	fmt.Println(res)

	// 序列化
	jsonData, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marschal error:", errs)
	}
	fmt.Println("")
	fmt.Println("-- map to Json -- ")
	fmt.Println("json data:", string(jsonData))

	// 反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsonData), &res2)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("--- json to map ---")
	fmt.Println("map data :", res2)

}
