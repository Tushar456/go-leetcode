package recurrsion

func fibonaci(n int) int { //0112358

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonaci(n-1) + fibonaci(n-2)

}

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)

}

func sumOfDigit(n int) int { //1234%10
	if n <= 0 {
		return 0
	}
	d := n % 10
	n = n / 10

	return d + sumOfDigit(n)

}

func reverese(arr []int, index int) []int { //12345
	if index <= len(arr)/2 {
		return arr
	}

	arr[index], arr[len(arr)-index-1] = arr[len(arr)-index-1], arr[index]

	return reverese(arr, index+1)

}

func search(arr []int, target int) int {
	return binarysearch(arr, target, 0, len(arr)-1)
}

func binarysearch(arr []int, target int, start int, end int) int {

	if start > end {
		return -1
	}

	m := start + (end-start)/2

	if arr[m] == target {
		return m
	} else if target < arr[m] {
		return binarysearch(arr, target, start, m-1)
	} else {
		return binarysearch(arr, target, m+1, end)
	}

}

func printSubsequences(arr []int, index int, currSub []int, result *[][]int) {

	if index == len(arr) {

		*result = append(*result, append([]int{}, currSub...))

		return
	}
	currSub = append(currSub, arr[index])
	printSubsequences(arr, index+1, currSub, result)

	currSub = currSub[:len(currSub)-1]
	printSubsequences(arr, index+1, currSub, result)

}

func SubsequencesSumEqualK(arr []int, target int) [][]int {

	result := [][]int{}
	currSub := []int{}
	currentSum := 0
	subsequencesSumEqualK(arr, 0, currentSum, target, currSub, &result)
	return result

}

func subsequencesSumEqualK(arr []int, index int, currentSum int, target int, currSub []int, result *[][]int) {

	if currentSum == target {

		*result = append(*result, append([]int{}, currSub...))

		return
	}

	if index == len(arr) {
		return
	}

	currSub = append(currSub, arr[index])
	subsequencesSumEqualK(arr, index+1, currentSum+arr[index], target, currSub, result)

	currSub = currSub[:len(currSub)-1]
	subsequencesSumEqualK(arr, index+1, currentSum, target, currSub, result)

}

func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)

}

func merge(left, right []int) []int {

	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {

		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result

}
