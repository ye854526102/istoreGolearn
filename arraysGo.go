package main

import "fmt"

func varArrayDefind() {
	var arr1 [5]int
	arr1[1] = 1
	arr1[0] = 3
	arr1[3] = 4
	arr1[2] = 5
	arr1[4] = 7
	fmt.Println(arr1)

	//隐式声明
	arr2 := [5]int{5, 2, 3, 1, 5}
	fmt.Println(arr2)

	//不给长度声明,编译器自己确认长度
	var arr3 = [...]int{5, 4, 3, 2, 1}
	fmt.Println(arr3)

	//二位数组
	var arr4 [4][5]int
	fmt.Println(arr4)

	//遍历数组一
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for k, v := range arr2 {
		fmt.Printf("array_key is %d, array_value is %d\r\n", k, v)
	}
}

func alice_sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

func alice_max(arr []int) int {
	var max int
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
func alice_min(arr []int) int {
	var max int
	for _, v := range arr {
		if v < max {
			max = v
		}
	}
	fmt.Println(arr)
	return max
}
func change_array(arr *[5]int) error {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
	return nil
}
func change_slice(arr []int) error {
	arr[0] = 5
	return nil
}
func main() {
	varArrayDefind()
	fmt.Println(alice_sum([]int{3, 2, 1}))
	fmt.Println(alice_max([]int{5, 6, 1, 10, 1000}))
	fmt.Println(alice_min([]int{5, 6, 1, 10, 100}))
	var arr3 = [...]int{5, 4, 3, 2, 1}
	fmt.Println(change_array(&arr3))
	fmt.Println(change_slice(arr3[:]))
	fmt.Println(arr3)

	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr4)
	s1 := arr4[3:6]
	fmt.Println(s1)
	s2 := s1[:4]
	fmt.Println(s2)
	s3 := s1[:7]
	fmt.Println(s3)
}
