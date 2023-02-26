package main

import (
	"fmt"
)

func main() {
	//var user="rechard"
	user := "rechard"
	const GO = "go"
	var result = fmt.Sprintf("%s: hello %s"+",hello world \n", user, GO)
	fmt.Printf(result)
	var num uint8 = 255
	result = fmt.Sprintf("%s: %v \n", user, num)
	fmt.Printf(result)
	var c bool = false
	fmt.Println(c)
	//var num2 int
	num2 := 2
	fmt.Println(num2)

	var s = 0
	for i := 0; i < 10; i++ {
		s += i
	}
	fmt.Println(s)

	var b int = 4
	var ptr *int
	ptr = &b
	println(b)
	println(ptr)
	println(*ptr)

	//cmd := exec.Command("sh")
	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	Cloneflags: syscall.CLONE_NEWUTS,
	//}
	//cmd.Stdin = os.Stdin
	//cmd.Stdout = os.Stdout
	//cmd.Stdin = os.Stdin
	//cmd.Stderr = os.Stderr
	//if err := cmd.Run(); err != nil {
	//	log.Fatal(err)
	//}
	m := make(map[string]int)
	m["a"] = 1
	m["c"] = 3
	println(m["c"])

	for k, v := range m {
		println(k, "=", v)
	}

	arr := make([]string, 4)
	println(arr[0])
	arr2 := [5]string{"a", "b", "c", "d", "e"}
	println(arr2[3])
	slice2 := arr2[2:4]
	println(cap(slice2), len(slice2))
	//slice2[2] = "f"
	newslice := append(slice2, "f")
	println(newslice[2])

	for i := 0; i < len(slice2); i++ {
		println(slice2[i])
	}
	println("====================")
	cli := closure()
	println(cli())
	println(cli())
	println(cli())

	//%% 就是为了在输出的字符里有%， 第一个%是转义符 %
	//0  0
	//%d 6
	//d  d
	//%06d
	println(fmt.Sprintf("%%0%dd", 6))
}

type Age uint

func (age Age) String() {
	println("the age is", age)
}

func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
