package main

import "fmt"

/*
*
channel close的用法
*/
func main() {

	c := make(chan int, 4)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//如果没有close会使得程序不会退出
		close(c)
	}()

	//for {
	//	if num, ok := <-c; ok {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//}

	//上面的代码可以用 range来简写
	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("main end")
}
