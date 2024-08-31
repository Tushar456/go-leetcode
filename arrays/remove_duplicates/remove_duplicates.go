package arrays

func RemoveDuplicateSlice(input []int) []int {
	dup := make(map[int]bool)
	var result []int
	for _, num := range input {

		if _, ok := dup[num]; !ok {

			dup[num] = true
			result = append(result, num)

		}

	}
	return result
}

func removeDuplicates(nums []int) int {
	dup := make(map[int]int)
	j := 0
	for i := 0; i < len(nums)-1; i++ {

		val, ok := dup[nums[i]]

		if !ok || val < 2 {

			dup[nums[i]] = dup[nums[i]] + 1
			nums[j] = nums[i]
			j++

		} else {
			dup[nums[i]] = dup[nums[i]] + 1
		}

	}

	return len(nums)

}

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/
func ContainsDuplicateSlice(input []int) bool {
	dup := make(map[int]struct{})
	for _, num := range input {

		if _, ok := dup[num]; ok {

			return true
		}

		dup[num] = struct{}{}

	}

	return false
}
