package slidingwindow

/*
Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:

Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.

*/
func findMaxConsecutiveOneWithoutFliping(nums []int) int {

	maxConsecutiveOneSubArray := 0
	counter := 0

	for _, num := range nums {

		if num == 0 {
			counter = 0
		} else {
			counter += 1
		}

		maxConsecutiveOneSubArray = max(maxConsecutiveOneSubArray, counter)

	}

	return maxConsecutiveOneSubArray

}

/*
Given a binary array, find the maximum number of consecutive 1s in this array if you can flip at most one 0.

Example 1:

Input: [1,0,1,1,0]
Output: 4
Explanation: Flip the first zero will get the the maximum number of consecutive 1s.
    After flipping, the maximum number of consecutive 1s is 4.
*/
func findMaxConsecutiveOneWithOneFliping(nums []int) int {

	maxConsecutiveOneSubArray := 0
	counter := 0
	L, R := 0, 0

	for R < len(nums) {

		if nums[R] == 0 {
			counter++
		}

		for counter > 1 {

			if nums[L] == 0 {
				counter--
			}
			L++

		}

		maxConsecutiveOneSubArray = max(maxConsecutiveOneSubArray, R-L+1)

		R++

	}
	return maxConsecutiveOneSubArray

}

func findMaxConsecutiveOneWithKFliping(nums []int, k int) int {

	maxConsecutiveOneSubArray := 0
	counter := 0
	L, R := 0, 0

	for R < len(nums) {

		if nums[R] == 0 {
			counter++
		}

		for counter > k {

			if nums[L] == 0 {
				counter--
			}
			L++

		}

		maxConsecutiveOneSubArray = max(maxConsecutiveOneSubArray, R-L+1)

		R++

	}
	return maxConsecutiveOneSubArray

}

func max(a int, b int) int {

	if a > b {
		return a
	}

	return b

}
