package main

func swap(a int, b int) {
	var tmp int
	tmp = a
	a = b
	b = tmp
}

func swap2(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}
func main() {

	var a, b int = 100, 200

	swap(a, b)
	//打印出来a= 100 b= 200
	//为什么a= 100中间有个空格?
	println("a=", a, "b=", b)

	swap2(&a, &b)
	println("a=", a, "b=", b)
}
