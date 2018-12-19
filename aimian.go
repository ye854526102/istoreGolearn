package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * AI核心代码,估值一个亿
 */
func aiMain() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("I am the AI:")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("There ware errors reading,exiting program.")
		}
		input = strings.Replace(input, "吗", "", -1)
		input = strings.Replace(input, "?", "!", -1)
		input = strings.Replace(input, "? ", "!", -1)
		fmt.Printf("%s\n", input)
	}

}
func main() {
	aiMain()
}
