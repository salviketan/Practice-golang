package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Average(nums []int) (float64, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("cannot average an empty list")
	}

	var total int
	for _, v := range nums {
		total += v
	}
	return  float64(total)/float64(len(nums)), nil
}


func Str2ints(strs []string) []int {
	ints := make([]int, 0, len(strs))

	for _, s := range strs{
		s := strings.TrimSpace(s)
		if s == "" {continue}
		v, err := strconv.Atoi(s)
		if err != nil {panic(err)}
		ints = append(ints, v)
	}

	return ints
}

func main()  {
	fmt.Println("Enter comma-separated values:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)

	strs := strings.Split(input, ",")

	values := Str2ints(strs)

	result, err := Average(values)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %.2f\n", result)
}