class MyListNode {
    key: number;
    value: number;
    next: MyListNode | null;

    constructor(key: number, value: number, next: MyListNode | null = null) {
        this.key = key;
        this.value = value;
        this.next = next;
    }
}

class MyLinkedList {
    head: MyListNode | null;

    constructor(head: MyListNode | null = null) {
        this.head = head;
    }

    get(key: number): number {
        let current: MyListNode | null = this.head;

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

        const newNode = new MyListNode(key, value, this.head);
        this.head = newNode;
    }

    remove(key: number): void {
        if (this.head === null) {
            return;
        }

        if (this.head.key === key) {
            this.head = this.head.next;
            return;
        }

        let current = this.head;

        while (current.next) {
            if (current.next.key === key) {
                current.next = current.next.next;
                return;
            }
            current = current.next;
        }
    }
}

export class MyHashMap {
    n: number;
    buckets: MyLinkedList[];

    constructor(n: number = 991) {
        this.n = n;
        this.buckets = new Array<MyLinkedList>(this.n);
        for (let i = 0; i < this.n; i++) {
            this.buckets[i] = new MyLinkedList();
        }
    }

    hash(x: number): number {
        return x % this.n;
    }

    put(key: number, value: number): void {
        const n = this.hash(key);
        this.buckets[n].put(key, value);
    }

    get(key: number): number {
        const n = this.hash(key);
        return this.buckets[n].get(key);
    }

    remove(key: number): void {
        const n = this.hash(key);
        this.buckets[n].remove(key);
    }
}