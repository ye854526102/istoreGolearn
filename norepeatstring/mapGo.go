package main

import "fmt"

func main() {
	result := findLongNoReapeatString("abcdaaabcdef")
	results := findLongNoReapeatString("我熬我看我我的奇偶奇偶就是ni")
	resultss := findLongNoReapeatString("小猪佩奇")
	fmt.Println(result)
	fmt.Println(results)
	fmt.Println(resultss)
}

func findLongNoReapeatString(s string) int {
	start, maxlength := 0, 0
	tmparr := make(map[rune]int)
	for i, v := range []rune(s) {
		if lasrtI, ok := tmparr[v]; ok && lasrtI >= start {
			start = lasrtI + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		tmparr[v] = i
	}
	return maxlength
}
