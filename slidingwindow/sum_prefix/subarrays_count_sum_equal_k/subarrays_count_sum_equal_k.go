package slidingwindow

/*
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.



Example 1:

Input: nums = [1,1,1], k = 2
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
Output: 2



Input: nums = [1, 4, -4, 5, -5, 0], k=0
Output: 6
*/
func findSubArrayWhoseSumEqualtoTarget(numbers []int, targert int) int {

	sum := 0
	mapSum := make(map[int]int)
	count := 0
	mapSum[0] = 1

	for _, num := range numbers {
		sum += num

		complement := sum - targert

		count += mapSum[complement]

		mapSum[sum] += 1
	}
	return count
}
