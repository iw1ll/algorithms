package main

import (
	f "algorithms/first"
	linkedList "algorithms/linkedList"
	sw "algorithms/slidingWindow"
	tp "algorithms/twoPointers"

	"fmt"
)

func main() {
	fmt.Println("Hello algorithm! MinSubArrayLen")

	// First
	f.TwoSum([]int{3, 2, 4}, 6)
	f.Fib(5)
	f.PlusOne([]int{4, 3, 2, 1})

	//Two Pointers
	reverseStr := []byte{'h', 'e', 'l', 'l', 'o'}
	tp.ReverseString(reverseStr)
	// fmt.Printf("Reverse string: %s\n", string(reverseStr))
	tp.IsPalindrome("A man, a plan, a canal: Panama")
	tp.TwoSum([]int{1, 2, 4}, 6)
	tp.ThreeSum([]int{-1, 0, 1, 2, -1, -4})

	// Sliding Window
	sw.FindMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
	sw.MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})

	// LinkedList
	linkedList.Init()
	designLinkedList := linkedList.Constructor()
	designLinkedList.AddAtHead(1)
	designLinkedList.AddAtTail(3)
	designLinkedList.AddAtIndex(1, 2)
	designLinkedList.Get(1)
	designLinkedList.DeleteAtIndex(1)
	designLinkedList.Get(1)
	designLinkedList.PrintList()

	// LinkedList Middle of the Linked List
	values := []int{1, 2, 3, 4, 5}
	head := linkedList.CreateLinkedList(values)
	middle := linkedList.MiddleNode(head)
	if middle != nil {
		fmt.Printf("Middle: %d\n", middle.Val)
	} else {
		fmt.Println("Empty!")
	}
}
