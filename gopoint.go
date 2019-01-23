package main

import "fmt"

func main() {
	var p int
	p = 10
	changeValue(&p)
	fmt.Println(p)

	a, b := 3, 4
	swapValue(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func changeValue(p *int) {
	*p++
}

func swapValue(a, b *int) {
	*a, *b = *b, *a
}
