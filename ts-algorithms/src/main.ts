import { testList } from './linked-list/design-linked-list.js';
import { ListNode, middleNode } from './linked-list/middle-off-the-linked-list.js';
import { reverseList } from './linked-list/reverse-linked-list.js';
import { containsNearbyDuplicate, countGoodSubstrings, longestOnes, longestSubarray, minimumDifference, minimumRecolors } from './sliding-window/sliding-window.js';
import { containsDuplicate, isPalindrome, reverseString, sortedSquares, strStr, twoSum } from './two-pointers/two-pointers.js';
import { sum } from './utils/math.js';

/** Math */
sum(2, 3);

/** Two pointers */
containsDuplicate([1, 2, 3]);
isPalindrome("A man, a plan, a canal: Panama");
strStr('abc', 'c');
reverseString(["h","e","l","l","o"]);
twoSum([2, 3, 4], 6);
sortedSquares([-4 ,-1 ,0 ,3 , 10]);

/** Sliding window  */
longestSubarray([1, 0, 1, 1, 1]);
longestOnes([0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1], 3);
containsNearbyDuplicate([1, 2, 3, 1], 3);
countGoodSubstrings('aababcabc');
minimumRecolors('WBWBBBW', 2);
minimumDifference([9, 4, 1, 7], 2);

/** Linked List */
const list = new ListNode(1, new ListNode(2, new ListNode(3, null)));
const middleLinkedList = middleNode(list);
const reverseLinkedList = reverseList(list);
testList();
