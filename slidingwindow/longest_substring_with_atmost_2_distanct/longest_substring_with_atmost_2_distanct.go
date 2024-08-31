package slidingwindow

/*
Given a string s , find the length of the longest substring t  that contains at most 2 distinct characters.

Example 1:

Input: "eceba"
Output: 3
Explanation: t is "ece" which its length is 3.
Example 2:

Input: "ccaabbb"
Output: 5
Explanation: t is "aabbb" which its length is 5.
*/
func longestSubstringWithAtmost2DistanctCharacter(input string) int {

	characters := []rune(input)

	L, R := 0, 0

	maxLength := 0

	mapCh := make(map[rune]int)

	for R < len(input) {

		mapCh[characters[R]]++

		for len(mapCh) > 2 {

			mapCh[characters[L]]--

			if mapCh[characters[L]] == 0 {
				delete(mapCh, characters[L])
			}
			L++
		}

		maxLength = max(maxLength, R-L+1)
		R++

	}
	return maxLength

}

func longestSubstringWithAtmostKDistanctCharacter(input string, k int) int {

	if k == 0 {
		return 0
	}
	characters := []rune(input)

	L, R := 0, 0

	maxLength := 0

	mapCh := make(map[rune]int)

	for R < len(input) {

		mapCh[characters[R]]++

		for len(mapCh) > k {

			mapCh[characters[L]]--

			if mapCh[characters[L]] == 0 {
				delete(mapCh, characters[L])
			}
			L++
		}

		maxLength = max(maxLength, R-L+1)
		R++

	}
	return maxLength

}
