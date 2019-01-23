package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
	"os"
	"strconv"
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
	var a = 18
	var s = "weipeng"
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
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func constDefind() {
	const (
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}
func eunmsDefind() {
	const (
		a = iota
		_
		aa
		bb
	)
	fmt.Println(a, aa, bb)
	//ioto作为种子,来表示 b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)

}
func branchIf() {
	const filename = "test.log"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", contents)
	}
}
func switchDefind() {

}
func switchDemoGrade(score int) string {
	s := ""
	switch {
	case score == 0 || score > 100:
		panic(fmt.Sprintf("you score is wrong %d", score))
	case score < 60:
		s = "F"
	case score < 80:
		s = "C"
	case score < 90:
		s = "B"
	case score < 100:
		s = "A"
	}
	return s
}
func forDefind() {

}

//已知二进制转换为x%2
func forDemoConvertToBin(number int) string {
	res := ""
	for ; number > 0; number /= 2 {
		lsb := number % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

//TODO 转8进制
func forDemoConvertToOCT(number int) string {
	res := ""
	return res
}

//TODO 16进制
func forDemoConvertToHex(number int) string {
	res := ""
	return res
}

//TODO 32进制?
func forDemoConvertToMore(number int) string {
	res := ""
	return res
}
func printFileLine(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("open file %s is wrong", file))
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
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

	fmt.Println("\nnext-func-is-eunmsDefind()")
	eunmsDefind()

	fmt.Println("\nnext-func-is-branchIf()")
	branchIf()

	fmt.Println("\nnext-func-is-switchDemoGrade()")
	fmt.Println(switchDemoGrade(60))
	fmt.Println(switchDemoGrade(59))
	fmt.Println(switchDemoGrade(69))
	fmt.Println(switchDemoGrade(89))
	fmt.Println(switchDemoGrade(100))

	fmt.Println("\nnext-func-is-forDemoConvertToBin()")
	fmt.Println(forDemoConvertToBin(100))
	fmt.Println(forDemoConvertToBin(8))

	fmt.Println("\nnext-func-is-printFileLine()")
	printFileLine("test.log")

	fmt.Println("\nnext-func-is-mian()")
	fmt.Println("Hello word")

	fmt.Printf("\n%v-%T--%v-%T--%v-%T--%v-%T", aa, aa, ss, ss, ff, ff, cc, cc)

}
