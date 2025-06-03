package slidingwindow

import (
	"math"
	"sort"
)

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

func MinimumRecolors(blocks string, k int) int {
	start := 0
	count := 0
	result := len(blocks)

	for end := range blocks {
		if blocks[end] == 'W' {
			count++
		}

		if end-start+1 == k {
			result = min(result, count)
			if blocks[start] == 'W' {
				count--
			}
			start++
		}
	}
	return result
}

func MinimumDifference(nums []int, k int) int {
	sort.Ints(nums)

	result := math.MaxInt

	for i := 0; i <= len(nums)-k; i++ {
		a := nums[i+k-1] - nums[i]

		result = min(a, result)
	}

	return result
}
