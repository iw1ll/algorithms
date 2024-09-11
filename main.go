package main

import (
	fn "algorithms/fibonacciNumber"
	po "algorithms/plusOne"
	ts "algorithms/twoSum"
	"fmt"
)

func main() {
	fmt.Println("Hello algorithm!")

	fmt.Println(ts.TwoSum([]int{3, 2, 4}, 6))

	fmt.Println(fn.Fib(5))

	po.PlusOne([]int{1})
}
