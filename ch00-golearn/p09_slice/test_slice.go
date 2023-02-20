package main

import "fmt"

/*
*
slice
有用的函数
len()
cap()
copy()
make()
*/
func main() {

	s1 := make([]int, 3)
	fmt.Printf("len=%d,cap=%d,%v \n", len(s1), cap(s1), s1)

	var s2 []int
	//截取
	s2 = s1[2:]
	fmt.Printf("len=%d,cap=%d,%v \n", len(s2), cap(s2), s2)

	s3 := make([]int, 3, 5)
	fmt.Printf("len=%d,cap=%d,%v \n", len(s3), cap(s3), s3)

	var s4 []int
	copy(s1, s4)
	fmt.Printf("len=%d,cap=%d,%v \n", len(s2), cap(s2), s2)

}
