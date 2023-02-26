package main

import (
	"fmt"
	"time"
)

/*
*
channel的缓冲机制，
chnnnel里的cap表示容量，如果满了，放的程序会阻塞住，直到拿的程序拿走，类似java里的queue
*/
func main() {

	c := make(chan int, 3)

	go func() {
		defer fmt.Println("sub process end")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Printf("channel len=%d,cap=%d \n", len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 4; i++ {
		fmt.Println("num=", <-c)
	}

	fmt.Println("main process end")
	time.Sleep(1 * time.Second)
}
