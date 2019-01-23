package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	q, r := div(1, 2)
	fmt.Println(q)
	fmt.Println(r)
	if result, eror := calc(q, r, "+x"); eror != nil {
		fmt.Println(eror)
	} else {
		fmt.Println(result)
	}
	result2 := apply(func(i int, i2 int) int {
		return i + i2
	}, 4, 5)
	fmt.Println(result2)

	result3 := apply(pow, 2, 3)
	fmt.Println(result3)

	result4 := sum(2, 3)
	fmt.Println(result4)
}
func div(a, b int) (q, r int) {
	return q + 1, r + 1
}

//可变参数函数
func sum(number ...int) int {
	i := 0
	for k := range number {
		i += number[k]
	}
	for _, v := range number {
		i += v
	}
	return i
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func calc(a, b int, expr string) (int, error) {
	var result int
	var err error
	switch expr {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		result = 0
		err = fmt.Errorf("unknow you input expr %s", expr)
	}
	return result, err
}

//参数是函数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opme := runtime.FuncForPC(p).Name()
	fmt.Println(opme)
	return op(a, b)
}
