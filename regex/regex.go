package main

import (
	"fmt"
	"regexp"
)

const name = "qq18018737358.w 854526102@qq.com yewiepeng@suntek.com.cn wishmyhzizi@163.com"

func main() {
	//regexpObj,_:= regexp.Compile(`[0-9]+\.[a-z]`)
	//res:=regexpObj.FindString(name)
	//fmt.Println(res)

	//regexpObj,_:= regexp.Compile(`([a-zA-Z0-9])+@([a-zA-Z0-9.]+)`)
	//res:=regexpObj.FindAllString(name,-1)
	//fmt.Println(res)

	regexpObj, _ := regexp.Compile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)`)
	res := regexpObj.FindAllStringSubmatch(name, -1)
	for _, value := range res {
		fmt.Println(value[0])
		fmt.Println(value[1])
		fmt.Println(value[2])
	}

}
