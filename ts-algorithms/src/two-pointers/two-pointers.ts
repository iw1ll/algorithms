export const containsDuplicate = (nums: number[]): boolean => {
    const sortNums = nums.sort((a, b) => a - b)
    for (let i = 0; i < nums.length; i++) {
        if (sortNums[i] === sortNums[i + 1]) {
            return true
        }
    }
    return false;
};