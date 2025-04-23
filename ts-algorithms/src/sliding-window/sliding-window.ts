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