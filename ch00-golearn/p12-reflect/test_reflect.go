package main

import (
	"fmt"
	"reflect"
)

func demo1() {
	var f float32 = 3.12
	//这会错误
	//println(reflect.ValueOf(arg))
	//println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(f))
	fmt.Println(reflect.TypeOf(f))
}

type Person struct {
	UserName string
	Age      int
}

//type Action interface {
//	action()
//}

// 注意方法名第一个是大写，表示公有
func (this Person) Action() {
	fmt.Println("action")
}

// 为什么这种也是NumMethod()里查不出来的
func (this *Person) Action1() {
	fmt.Println("action")
}

// 注意方法名第一个是小写，表示私有
func (this Person) action2() {
	fmt.Println("action")
}

func demo2(arg interface{}) {
	argType := reflect.TypeOf(arg)
	fmt.Println("type is", argType.Name())

	//为什么这里的method数量是0
	fmt.Println("method number is", argType.NumMethod())

	fmt.Println("field number is", argType.NumField())

	argValue := reflect.ValueOf(arg)
	fmt.Println(argValue)

	//找到所有的field
	for i := 0; i < argType.NumField(); i++ {
		// 不好理解，为什么value不放在filed里
		field := argType.Field(i)
		value := argValue.Field(i).Interface()

		fmt.Printf("%s :%v=%v \n", field.Name, field.Type, value)
	}

	//找到所有的方法
	for i := 0; i < argType.NumMethod(); i++ {
		method := argType.Method(i)
		fmt.Printf("%s:%v \n", method.Name, method.Type)
	}
}

func main() {
	person := Person{
		"rechard",
		18,
	}

	//不能够将地址传过去
	//demo2(&person)

	//这样也不行
	//var a Action
	//a = &person
	//demo2(a)

	//必须要这样
	demo2(person)

}
