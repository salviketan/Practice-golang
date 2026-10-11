package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter digit (0-9):")

	if scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		input, err := strconv.Atoi(text)
		fmt.Printf("input type %T\n", input)
		if err != nil || input < 0 || input > 9 {
			fmt.Println("please enter a single digit 0-9")
			return
		}
		n := input
		nn := 10*n+n
		nnn := 10*nn+n
		result := n + nn + nnn
		fmt.Println("1 result: ",result)
	
		// Loop approach
		result = 0
		var i int = 0
		var value int
		for i < 3 {
			value = 10 * value + input
			result += value
			i++
		}
		fmt.Println("2 result: ",result)
	}


}
