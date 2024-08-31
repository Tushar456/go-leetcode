package arrays

func topKFrequent(nums []int, k int) []int {
	n := len(nums)

	// Create count map and count the frequencies
	countMap := make(map[int]int)
	for i := 0; i < n; i++ {
		countMap[nums[i]]++
	}

	// Create bucket array, with n+1 lenght
	// Note: making interface to eliminate the default 0 value of slice
	bucket := make([][]int, n+1)

	// Fill the bucket array
	for e, c := range countMap {
		bucket[c] = append(bucket[c], e)
	}

	ans := make([]int, 0)
	// Loop from back and fill the ans array
	for i := n; i >= 0; i-- {
		e := bucket[i]
		if e == nil {
			continue
		}

		ans = append(ans, e...)

		// If we have k elements
		if len(ans) >= k {
			ans = ans[:k]
			break
		}
	}

	return ans
}
