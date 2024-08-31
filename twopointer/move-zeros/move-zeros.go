package twopointer

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]
*/

//0,1,0,3,12
//1,0,0,3,12
func moveZeros(input []int) []int {

	n := len(input)

	if n < 2 {
		return input
	}
	L := 0 // if L is zero replace with right element
	R := 1 // find the next non zero

	for R < n {

		if input[L] != 0 {
			L++
			R++
		} else if input[R] == 0 {
			R++
		} else {
			input[L], input[R] = input[R], input[L]
		}

	}

	return input

}

func moveZeroSimple(input []int) []int {

	nonZeroIndex := 0

	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			input[i], input[nonZeroIndex] = input[nonZeroIndex], input[i]
			nonZeroIndex++
		}

	}

	return input

}
