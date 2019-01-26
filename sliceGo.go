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

	fmt.Println("reslice")
	sliceAction()
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

//slice 的操作
func sliceAction() {
	//slice 再次slice
	arr := [...]string{"go语言", "php语言", "java语言", "js语言", "c语言", "c++语言"}

	s := arr[:5]

	fmt.Println(s)

	s_r_s := s[:2]

	fmt.Println(s_r_s)

	//slice的扩展

	s_p := arr[2:4]
	fmt.Println(s_p)

	//查出实体的arr,s_p[1:5]会报错,最多到[1,4]
	s_p_s := s_p[1:4]

	fmt.Println(s_p_s)

	fmt.Printf("%v--len(s)=%d--cap(s)=%d", s_p, len(s_p), cap(s_p))

	//如果利用append函数对slice添加元素,如果超出课cap的长度,go则会自动开辟一个新空间,生成一个新的数组

	a_s := arr[:len(arr)-2]
	fmt.Println(a_s)

	a_s_p1 := append(a_s, "python语言")
	fmt.Println(a_s_p1)

	a_s_p2 := append(a_s_p1, "ruby语言")
	fmt.Println(a_s_p2)
	fmt.Println("now arr is", arr)
	//这个不是对arr的视图
	a_s_p3 := append(a_s_p2, "R语言")
	fmt.Println(a_s_p3)
	fmt.Println("now arr is", arr)

}
