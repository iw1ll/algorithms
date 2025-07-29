class ListNode {
    val: number;
    next: ListNode | null;
    constructor(val: number, next: ListNode | null) {
        this.val = val;
        this.next = next;
    }
}

class MyLinkedList {
    head: ListNode | null;
    size: number;
    constructor(head: ListNode | null = null, size: number = 0) {
        this.head = head;
        this.size = size;
    }

    get(index: number): number {
        if (index >= this.size || index < 0) return -1;
        let current = this.head;

        for (let i = 0; i < index; i++) {
            if (current !== null) {
                current = current.next
            }
        }

        return current ? current.val : -1;
    }

    addAtHead(val: number): void {
        this.addAtIndex(0, val);
    }

    addAtTail(val: number): void {
        this.addAtIndex(this.size, val);
    }

    addAtIndex(index: number, val: number): void {
        if (index > this.size || index < 0) {
            return;
        }

        if (index === 0) {
            const newNode = new ListNode(val, this.head)
            this.head = newNode;
            this.size++;
            return;
        }

        let current = this.head;

        for (let i = 0; i < index - 1; i++) {
            if (current === null ) break;
            current = current.next;
        }

        if (current === null ) return;

        const oldNode = current.next;
        current!.next = new ListNode(val, oldNode);
        this.size++;
    }

    deleteAtIndex(index: number): void {
        if (index >= this.size || index < 0) {
            return;
        }

         if (index === 0) {
            this.head = this.head!.next;
            this.size--;
            return;
        }

        let current = this.head;

        for (let i = 0; i < index - 1; i++) {
            if (current === null ) break;
            current = current.next;
        }

        current!.next = current!.next!.next;

        this.size--;
    }
}

export const testList = (): MyLinkedList => {
    const list = new MyLinkedList();
    list.addAtHead(1);
    list.addAtIndex(1, 2);
    list.addAtIndex(2, 3);
    list.deleteAtIndex(1)
    list.addAtTail(4);
    const param = list.get(0)
    // console.log(new ListNode(1, new ListNode(2, new ListNode(3, null))));
    console.log(param)
    return list;
}
