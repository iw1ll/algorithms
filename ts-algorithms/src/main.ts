import { deleteDuplicates, deleteMiddle, isPalindrome, middleNode, swapPairs } from './linked-list/linked-list-repeat.js';
import { ListNode } from './linked-list/middle-off-the-linked-list.js';
import { reverseList } from './linked-list/reverse-linked-list.js';
import { containsNearbyDuplicate, countGoodSubstrings, longestOnes, longestSubarray, minimumDifference, minimumRecolors, totalFruit } from './sliding-window/sliding-window.js';
import { containsDuplicate, isPalindromeStr, maxArea, removeDuplicates, reverseString, sortedSquares, strStr, twoSum, } from './two-pointers/two-pointers.js';
import { sum } from './utils/math.js';


console.log("Start algorithms");
/** Math */
sum(2, 3);

/** Two pointers */
containsDuplicate([1, 2, 3]);
isPalindromeStr("A man, a plan, a canal: Panama");
strStr('abc', 'c');
reverseString(["h","e","l","l","o"]);
twoSum([2, 3, 4], 6);
sortedSquares([-4 ,-1 ,0 ,3 , 10]);
removeDuplicates([0,0,1,1,1,2,2,3,3,4]);
maxArea([1,8,6,2,5,4,8,3,7]);
https://leetcode.com/problems/3sum/ ---> решаем на свежую голову

/** Sliding window  */
longestSubarray([1, 0, 1, 1, 1]);
longestOnes([0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1], 3);
containsNearbyDuplicate([1, 2, 3, 1], 3);
countGoodSubstrings('aababcabc');
minimumRecolors('WBWBBBW', 2);
minimumDifference([9, 4, 1, 7], 2);
totalFruit([1,2,3,2,2]);

/** Linked List */
const list = new ListNode(1, new ListNode(2, new ListNode(3 , null)));
middleNode(list);
reverseList(list);
deleteMiddle(list)
const palindromeLinkerList = new ListNode(1, new ListNode(2, new ListNode(2, new ListNode(1))));
isPalindrome(palindromeLinkerList)
const duplicates = new ListNode(1, new ListNode(1, new ListNode(2 , null)));
deleteDuplicates(duplicates);
const listPairs = new ListNode(1, new ListNode(2, new ListNode(3 , new ListNode(4))));
swapPairs(listPairs);

/** Hash Table */