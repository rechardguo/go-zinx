package main

import (
	"time"
)

func main() {

	go func() {
		defer println("A defer")
		func() {
			defer println("B defer")
			//全局退出
			//runtime.Goexit()

			//只是里面的匿名函数退出，外面的不退
			//return
			println("B end")
		}()
		println("A end")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
