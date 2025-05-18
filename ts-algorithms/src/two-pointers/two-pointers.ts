export const containsDuplicate = (nums: number[]): boolean => {
    const sortNums = nums.sort((a, b) => a - b)
    for (let i = 0; i < nums.length; i++) {
        if (sortNums[i] === sortNums[i + 1]) {
            return true
        }
    }
    return false;
};

export const isPalindrome = (s: string): boolean => {
    const clean = s.toLowerCase().replace(/[^a-z]/g, "");

    let left = 0;
    let right = clean.length - 1;

    while (left < right) {
        if (clean[left] !== clean[right]) return false;
        left++;
        right--;
    }
    return true;
};

export const strStr = (haystack: string, needle: string): number => {
    if (needle.length === 0) return 0;

    const haystackLen = haystack.length;
    const needleLen = needle.length;

    const maxStartIndex = haystackLen - needleLen;

    for (let i = 0; i <= maxStartIndex; i++) {
        let j = 0;

        while (j < needleLen && haystack[i + j] === needle[j]) {
            j++;
        }

        if (j === needleLen) {
            return i;
        }
    }

    return -1;
};