package main

import (
	f "algorithms/first"
	tp "algorithms/twoPointers"

	"fmt"
)

func main() {
	fmt.Println("Hello algorithm!")

	// First
	fmt.Println(f.TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(f.Fib(5))
	fmt.Println(f.PlusOne([]int{4, 3, 2, 1}))

	//TwoPointers
	check := []byte{'h', 'e', 'l', 'l', 'o'}
	tp.ReverseString(check)
	fmt.Println(string(check))
}
