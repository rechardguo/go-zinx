package main

import "fmt"

// channel与select的关系
// select 可以监听多个channel
/**
下列主要是fibonacci的过程
*/
func fibonacci(num, quit chan int) {
	x, y := 1, 1
	for {
		select {
		//如果num能写入成功
		case num <- x:
			x = y
			y = x + y
			//如果quit可读就退出
		case <-quit:
			return
		}
	}
}
func main() {

	num := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			//num这里会有阻塞同步的作用
			fmt.Println(<-num)
		}
		quit <- 0
	}()

	fibonacci(num, quit)

	fmt.Printf("main end")
}
