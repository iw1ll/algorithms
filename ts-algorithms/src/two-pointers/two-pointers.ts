export const containsDuplicate = (nums: number[]): boolean => {
    const sortNums = nums.sort((a, b) => a - b);

    for (let i = 0; i < nums.length; i++) {
        if (sortNums[i] === sortNums[i + 1]) {
            return true
        }
    }

    return false;
};

export const isPalindromeStr = (s: string): boolean => {
    const clean = s.toLowerCase().replace(/[^a-z0-9]/g, '');
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

export const reverseString = (s: string[]): void => {
    let left = 0;
    let right = s.length - 1;

    while (left < right) {
        const temp = s[left];
        s[left] = s[right];
        s[right] = temp;
        left++;
        right--;
    }
};

export const twoSum = (numbers: number[], target: number): number[] => {
    let left = 0;
    let right = numbers.length - 1;

    while (left < right) {
        const currentSum = numbers[left] + numbers[right];

        if (currentSum === target) {
            return [left + 1, right + 1]
        } else if (currentSum > target) {
            right--;
        } else {
            left++;
        }
    }

    return [-1];
}

export const sortedSquares = (nums: number[]): number[] => {
    const result = [];
    let left = 0;
    let right = nums.length - 1;

    for (let i = nums.length - 1; i >= 0; i--) {
        const rightSqrt = nums[right] * nums[right]; 
        const leftSqrt = nums[left] * nums[left]; 

        if (rightSqrt > leftSqrt) {
            result[i] = rightSqrt;
            right--;
        } else {
            result[i] = leftSqrt;
            left++;
        }
    }

    return result;
};

export const removeDuplicates = (nums: number[]): number => {
    let k = 0;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] !== nums[k]) {
            k++;
            nums[k] = nums[i];
        }
    }

    return k + 1;
};

export const tp = (nums: number[], target: number): number[] => {
    for(let i = 0; i < nums.length; i++) {
        for (let j = i + 1; j < nums.length; j++) {
            if (nums[i] + nums[j] === target) {
                return [i, j];
            }
        }
    }

    return []; 
};

export const maxArea = (height: number[]): number => {
    let left = 0;
    let right = height.length - 1;
    let maxWater = 0;

    while (left < right) {
        const width = right - left;
        const minHeight = Math.min(height[left], height[right]);
        const water = minHeight * width;
        maxWater = Math.max(maxWater, water);

        if (height[left] < height[right]) {
            left++;
        } else {
            right--;
        }
    }

    return maxWater;
};

function threeSum(nums: number[]): number[][] {
    const result: number[][] = [];
    
    // 1. СОРТИРУЕМ МАССИВ
    // Без сортировки нельзя использовать два поинтера
    nums.sort((a, b) => a - b);
    
    // 2. ФИКСИРУЕМ ПЕРВОЕ ЧИСЛО
    for (let i = 0; i < nums.length - 2; i++) {
        
        // ПРОПУСК ДУБЛИКАТОВ для первого числа
        // Если nums[i] такой же как предыдущий — эту тройку уже нашли
        if (i > 0 && nums[i] === nums[i - 1]) {
            continue;
        }
        
        // ОПТИМИЗАЦИЯ: если первое число > 0,
        // то сумма трёх положительных чисел не может быть 0
        if (nums[i] > 0) {
            break;
        }
        
        // 3. ДВА ПОИНТЕРА ДЛЯ ОСТАВШИХСЯ ДВУХ ЧИСЕЛ
        let left = i + 1;
        let right = nums.length - 1;
        
        while (left < right) {
            const sum = nums[i] + nums[left] + nums[right];
            
            if (sum === 0) {
                // НАШЛИ ТРОЙКУ!
                result.push([nums[i], nums[left], nums[right]]);
                
                // ПРОПУСКАЕМ ДУБЛИКАТЫ слева
                while (left < right && nums[left] === nums[left + 1]) {
                    left++;
                }
                
                // ПРОПУСКАЕМ ДУБЛИКАТЫ справа
                while (left < right && nums[right] === nums[right - 1]) {
                    right--;
                }
                
                // Двигаем оба указателя
                left++;
                right--;
                
            } else if (sum < 0) {
                // Сумма меньше нуля → нужно увеличить
                // Двигаем left вправо (к бОльшим числам)
                left++;
            } else {
                // Сумма больше нуля → нужно уменьшить
                // Двигаем right влево (к меньшим числам)
                right--;
            }
        }
    }
    
    return result;
}

export const isSubsequence = (s:string, t: string): boolean => {
    let sNum = 0;

    for (let i = 0; i < t.length; i++) {
        if (t[i] === s[sNum]) {
            sNum++;
        }
    }
    if (s.length === sNum) {
        return true;
    }

    return false;
};

export const moveZeroes = (nums: number[]): void => {
    let left = 0;

    for(let right = 0; right < nums.length; right++) {
        if (nums[right] !== 0) {
            [nums[left], nums[right]] = [nums[right], nums[left]];
            left++;
        }
    }
};

export const merge = (nums1: number[], m: number, nums2: number[], n: number): void => {
    let p1 = m - 1;
    let p2 = n - 1;
    let p = m + n - 1;

    while (p2 >= 0) {
        if (p1 >= 0 && nums1[p1] > nums2[p2]) {
            nums1[p] = nums1[p1];
            p1--;
        } else {
            nums1[p] = nums2[p2];
            p2--;
        }
        p--;
    }
};