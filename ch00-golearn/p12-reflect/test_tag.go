package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserName string `info:"name" doc:"your name"`
	Age      int    `info:"age" doc:"your age"`
}

// 下面的例子里是演示的不同值传递的时候，reflect获取的方式不同
// 传递的是指针，就用TypeOf().Elem()。传递的是对象本身，就使用TypeOf()
func main() {
	user := User{
		"rechard",
		20,
	}

	//注意这里传的是实际的类型
	userType := reflect.TypeOf(user)
	fmt.Printf("%d \n", userType.NumField())

	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		fmt.Printf("filed=%v,tag=%s \n", field.Type, field.Tag)
	}

	//注意这里传的是指针类型
	userType2 := reflect.TypeOf(&user)
	fmt.Printf("%d \n", userType2.Elem().NumField())

	for i := 0; i < userType2.Elem().NumField(); i++ {
		field := userType2.Elem().Field(i)
		fmt.Printf("filed=%v,tag=%s \n", field.Type, field.Tag)
	}

}
