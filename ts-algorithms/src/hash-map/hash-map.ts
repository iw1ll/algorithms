// class MyHashMap {
// 	n:number;
// 	buckets: LinkedList[];

//     constructor(n:number, buckets: LinkedList[]) {
//         this.n = n;
//         this.buckets = buckets;
//     }
// }

class ListNode {
	key: number;
	value: number;
	next: ListNode | null;

    constructor(key:number, value: number, next: ListNode | null = null) {
        this.key = key;
        this.value = value;
        this.next = next
    }
}

class LinkedList {
	head: ListNode | null;

    constructor(head: ListNode | null = null) {
        this.head = head;
    }

    get(key: number): number {
        let current: ListNode | null = this.head;

        while (current) {
            if (current.key === key) {
                return current.value;
            }
            current = current.next;
        }
        return -1;
    }

    put(key: number, value: number): void {
        let current = this.head;

        while (current) {
            if (current.key === key) {
                current.value = value;
                return;
            }
            current = current.next;
        }

        let newNode = new ListNode(key, value, this.head);
        this.head = newNode; 
    }

    remove(key: number): void {
         
    }
}

class MyHashMap {
	n:number;
	buckets: LinkedList[];

    constructor(n:number, buckets: LinkedList[]) {
        this.n = n;
        this.buckets = buckets;
    }

    put(key: number, value: number): void {
        
    }

    get(key: number): number {
        return -1;
    }

    remove(key: number): void {
        
    }
}