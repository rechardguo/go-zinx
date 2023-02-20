package main

// defer相当于java里的final,会在最后执行
// defer调用是按栈的形式先进后出
func foo() int {
	defer println("defer1")
	defer println("defer2")
	defer println("defer3")
	return 1
}
func main() {
	println(foo())
}
