export class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = val === undefined ? 0 : val;
        this.next = (next===undefined ? null : next)
    }
}

export const middleNode = (head: ListNode | null): ListNode | null => {
    let slow = head;
    let fast = head;

    while (fast !== null && fast.next !== null && slow !== null) {
        fast = fast.next.next;
        slow = slow.next;
    }

    return slow;
};
