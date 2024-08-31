package twopointer

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
*/
func findMaxArea(numbers []int) int {

	n := len(numbers)
	L := 0
	R := n - 1

	maxArea := 0

	for R > L {
		currentArea := 0
		if numbers[L] < numbers[R] {
			currentArea = (R - L) * numbers[L]
			L++
		} else {
			currentArea = (R - L) * numbers[R]
			R--
		}

		if currentArea > maxArea {
			maxArea = currentArea
		}

	}

	return maxArea

}
