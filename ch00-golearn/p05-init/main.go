package main

import (
	mylib1 "zinx/ch00-golearn/p05-init/lib1"
	_ "zinx/ch00-golearn/p05-init/lib2"
)

/*
*
init 早于 main函数

导包方式
1. mylib1 "zinx/ch00-golearn/p05-init/lib1" 别名
2. _ "zinx/ch00-golearn/p05-init/lib2" 匿名
3. . "zinx/ch00-golearn/p05-init/lib2" 将lib2引入到当前，使用不多，避免使用
例如 lib1和lib2里都有同名函数，使用.导入2个，会报错
*/
func main() {
	//lib1.Test()
	mylib1.Test()
	//lib2.Test()
}
