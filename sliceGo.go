package main

import (
	"fmt"
)

func main() {

	//数据
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//array to slice
	printArray(arr[:])

	//slice 也说是对数组的是视图
	fmt.Println("arr[2:5]", arr[2:5])

	s1 := arr[:]
	fmt.Println("arr[:] is =", s1)
	s2 := arr[:3]
	fmt.Println("arr[:] is =", s2)
	fmt.Println("arr is =", arr)

	//更新
	ss1 := arr[:]
	updateSlice(ss1)
	fmt.Println("arr[:] is =", ss1)
	ss2 := arr[:3]
	updateSlice(ss2)
	fmt.Println("arr[:] is =", ss2)
	fmt.Println("arr is =", arr)
}

func updateSlice(s []int) {
	s[0] = 100
}

func printArray(s []int) {
	s[1] = 200
	for _, v := range s {
		fmt.Println(v)
	}
}
