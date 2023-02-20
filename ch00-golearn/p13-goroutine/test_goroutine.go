package main

import (
	"fmt"
	"time"
)

func NewTask() {
	i := 1

	for {
		i++
		fmt.Printf("child task counting %d \n", i)
		time.Sleep(time.Second * 1)
	}
}

// main如果退出了， 则里面的task也会退出
func main() {
	//go关键字起一个go程
	go NewTask()

	i := 1
	for {
		i++
		fmt.Printf("main counting %d \n", i)
		time.Sleep(time.Second * 1)
	}
}
