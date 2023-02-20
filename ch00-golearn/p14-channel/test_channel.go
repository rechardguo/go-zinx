package main

/*
*
1. channel 是2个携程之间的通信中介
2. channel有阻塞等待的作用
*/
func main() {

	ch := make(chan int)
	go func() {
		defer println("sub defer end")

		println("sub end")
		ch <- 888
	}()

	num := <-ch
	println("num=", num)
}
