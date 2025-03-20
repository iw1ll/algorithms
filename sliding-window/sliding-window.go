package slidingwindow

import "math"

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
	// how much 0?
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
