export class ListNode {
    val: number;
    next: ListNode | null;
    constructor(val: number, next: ListNode | null) {
        this.val = val;
        this.next = next;
    }
}

export class MyLinkedList {
    head: ListNode | null;
    size: number;

    constructor(head: ListNode | null = null, size: number = 0) {
        this.head = head;
        this.size = size;
    }

    get(index: number): number {
        if (index >= this.size || index < 0) {
            return -1;
        }

        let current = this.head;

        for (let i = 0; i < index; i++) {
            if (current === null) {
                break;
            }
            current = current.next;
        }
        return current ? current.val : - 1;
    }

    addAtIndex(index: number, val: number) {
        if (index > this.size || index < 0) {
            return -1;
        }

        let current = this.head;

        if (index === 0) {
            const newNode = new ListNode(val, this.head);
            this.head = newNode;
            this.size++;
            return
        }

        for (let i = 0; i < index - 1; i++) {
            if (current === null || current.next === null) {
                break;
            }
            current = current.next;
        }

        if (current) {
            const oldNode = current.next;
            current.next = new ListNode(val, oldNode);
            this.size++;
        }
    }

    addAtHead(val: number): void {
        this.addAtIndex(0, val);
    }

    addAtTail(val: number): void {
        this.addAtIndex(this.size, val);
    }

    deleteAtIndex(index: number): void {
        if (index >= this.size || index < 0) {
            return;
        }

        if (index === 0) {
            if(this.head !== null) {
                this.head = this.head?.next;
                this.size--;
            }
            return;
        }

        let current = this.head;

        for (let i = 0; i < index - 1; i++) {
            if (current === null || current.next === null) {
                break;
            }
            current = current.next;
        }

        if (current !== null && current.next !== null) {
            current.next = current.next.next;
        }

        this.size--;
    }
}

export const middleNode = (head: ListNode | null): ListNode | null => {
    let slow = head;
    let fast = head;

    while (fast && fast.next) {
        fast = fast.next.next;
        slow = slow!.next;
    }

    return slow;
};

export const deleteMiddle = (head: ListNode | null): ListNode | null => {
    if (head === null || head.next === null) {
        return null;
    }

    let slow: ListNode | null = head;
    let fast = head?.next?.next;

    while (slow && fast && fast.next) {
        fast = fast.next.next;
        slow = slow!.next;
    }

    if (slow && slow.next) {
        slow!.next = slow!.next!.next;
    }

    return head;
};



function isPalindrome(head: ListNode | null): boolean {
    function middle(head: ListNode | null): ListNode | null {
        let slow = head;
        let fast = head;

        while (slow && fast && fast.next) {
            fast = fast.next.next;
            slow = slow.next;
        }
        return slow;
    }
    
    function reverse(head: ListNode | null): ListNode | null {
        let current = head;
        let prev: ListNode | null = null;

        while (current) {
            let tmp: ListNode | null = current.next;
            current.next = prev;
            prev = current;
            current = tmp;
        }
        return prev;
    }
    
    let mid = middle(head);
    let second = reverse(mid)
    let first = head;
     
    while (first && second) {
        if (first.val !== second.val) {
            return false;
        } else {
            first = first.next;
            second = second.next;
        }
    }
    return true;
};