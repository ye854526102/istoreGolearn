package main

import "fmt"

func main() {

	createSlice()
	copySlice()

}
func copySlice() {
	//
	s := []int{3, 4, 5, 6, 7, 8, 9}

	//smc 里面的值根据len来看取多少
	s_m_c := make([]int, 16, 32)
	copy(s_m_c, s)
	printSlice(s_m_c)

	//slice删除中间,主要通过append,和[]int...不可定参数传递
	s_m_c_n := append(s_m_c[:6], s_m_c[7:]...)
	printSlice(s_m_c_n)

	//slic删除头部的数据,头部拿去会是cap减少,因为slice不可向前扩展
	s_m_c_h := s[1:]
	printSlice(s_m_c_h)

	//slic删除头部的数据
	s_m_c_t := s[:len(s)-1]
	printSlice(s_m_c_t)

}

func createSlice() {
	//声明方式1
	var s []int //zorevalue is nil

	fmt.Println(s == nil)

	//隐式声明,可以说是对数组2,3,4的视图
	s_s := []int{2, 3, 4}
	printSlice(s_s)

	//make内建函数声明
	s_m := make([]int, 10)
	printSlice(s_m)

	s_m_c := make([]int, 6, 32)
	printSlice(s_m_c)
}

func printSlice(s []int) {
	fmt.Printf("value is=%v--len(%d)--cap(%d)\n", s, len(s), cap(s))
}
