package slidingwindow

import (
	"math"
)

// Maximum Average Subarray I
// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.

// Example 1:
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

// Example 2:
// Input: nums = [5], k = 1
// Output: 5.00000
// Constraints:
// n == nums.length
// 1 <= k <= n <= 105
// -104 <= nums[i] <= 104

func FindMaxAverage(nums []int, k int) float64 {
	begin := 0
	windowState := 0
	var result float64 = math.MinInt64

	for end := range nums {
		windowState += nums[end]
		if end-begin+1 == k {
			result = math.Max(float64(result), float64(windowState))
			windowState -= nums[begin]
			begin += 1
		}
	}

	return result / float64(k)
}

// Minimum Size Subarray Sum

// Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// Example 1:
// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.

// Example 2:
// Input: target = 4, nums = [1,4,4]
// Output: 1

// Example 3:
// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
// Output: 0

// Constraints:
// 1 <= target <= 109
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 104

func MinSubArrayLen(target int, nums []int) int {
	begin := 0
	windowState := 0
	result := math.MaxInt64

	for end := range nums {
		windowState += nums[end]
		for windowState >= target {
			windowSize := end - begin + 1
			result = min(result, windowSize)
			windowState -= nums[begin]
			begin++
		}
	}
	if result == math.MaxInt64 {
		return 0
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func LongestOnes(nums []int, k int) int {
	begin := 0
	// how much 0?
	windowState := 0
	result := 0

	for end := range nums {
		if nums[end] == 0 {
			windowState++
		}
		for windowState > k {
			if nums[begin] == 0 {
				windowState--
			}
			begin++
		}
		currentLength := end - begin + 1
		result = max(result, currentLength)
	}
	return result
}

func LongestSubarray(nums []int) int {
	begin := 0
	windowState := 0
	result := 0

	for end := range nums {
		if nums[end] == 0 {
			windowState++
		}
		for windowState > 1 {
			if nums[begin] == 0 {
				windowState--
			}
			begin++
		}
		currentLength := end - begin
		result = max(result, currentLength)
	}
	return result
}

// Fruit Into Baskets

// You are visiting a farm that has a single row of fruit trees arranged from left to right. The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.

// You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

// You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
// Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
// Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
// Given the integer array fruits, return the maximum number of fruits you can pick.

// Example 1:

// Input: fruits = [1,2,1]
// Output: 3
// Explanation: We can pick from all 3 trees.
// Example 2:

// Input: fruits = [0,1,2,2]
// Output: 3
// Explanation: We can pick from trees [1,2,2].
// If we had started at the first tree, we would only pick from trees [0,1].
// Example 3:

// Input: fruits = [1,2,3,2,2]
// Output: 4
// Explanation: We can pick from trees [2,3,2,2].
// If we had started at the first tree, we would only pick from trees [1,2].

func TotalFruit(fruits []int) int {
	begin := 0
	windowState := make(map[int]int)
	result := 0

	for end := range fruits {
		windowState[fruits[end]]++
		for len(windowState) > 2 {
			windowState[fruits[begin]]--
			if windowState[fruits[begin]] == 0 {
				delete(windowState, fruits[begin])
			}
			begin++
		}
		currentLength := end - begin + 1
		result = max(result, currentLength)
	}
	return result
}

// t:41:22
