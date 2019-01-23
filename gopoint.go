package main

import "fmt"

func main() {
	var p int
	p = 10
	changeValue(&p)
	fmt.Println(p)

	a, b := 3, 4
	swapValueByPoint(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
	c, d := swapValueBySwap(a, b)
	fmt.Println(c)
	fmt.Println(d)
}

func changeValue(p *int) {
	*p++
}

func swapValueByPoint(a, b *int) {
	*a, *b = *b, *a
}

func swapValueBySwap(a int, b int) (int, int) {
	return b, a
}
