package twopointer

import (
	"math"
	"sort"
)

/*

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

*/
//[2,7,11,15]
func ThreeSumUnorderedListBruteForce(input []int, target int) []int {

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				sum := input[i] + input[j] + input[k]

				if sum == target {
					return []int{i, j, k}
				}

			}

		}

	}
	return []int{-1, -1, -1}
}

func ThreeSumUnorderedList(input []int, target int) []int {

	sort.Ints(input)
	for i := 0; i < len(input)-2; i++ {
		j := i + 1
		k := len(input) - 1
		for j < k {
			sum := input[i] + input[j] + input[k]
			if sum == target {
				return []int{input[i], input[j], input[k]}
			} else if sum > target {
				k--
			} else {
				j++
			}

		}

	}
	return []int{-1, -1, -1}

}

// [2,7,11,15]
func ThreeSumOrderedList(input []int, target int) []int { //save target - num -----> in map and then search

	for i := 0; i < len(input)-2; i++ {
		j := i + 1
		k := len(input) - 1
		for j < k {
			sum := input[i] + input[j] + input[k]
			if sum == target {
				return []int{i, j, k}
			} else if sum > target {
				k--
			} else {
				j++
			}

		}

	}
	return []int{-1, -1, -1}

}

/*Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.



Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
Example 2:

Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).

*/

func threeSumClosest(input []int, target int) int {
	var result int = math.MaxInt
	var minAbsDiff int = math.MaxInt
	for i := 0; i < len(input)-2; i++ {
		j := i + 1
		k := len(input) - 1

		for j < k {
			sum := input[i] + input[j] + input[k]
			absDiff := abs(sum - target)
			if sum == target {
				return sum
			} else if sum > target {
				k--
			} else {
				j++
			}

			if absDiff < minAbsDiff {

				minAbsDiff = absDiff
				result = sum

			}

		}

	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
