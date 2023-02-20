package main

import (
	"fmt"
	"sync"
)

func hello(i int, group *sync.WaitGroup) {
	fmt.Printf("hello %d \n", i)
	group.Done()
}
func main() {
	var group sync.WaitGroup
	group.Add(100)
	for i := 0; i < 100; i++ {
		go hello(i, &group)
	}
	println("main")
	group.Wait()
}
