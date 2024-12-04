package hashmap

// Design a HashMap without using any built-in hash table libraries.

// Implement the MyHashMap class:

// MyHashMap() initializes the object with an empty map.
// void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
// int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
// void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.

// Example 1:
// ["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
// [[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
// Output
// [null, null, null, 1, -1, null, 1, null, -1]

// Explanation
// MyHashMap myHashMap = new MyHashMap();
// myHashMap.put(1, 1); // The map is now [[1,1]]
// myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
// myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
// myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
// myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
// myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
// myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
// myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]

// Constraints:
// 0 <= key, value <= 106
// At most 104 calls will be made to put, get, and remove.

type MyHashMap struct {
	n       int
	buckets []*LinkedList
}

type Node struct {
	Key   int
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Get(key int) int {
	current := ll.Head

	for current != nil {
		if current.Key == key {
			return current.Value
		}
		current = current.Next
	}
	return -1
}

func (ll *LinkedList) Put(key int, value int) {
	current := ll.Head

	for current != nil {
		if current.Key == key {
			current.Value = value
			return
		}
		current = current.Next
	}

	newNode := &Node{
		Value: value,
		Key:   key,
		Next:  ll.Head,
	}

	ll.Head = newNode
}

func (ll *LinkedList) Remove(key int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Key == key {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head

	for current.Next != nil {
		if current.Next.Key == key {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (mh *MyHashMap) NewMyHashMap() *MyHashMap {
	n := 991
	buckets := make([]*LinkedList, n)
	for i := range buckets {
		buckets[i] = &LinkedList{}
	}
	return &MyHashMap{
		n:       n,
		buckets: buckets,
	}
}

func (mh *MyHashMap) Hash(x int) int {
	return x / mh.n
}

func (mh *MyHashMap) Put(key int, value int) {
	n := mh.Hash(key)
	mh.buckets[n].Put(key, value)
}

func (mh *MyHashMap) Get(key int) int {
	n := mh.Hash(key)
	return mh.buckets[n].Get(key)
}

func (mh *MyHashMap) Remove(key int) {
	n := mh.Hash(key)
	mh.buckets[n].Remove(key)
}
