package slidingwindow

func lengthOfLongestSubstringBruteForce(s string) int {

	charcters := []rune(s)
	maxLength := 0

	for i := 0; i < len(charcters); i++ {

		mapCh := make(map[rune]int)

		for j := i; j < len(charcters); j++ {

			mapCh[charcters[j]]++

			if len(mapCh) < j-i+1 {
				break
			}

			maxLength = max(maxLength, j-i+1)

		}

	}
	return maxLength
}

// abcdaababa
func lengthOfLongestSubstring(s string) (int, string) {

	R, L := 0, 0
	longestSubstringLength := 0
	longestSubstringLengthPos := [2]int{}
	characters := []rune(s)
	characcterMap := make(map[rune]int)

	for R < len(s) {

		characcterMap[characters[R]]++

		for len(characcterMap) != R-L+1 {

			characcterMap[characters[L]]--

			if characcterMap[characters[L]] == 0 {
				delete(characcterMap, characters[L])
			}

			L++

		}
		if R-L+1 >= longestSubstringLength {
			longestSubstringLength = R - L + 1
			longestSubstringLengthPos[0] = L
			longestSubstringLengthPos[1] = R

		}
		R++

	}

	subChar := characters[longestSubstringLengthPos[0] : longestSubstringLengthPos[1]+1]

	return longestSubstringLength, string(subChar)
}

//aabccbb

func lengthOfLongestSubstringHavingSameLetter(s string, k int) int {
	characters := []rune(s)
	R, L := 0, 0
	maxRepeatLetterCount := 0
	maxLength := 0
	mapLetterFrequency := make(map[rune]int)

	for R < len(s) {

		mapLetterFrequency[characters[R]]++
		maxRepeatLetterCount = max(maxRepeatLetterCount, mapLetterFrequency[characters[R]])

		for R-L+1-maxRepeatLetterCount > k {

			mapLetterFrequency[characters[L]]--

			if mapLetterFrequency[characters[L]] == 0 {
				delete(mapLetterFrequency, characters[L])
			}

			L++

		}
		maxLength = max(maxLength, R-L+1)
		R++
	}

	return maxLength

}

// ababaedd
func longestPalindrome(s string) string {

	var result string

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {

			if s[i] != s[j] {
				continue
			}

			if isPalindrome(i, j, s) {
				palidromeString := s[i : j+1]

				if len(palidromeString) > len(result) && len(palidromeString) > 2 {
					result = palidromeString
				}
				break
			}

		}
	}
	return result
}

func isPalindrome(start, end int, s string) bool {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
