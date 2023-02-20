package main

import (
	"io"
	"os"
)

/*
 *static interface: 类型在编译过程中已知
 *dynamic interface: arg interface{}
 */
func main() {
	//f, err := os.OpenFile(os.Stdout, os.O_RDONLY, 0)
	f := os.Stdout
	var r io.Reader
	//由于file有实现了Reader的方法 Read(p []byte) (n int, err error)
	//所以这里可以直接转
	r = f

	//以下是错误的
	//这是go的断言,由于f是static interface是已知的，就不能用断言
	//换言之：要能使用断言，必须左边是接口类型
	//r = f.(io.Reader)

	var w io.Writer
	w = r.(io.Writer)
	w.Write([]byte("HELLO WORLD!"))
}
