package main

func foo1(name string) string {
	return name
}

func foo2(name string, age int) (string, int) {
	return name, age
}

// 多返回值
func foo3(name string, age int) (r1 string, r2 int) {
	r1 = name
	r2 = age
	return
}

func foo4(name string, age int) (r1 string, r2 int) {
	//默认值
	println(r1)
	println(r2)
	r1 = name
	r2 = age
	return
}
func main() {
	println("===========foo1============")
	println(foo1("rechard"))

	println("===========foo2============")
	s, i := foo2("rechard", 20)
	println(s, i)

	println("===========foo3============")
	r1, r2 := foo3("rechard", 20)
	println(r1, r2)

	println("===========foo4============")
	foo4r1, foo4r2 := foo4("rechard", 20)
	println(foo4r1, foo4r2)
}
