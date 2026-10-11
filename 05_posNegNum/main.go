package main

import "fmt"


func main()  {
	num := 1

	if num < 0 {
		fmt.Println("Negative no. ", num)
	} else {
		fmt.Printf("Positive no. %d\n", num)
	}
}
