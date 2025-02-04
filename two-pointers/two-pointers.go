package twopointers

import (
	"slices"
	"unicode"
)

// Reverse String
// Write a function that reverses a string. The input string is given as an array of characters s.
// You must do this by modifying the input array in-place with O(1) extra memory.

// Example 1:
// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]

// Example 2:
// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]

func ReverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// Valid Palindrome

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

// Example 2:
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.

// Example 3:
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

func IsPalindrome(s string) bool {
	cleaned := make([]rune, 0, len(s))
	for _, char := range s {
		if isAlphanumeric(char) {
			cleaned = append(cleaned, unicode.ToLower(char))
		}
	}
	left, right := 0, len(cleaned)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
func isAlphanumeric(r rune) bool {
	return ('a' <= r && r <= 'z') || ('0' <= r && r <= '9')
}

// Two Sum II - Input Array Is Sorted
// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
// Your solution must use only constant extra space.

// Example 1:

// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

// Example 2:
// Input: numbers = [2,3,4], target = 6
// Output: [1,3]
// Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

// Example 3:
// Input: numbers = [-1,0], target = -1
// Output: [1,2]
// Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

func TwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		currentNum := numbers[left] + numbers[right]
		if currentNum == target {
			return []int{left + 1, right + 1}
		}
		if currentNum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

// Squares of a Sorted Array

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

// Example 1:

// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].

// Example 2:
// Input: nums = [-7,-3,2,3,11]
// Output: [4,9,9,49,121]

// Constraints:

// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums is sorted in non-decreasing order.

func SortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1

	for i := n - 1; i >= 0; i-- {
		rightSq := nums[right] * nums[right]
		leftSq := nums[left] * nums[left]
		if rightSq > leftSq {
			result[i] = rightSq
			right--
		} else {
			result[i] = leftSq
			left++
		}
	}

	return result
}

// 3Sum

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.

// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

// Example 2:
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// Example 3:
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.

// Constraints:
// 3 <= nums.length <= 3000
// -105 <= nums[i] <= 105

func ThreeSum(nums []int) [][]int {
	// slices.Sort(nums)
	// result := [][]int{}
	// n := len(nums)
	// for i := range n {
	// 	target := -nums[i]
	// 	left := i + 1
	// 	right := n - 1
	// 	for left < right {
	// 		currentSum := nums[left] + nums[right]
	// 		if currentSum == target {
	// 			result = append(result, []int{target, nums[left], nums[right]})
	// 		}
	// 		if currentSum > target {
	// 			right--
	// 		} else {
	// 			left++
	// 		}
	// 	}
	// }
	// return result
	// Сортируем массив по возрастанию
	slices.Sort(nums)

	// Инициализируем пустой слайс для хранения результатов
	result := [][]int{}

	// Сохраняем длину массива
	n := len(nums)

	// fmt.Println("Исходный массив:", nums)
	// fmt.Printf("Длина массива: %d\n", n)

	// Цикл по всем возможным первым числам тройки
	for i := 0; i < n-2; i++ {
		// fmt.Printf("Обрабатываем элемент под индексом %d (%d)\n", i, nums[i])

		// Проверяем, нет ли дубликата первого числа
		if i > 0 && nums[i] == nums[i-1] {
			// fmt.Println("Пропускаем дубликат первого числа")
			continue
		}

		// Определяем целевое значение для суммы трех чисел
		target := -nums[i]
		// fmt.Printf("Целевая сумма: %d (отрицание %d)\n", target, nums[i])

		// Инициализируем указатели для поиска второй и третьей цифры
		left, right := i+1, n-1

		// fmt.Printf("Левый указатель: %d, Правый указатель: %d\n", left, right)

		// Поиск тройки
		for left < right {
			currentSum := nums[left] + nums[right]

			// fmt.Printf("Текущая сумма: %d\n", currentSum)

			if currentSum == target {
				// fmt.Println("Найдена тройка!")

				result = append(result, []int{nums[i], nums[left], nums[right]})
				// fmt.Printf("Добавлена тройка: [%d, %d, %d]\n", nums[i], nums[left], nums[right])

				// Пропускаем дубликаты второго и третьего чисел
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// fmt.Printf("Обновленные указатели: Left=%d, Right=%d\n", left, right)

				left++  // Увеличиваем left для поиска новой пары
				right-- // Уменьшаем right для поиска нового третьего числа
			} else if currentSum < target {
				// fmt.Println("Текущая сумма меньше целевой, двигаемся влево")
				left++
			} else {
				// fmt.Println("Текущая сумма больше целевой, двигаемся вправо")
				right--
			}

			// fmt.Printf("После движения: Left=%d, Right=%d\n", left, right)
		}
	}

	// fmt.Println("\nИтоговый результат:")
	// for _, triplet := range result {
	// 	fmt.Println(triplet)
	// }

	return result // Возвращаем список всех уникальных тройек
	//  прорешать еще раз
}

// Remove Duplicates from Sorted Array

// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
// Return k.
// Custom Judge:
// The judge will test your solution with the following code:

// int[] nums = [...]; // Input array
// int[] expectedNums = [...]; // The expected answer with correct length
// int k = removeDuplicates(nums); // Calls your implementation
// assert k == expectedNums.length;
// for (int i = 0; i < k; i++) {
//     assert nums[i] == expectedNums[i];
// }
// If all assertions pass, then your solution will be accepted.

// Example 1:
// Input: nums = [1,1,2]
// Output: 2, nums = [1,2,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).

// Example 2:
// Input: nums = [0,0,1,1,1,2,2,3,3,4]
// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).

func RemoveDuplicates(nums []int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}

	return k + 1
}

// Move Zeroes
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.

// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Example 2:
// Input: nums = [0]
// Output: [0]

// Constraints:

// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1

func MoveZeroes(nums []int) {
	w := 0

	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[w] = nums[w], nums[i]
			w++
		}
	}
}
