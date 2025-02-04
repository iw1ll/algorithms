package main

import (
	f "algorithms/first"
	hashMap "algorithms/hash-map"
	ll "algorithms/linked-list"
	sw "algorithms/sliding-window"
	tp "algorithms/two-pointers"

	"fmt"
)

func main() {
	fmt.Println("Hello algorithm! MinSubArrayLen")
	moveZeroesSlice := []int{0, 1, 0, 3, 12}

	// Start
	f.TwoSum([]int{3, 2, 4}, 6)
	f.Fib(5)
	f.PlusOne([]int{4, 3, 2, 1})

	//Two Pointers
	reverseStr := []byte{'h', 'e', 'l', 'l', 'o'}
	tp.ReverseString(reverseStr)
	// fmt.Println"Reverse string: %s\n", string(reverseStr)
	tp.IsPalindrome("A man, a plan, a canal: Panama")
	tp.TwoSum([]int{1, 2, 4}, 6)
	tp.ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	tp.SortedSquares([]int{-4, -1, 0, 3, 10})
	tp.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	tp.RemoveDuplicates([]int{0, 1, 0, 3, 12})
	tp.MoveZeroes(moveZeroesSlice)

	// Sliding Window
	sw.FindMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
	sw.MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})

	// LinkedList
	ll.Init()
	designLinkedList := ll.Constructor()
	designLinkedList.AddAtHead(1)
	designLinkedList.AddAtTail(3)
	designLinkedList.AddAtIndex(1, 2)
	designLinkedList.Get(1)
	designLinkedList.DeleteAtIndex(1)
	designLinkedList.Get(1)
	designLinkedList.PrintList()

	// LinkedList Middle of the Linked List
	values := []int{1, 2, 3, 4, 5}
	head := ll.CreateLinkedList(values)
	middle := ll.MiddleNode(head)
	if middle != nil {
		fmt.Printf("Middle: %d\n", middle.Val)
	} else {
		fmt.Println("Empty")
	}

	// Hash Map
	h := &hashMap.MyHashMap{}
	myHashMap := h.NewMyHashMap()

	myHashMap.Put(1, 1)
	myHashMap.Put(2, 2)

	myHashMap.Get(1)
	myHashMap.Get(2)
	myHashMap.Get(3)

	myHashMap.Remove(2)
	myHashMap.Get(2)

}
