package main

import "fmt"

func main() {
	q, r := div(1, 2)
	fmt.Println(q)
	fmt.Println(r)

}
func div(a, b int) (q, r int) {
	return q + 1, r + 1
}
