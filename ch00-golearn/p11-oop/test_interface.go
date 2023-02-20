package main

import "fmt"

func foo(arg interface{}) {
	//相当于 java里的instanceof
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println(value)
	}
}

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}
type Book struct {
	author  string
	content string
}

func (this *Book) ReadBook() {
	println("reading book")
}

func (this *Book) WriteBook() {
	println("writing book")
}
func main() {
	//pair (type:Book,value:...)
	book := &Book{
		"rechard",
		"go learning",
	}

	var arg interface{}
	arg = book
	//要使用断言，必须是接口类型，不能使用实际的类
	reader, ok := arg.(Reader)
	if ok {
		reader.ReadBook()
	}

	//pair (type:,value:)
	var r Reader
	//pair (type:Book,value:...)
	r = book
	r.ReadBook()

	//pair (type:,value:)
	var w Writer
	//这里断言成功是因为 type 是 Book，而 Book 实现了 Writer 接口
	//pair (type:Book,value:...)
	w = r.(Writer)
	w.WriteBook()

	//var w Write
	//w = book
	//w.WriteBook()

}
