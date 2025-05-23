export const longestSubarray = (nums: number[]): number => {
    let begin = 0;
    let windowsState = 0;
    let result = 0;

    for (let end = 0; end < nums.length; end++) {
        if (nums[end] == 0) {
            windowsState++;
        }

        while (windowsState > 1) {
            if (nums[begin] == 0) {
                windowsState--;
            }
            begin++;
        }
        let currentLength = end - begin;
        result = Math.max(result, currentLength);
    }
    return result;
};

export const longestOnes = (nums: number[], k: number): number => {
    let begin = 0
	let windowState = 0
	let result = 0

    for (let end = 0; end < nums.length; end++) {
        if (nums[end] === 0) {
            windowState++;
        }
        while (windowState > k) {
            if (nums[begin] === 0) {
                windowState--
            }
            begin++;
        }
        result = Math.max(result, end - begin + 1)
    }
    return result;
};

export const containsNearbyDuplicate = (nums: number[], k: number): boolean => {
    let begin  = 0;

    for (let end = 0; end < nums.length; end++) {
        for (let i = begin; i < end; i++) {
            if (nums[i] === nums[end]) {
                return true;
            }
    }
        if (end - begin >= k) {
                begin++;
        }
    }

    return false;
};