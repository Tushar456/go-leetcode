package twopointer

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
func TwoSumUnorderedListBruteForce(input []int, target int) []int {

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			sum := input[i] + input[j]

			if sum == target {
				return []int{i, j}
			}

		}

	}
	return []int{-1, -1}
}

func TwoSumUnorderedList(input []int, target int) []int { //save target - num -----> in map and then search

	mapDiff := make(map[int]int)

	for index, num := range input {

		if i, ok := mapDiff[num]; ok {

			return []int{i, index}

		} else {
			mapDiff[target-num] = index
		}

	}

	return []int{-1, -1}

}

//[2,7,11,15]
func TwoSumOrderedList(input []int, target int) []int { //save target - num -----> in map and then search

	i := 0
	j := len(input) - 1
	for i < j {
		sum := input[i] + input[j]
		if sum == target {
			return []int{i, j}
		} else if sum > target {
			j--
		} else {
			i++
		}

	}
	return []int{-1, -1}
}

func TwoSumUnorderedListMultiple(input []int, target int) [][]int { //save target - num -----> in map and then search

	mapDiff := make(map[int]int)

	var result [][]int

	for index, num := range input {

		if i, ok := mapDiff[num]; ok {

			result = append(result, []int{i, index})

		} else {
			mapDiff[target-num] = index
		}

	}
	if len(result) > 0 {
		return result
	}
	return [][]int{{-1, -1}}

}

func TwoSumOrderedListMultiple(input []int, target int) [][]int {
	var result [][]int
	i := 0
	j := len(input) - 1
	for i < j {
		sum := input[i] + input[j]
		if sum == target {
			result = append(result, []int{i, j})
			i++
			j--
		} else if sum > target {
			j--
		} else {
			i++
		}

	}
	if len(result) > 0 {
		return result
	}
	return [][]int{{-1, -1}}
}

func TwoSumWqualThirdElement(input []int) []int { //save target - num -----> in map and then search

	for k := len(input) - 1; k >= 0; k-- {

		i := 0
		j := k - 1
		for i < j {
			sum := input[i] + input[j]
			if sum == input[k] {
				return []int{i, j, k}
			} else if sum > input[k] {
				j--
			} else {
				i++
			}

		}

	}
	return []int{-1, -1, -1}
}
