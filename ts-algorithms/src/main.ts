import { containsNearbyDuplicate, countGoodSubstrings, longestOnes, longestSubarray, minimumDifference, minimumRecolors } from './sliding-window/sliding-window.js';
import { containsDuplicate, isPalindrome, strStr } from './two-pointers/two-pointers.js';
import { sum } from './utils/math.js';

// math
sum(2, 3);

// Two pointers
containsDuplicate([1,2,3,]);
isPalindrome("aaaa");
strStr('abc', 'c');

// Sliding window
longestSubarray([1, 0, 1, 1, 1]);
longestOnes([0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], 3);
containsNearbyDuplicate([1,2,3,1], 3);
countGoodSubstrings('aababcabc');
minimumRecolors('WBWBBBW', 2);
console.log(minimumDifference([9,4,1,7], 2));