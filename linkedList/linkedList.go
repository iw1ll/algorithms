package linkedlist

import "fmt"

// Design Linked List

// Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
// A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next is a pointer/reference to the next node.
// If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.
// Implement the MyLinkedList class:
// MyLinkedList() Initializes the MyLinkedList object.
// int get(int index) Get the value of the indexth node in the linked list. If the index is invalid, return -1.
// void addAtHead(int val) Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
// void addAtTail(int val) Append a node of value val as the last element of the linked list.
// void addAtIndex(int index, int val) Add a node of value val before the indexth node in the linked list. If index equals the length of the linked list, the node will be appended to the end of the linked list. If index is greater than the length, the node will not be inserted.
// void deleteAtIndex(int index) Delete the indexth node in the linked list, if the index is valid.

// Example 1:

// Input
// ["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
// [[], [1], [3], [1, 2], [1], [1], [1]]
// Output
// [null, null, null, null, 2, null, 3]

// Explanation
// MyLinkedList myLinkedList = new MyLinkedList();
// myLinkedList.addAtHead(1);
// myLinkedList.addAtTail(3);
// myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
// myLinkedList.get(1);              // return 2
// myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
// myLinkedList.get(1);              // return 3

type Node struct {
	Value int
	Next  *Node
}

type MyLinkedList struct {
	Head *Node
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (list *MyLinkedList) Get(index int) int {
	if index < 0 || index >= list.Size {
		return -1
	}

	current := list.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (list *MyLinkedList) AddAtHead(val int) {
	list.AddAtIndex(0, val)
}

func (list *MyLinkedList) AddAtTail(val int) {
	list.AddAtIndex(list.Size, val)
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > list.Size {
		return
	}

	list.Size++

	if index == 0 {
		newNode := &Node{
			Value: val,
			Next:  list.Head,
		}
		list.Head = newNode
		return
	}

	current := list.Head

	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	oldNext := current.Next
	current.Next = &Node{
		Value: val,
		Next:  oldNext,
	}

}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= list.Size {
		return
	}

	list.Size--

	if index == 0 {
		list.Head = list.Head.Next
		return
	}

	current := list.Head

	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
}

func Init() {
	fmt.Println("init")
}

func (list *MyLinkedList) PrintList() {
	current := list.Head
	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
