package main

import "fmt"

func printArrByIndex(arr [4]int, index int) {
	println(arr[index])
}

// 动态数组
func printArrByIndex2(arr []int, index int) {
	println(arr[index])
}
func main() {
	myarr1 := [4]int{1, 2, 3, 4}

	//loop array
	for i := 0; i < len(myarr1); i++ {
		println(myarr1[i])
	}

	for index, value := range myarr1 {
		fmt.Printf("myarr1[%d]=%d \n", index, value)
	}

	//传值方式
	printArrByIndex(myarr1, 3)

	//编译不过报错，要求是4个元素
	//myarr2 := [10]int{1, 2, 3, 4}
	//printArrByIndex(myarr2, 5)

	//动态数组,不指定[]里的数字
	myarr3 := []int{1, 2, 3, 4, 5, 6, 7}
	printArrByIndex2(myarr3, 5)

}
