package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)




func str2Int(strs []string) []int {
	ints := make([]int, 0, len(strs))

	for _, s := range strs {
		fmt.Println("s", s)
		s := strings.TrimSpace(s)
		fmt.Println("s", s)
		if s == "" {
			continue
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints = append(ints, v)
	}
	return ints
}

func main()  {
	fmt.Println("Enter comma-separated values:")

	reader := bufio.NewReader(os.Stdin)
	strsInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("problem reading user input %v\n", strsInput)
	}

	strs := strings.Split(strsInput, ",")
	
	values := str2Int(strs)
	fmt.Println(len(values),"-",values)
	if len(values) != 5 {
		fmt.Println("Only 5 comma- separated values are accepted.")
		return
	}

	total := 0
	for _, v := range values {
		if v <0 || v > 100{
			fmt.Println("Marks should be in range 0-100")
		}
		total += v
	}

	grade := (float64(total)/float64(500))*100

	fmt.Println("Grade:", grade)


}
