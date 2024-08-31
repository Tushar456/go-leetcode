package arrays

/*

The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.

If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.



Example 1:

Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The pivot index is 3.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11
Example 2:

Input: nums = [1,2,3]
Output: -1
Explanation:
There is no index that satisfies the conditions in the problem statement.
Example 3:

Input: nums = [2,1,-1]
Output: 0
Explanation:
The pivot index is 0.
Left sum = 0 (no elements to the left of index 0)
Right sum = nums[1] + nums[2] = 1 + -1 = 0

*/
//[1,7,3,6,5,6]
func pivot_index(numbers []int) int {

	for i := 0; i < len(numbers); i++ {

		LSum := 0
		RightSum := 0

		for j := i + 1; j < len(numbers); j++ {

			RightSum += numbers[j]

		}

		for k := i - 1; k >= 0; k-- {
			LSum += numbers[k]

		}

		if LSum == RightSum {
			return i
		}

	}
	return -1

}

//[1,7,3,6,5,6]
func pivotIndexOptimized(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	lSum := 0
	rSum := sum

	for i := 0; i < len(numbers); i++ {

		rSum -= numbers[i]
		if lSum == rSum {

			return i

		}
		lSum += numbers[i]

	}
	return -1
}
