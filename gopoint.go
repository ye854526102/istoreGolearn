package main

import "fmt"

func main() {
	var p int
	p = 10
	changeValue(&p)
	fmt.Println(p)
}

func changeValue(p *int) {
	*p++
}
