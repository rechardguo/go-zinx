package main

import "fmt"

func main() {

	//map的3种建立方式
	//1
	var m1 map[string]string
	if m1 == nil {
		println("map is not initialized")
	}
	//2
	m2 := make(map[string]string)
	m2["good"] = "rechard"
	m2["bad"] = "tom"
	fmt.Println(m2)

	//3
	m3 := map[int]string{
		1: "rechard",
		2: "tom",
	}
	//println 是buildin的函数不同于fmt.Println
	fmt.Println(m3)

	println("--------------------")
	//map的增删改查
	m4 := map[string]string{
		"china": "beijing",
		"uk":    "london",
		"us":    "dc",
	}
	m4["japan"] = "tokoyo"
	printMap(m4)

	//增加
	println("--------------------")
	m4["us"] = "washiton dc"
	printMap(m4)

	//删除
	println("--------------------")
	delete(m4, "us")
	printMap(m4)

	//参数传递是按地址来传递,不同于结构体是按副本传递
	println("--------------------")
	changeMap(m4)
	printMap(m4)

}

func printMap(mymap map[string]string) {
	for country, capital := range mymap {
		fmt.Printf("%s=%s \n", country, capital)
	}
}

func changeMap(mymap map[string]string) {
	mymap["us"] = "washiton dc"
}
