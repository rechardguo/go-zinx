package main

import (
	"fmt"
)

/*
*
%d 数字
%p 指针
%v 打印整体
%+v 整体加上字段名
%T type

*/

type Person struct {
	UserName string
	Age      int
}

func main() {
	person := Person{
		"rechard",
		18,
	}

	fmt.Printf("%v \n", person)
	//加上字段名
	fmt.Printf("%+v \n", person)
	fmt.Printf("%T \n", person)
	var arg interface{}
	arg = &person

	fmt.Printf("------------------- \n")
	fmt.Printf("%v \n", arg)
	fmt.Printf("%+v \n", arg)
	//*main.Person 前面的*表示是指针类型
	fmt.Printf("%T \n", arg)

}
