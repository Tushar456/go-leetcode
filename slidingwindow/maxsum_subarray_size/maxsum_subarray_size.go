package slidingwindow

/*

You are given an array of integers nums find the maxsum of subarray of size k

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: 16

*/

func findMaxSumOfSubArrayOfSizeKBruteForce(numbers []int, k int) int {
	maxSum := 0
	for i := 0; i < len(numbers)-k+1; i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += numbers[j]
		}

		maxSum = max(maxSum, sum)

	}

	return maxSum

}

//[1,3,-1,-3,5,3,6,7
func findMaxSumOfSubArrayOfSizeK(numbers []int, k int) int {

	maxSum := 0
	windowSum := 0

	for i := 0; i < k; i++ {

		windowSum += numbers[i]

	}

	L := 0
	R := k

	for R < len(numbers) {

		windowSum = windowSum + numbers[R] - numbers[L]
		if windowSum > maxSum {
			maxSum = windowSum
		}
		L++
		R++
	}
	return maxSum
}

/*

You are given an array of integers nums find the maxsum of subarray of size k

Input: nums = [1,3,-1,-35,3,6,7], k = 3
Output: 16
[]int{1, 3, -1, -3, 5, 3, 6, 7}, k: 2
*/

func findAverageSumofSubArrayOfSizeKKBruteForce(numbers []int, k int) []float64 {
	var result []float64
	for i := 0; i < len(numbers)-k+1; i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += numbers[j]
		}

		result = append(result, float64(sum)/float64(k))

	}

	return result

}
func findAverageSumofSubArrayOfSizeK(numbers []int, k int) []float64 {

	windowSum := 0
	var result []float64

	for i := 0; i < k; i++ {

		windowSum += numbers[i]

	}

	result = append(result, float64(windowSum/k))

	L := 0
	R := k

	for R < len(numbers) {

		windowSum = windowSum + numbers[R] - numbers[L]
		result = append(result, float64(windowSum)/float64(k))
		L++
		R++
	}
	return result
}
