class Node {
    key: number;
    value: number;
    prev: Node | null;
    next: Node | null;

    constructor(key: number = 0, value: number = 0, prev: Node | null = null, next: Node | null = null) {
        this.key = key;
        this.value = value;
        this.prev = prev;
        this.next = next;
    }
}

export class LRUCache {
    capacity: number;
    data: Record<number, Node>;
    head: Node | null;
    tail: Node | null;

    constructor(capacity: number) {
        this.capacity = capacity;
        this.data = {};
        this.head = new Node();
        this.tail = new Node();
        this.head.next = this.tail;
        this.tail.prev = this.head;
    }

    remove(node: Node): void {
        if (node.prev) {
            node.prev.next = node.next;
        }
        if (node.next) {
            node.next.prev = node.prev;
        }
    }

    addToHead(node: Node): void {
        if (this.head?.next) {
            node.next = this.head.next;
        }
        node.prev = this.head;
        if (this.head && this.head.next) {
            this.head.next.prev = node;
        }
        if (this.head) {
            this.head.next = node;
        }
    }

    get(key: number): number {
        const node = this.data[key];

        if (!node) {
            return -1;
        }

        this.remove(node);
        this.addToHead(node);

        return node.value;
    }

    put(key: number, value: number): void {
        const node = this.data[key];

        if (node) {
            this.remove(node);
            this.addToHead(node);
            node.value = value;
            return;
        }

        if (Object.keys(this.data).length === this.capacity) {
            const tailNode = this.tail!.prev!;
            this.remove(tailNode);
            delete this.data[tailNode.key];
        }

        const newNode = new Node(key, value);
        this.data[key] = newNode;
        this.addToHead(newNode);
    }
}