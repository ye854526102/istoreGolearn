package main

import "fmt"

func main() {
	result := findLongNoReapeatString("abcdaaabcdef")
	fmt.Println(result)
}

func findLongNoReapeatString(s string) int {
	start, maxlength := 0, 0
	var lastImax int
	fmt.Printf("最长的字符串未%v\n", []byte(s))

	tmparr := make(map[byte]int)
	for i, v := range []byte(s) {
		fmt.Println(i)
		if lasrtI, ok := tmparr[v]; ok && lasrtI >= start {
			start = lasrtI + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
			lastImax = i
		}
		tmparr[v] = i
	}
	fmt.Printf("最长的字符串未%v\n", lastImax)
	return maxlength
}
