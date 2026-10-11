package main

import "fmt"


func main()  {
	num := 123

	var rev_num int
	for num != 0 {
		rev_num = 10 * rev_num + (num%10)
		num /= 10
	}

	fmt.Printf("Reverse number: %d\n", rev_num)

}
