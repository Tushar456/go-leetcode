package slidingwindo

import "math"

//1, 3, 5, -3, 2, 3, 6, 7  target is 8
func findMaxSubarraySumEqualTarget(numbers []int, targetSum int) []int {

	windowSum := 0
	L, R := 0, 0
	maxSubarraySize := 0
	maxSubArrayStartingEndingIndex := [2]int{}

	for R < len(numbers) {

		windowSum += numbers[R]

		for windowSum > targetSum {

			windowSum = windowSum - numbers[L]
			L++

			if windowSum == targetSum {

				if maxSubarraySize < R-L+1 {
					maxSubarraySize = R - L + 1
					maxSubArrayStartingEndingIndex[0] = L
					maxSubArrayStartingEndingIndex[1] = R

				}

			}

		}

		R++

	}

	return numbers[maxSubArrayStartingEndingIndex[0] : maxSubArrayStartingEndingIndex[1]+1]

}

//1, 3, 5, -3, 2, 3, 6, 7  target is 8
func findMinSubarraySumEqualTarget(numbers []int, targetSum int) []int {

	windowSum := 0
	R, L := 0, 0
	minSubarraySize := math.MaxInt
	maxWindowStartingEndingIndex := [2]int{}

	for R < len(numbers) {

		windowSum += numbers[R]

		for windowSum > targetSum {

			windowSum = windowSum - numbers[L]
			L++

			if windowSum == targetSum {

				if R-L+1 < minSubarraySize {
					minSubarraySize = R - L + 1
					maxWindowStartingEndingIndex[0] = L
					maxWindowStartingEndingIndex[1] = R

				}

			}

		}

		R++

	}
	if maxWindowStartingEndingIndex[0] == 0 && maxWindowStartingEndingIndex[1] == 0 {
		return nil
	}

	return numbers[maxWindowStartingEndingIndex[0] : maxWindowStartingEndingIndex[1]+1]

}

/*
Given an array of positive integers nums and a positive integer target, return the minimal length of a
subarray

	whose sum is  equal to target. If there is no such subarray, return 0 instead.

Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1
Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/

func findMinSubarraySumEqualGreaterTarget(numbers []int, targetSum int) int {

	windowSum := 0
	L, R := 0, 0
	minSubArraySize := math.MaxInt

	for R < len(numbers) {

		windowSum += numbers[R]

		for windowSum >= targetSum {
			if R-L+1 < minSubArraySize {
				minSubArraySize = R - L + 1
			}

			windowSum = windowSum - numbers[L]
			L++

		}
		R++

	}

	if minSubArraySize > len(numbers) {

		return 0

	}

	return minSubArraySize

}
