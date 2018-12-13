package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"io/ioutil"
)

var (
	aa = 21
	ss = "golang"
	ff = 6.15183017521
	cc = 1 + 2i
)

func varZeroValue() {
	var a int
	var s string
	fmt.Printf("%d---%q", a, s)
}

func varInitialValue() {
	var a int = 18
	var s string = "weipeng"
	fmt.Println(a, s)
}
func varTypeDeduction() {
	var a = 18
	var s = "ye"
	var m, n = true, 3.14

	fmt.Println(a, s)
	fmt.Printf("%v-%T--%v-%T", m, m, n, n)
}
func varShorter() {
	a, s, m, n := 18, "weipeng", false, 63.013
	a = 20
	//notice1 函数内部:短定义 不能再 a:=20这样,否则会报错 // no new variables on left side of :=
	//notice2 :=短定义 不能再函数外面定义
	//a:=20

	fmt.Printf("%v-%T--%v-%T--%v-%T--%v-%T", a, a, s, s, m, m, n, n)
}
func varComplx() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Printf("%v---%T", c, c)
}
func euler() {
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func constDefind()  {
	const(
		a,b=3,4
	)
	var c int
	c = int(math.Sqrt((a*a + b*b)))
	fmt.Println(c)
}
func eunmsDefind()  {
	const(
		a = iota
		_
		aa
		bb
	)
	fmt.Println(a,aa,bb)
	//ioto作为种子,来表示 b,kb,mb,gb,tb,pb
	const(
		b = 1<<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b,kb,mb,gb,tb,pb)

}
func branchIf(){
	const filename="test.log"
	if contents,err :=ioutil.ReadFile(filename);err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s",contents)
	}
}
func main() {
	fmt.Println("next-func-is-varShorter()")
	varShorter()

	fmt.Println("\nnext-func-is-varZeroValue()")
	varZeroValue()

	fmt.Println("\nnext-func-is-varTypeDeduction()")
	varTypeDeduction()

	fmt.Println("\nnext-func-is-varInitialValue()")
	varInitialValue()

	fmt.Println("\nnext-func-is-varComplx()")
	varComplx()

	fmt.Println("\nnext-func-is-euler()")
	euler()

	fmt.Println("\nnext-func-is-triangle()")
	triangle()

	fmt.Println("\nnext-func-is-constDefind()")
	constDefind()

	fmt.Println("\nnext-func-is-constDefind()")
	eunmsDefind()

	fmt.Println("\nnext-func-is-branchIf()")
	branchIf()

	fmt.Println("\nnext-func-is-mian()")
	fmt.Println("Hello word")

	fmt.Printf("\n%v-%T--%v-%T--%v-%T--%v-%T", aa, aa, ss, ss, ff, ff, cc, cc)

}