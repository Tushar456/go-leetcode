package slidingwindow

import "math"

/*
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

For example, S = "ADOBECODEBANC" T = "ABC" Minimum window is "BANC".

Note: If there is no such window in S that covers all characters in T, return the empty string "".

If there are multiple such windows, you are guaranteed that there will always be only one unique minimum window in S.
*/
func findMinimumWindoeSubstring(s string, t string) string {

	mapCh := make(map[rune]int)

	for _, ch := range t { //// store all the character of pattern in map
		mapCh[ch]++
	}

	L, R := 0, 0
	count := 0

	minSubstringLength := math.MaxInt

	minSubstring := ""

	characters := []rune(s)

	for R < len(s) { //"ADOBECODEBANC" T = "ABC"

		_, ok := mapCh[characters[R]]

		if ok {

			mapCh[characters[R]]-- // if charcter found in the map decrement the character
			if mapCh[characters[R]] >= 0 {
				count++
			}

		}

		for count == len(t) {

			if minSubstringLength > R-L+1 {
				minSubstringLength = R - L + 1
				minSubstring = string(characters[L : R+1])
			}

			_, ok := mapCh[characters[L]]

			if ok {
				mapCh[characters[L]]++
				if mapCh[characters[L]] > 0 {
					count--
				}

			}

			L++
		}
		R++

	}
	return minSubstring
}

/*
Given a string and a pattern, find out if the string contains any permutation of the pattern.

Example 1:

Input: String="oidbcaf", Pattern="abc"
Output: true
Explanation: The string contains "bca" which is a permutation of the given pattern.

Example 2:

Input: String="odicf", Pattern="dc"
Output: false
Explanation: No permutation of the pattern is present in the given string as a substring.

Example 3:

Input: String="bcdxabcdy", Pattern="bcdyabcdx"
Output: true
Explanation: Both the string and the pattern are a permutation of each other.
*/

func StringPermutationCheck(input string, pattern string) bool {

	characters := []rune(input)

	patterMap := make(map[rune]int)

	for _, ch := range pattern { // store all the character of pattern in map
		patterMap[ch]++

	}

	L, R := 0, 0

	matcher := 0

	for R < len(input) { ////"ADOBECODEBANC" T = "ABC"

		if _, ok := patterMap[characters[R]]; ok { // if charcter found in the map decrement the character
			patterMap[characters[R]]--
			if patterMap[characters[R]] == 0 { // if count of character in pattern map is equal to 0 found a match
				matcher++
			}
		}

		if matcher == len(patterMap) { // if the length of patternMap == match count we found the permutation of pattern
			return true
		}

		if R-L+1 >= len(pattern) { // if the window size is greater than size of pattern contract the window by sliding
			_, ok := patterMap[characters[L]]

			if ok {
				if patterMap[characters[L]] == 0 {
					matcher--
				}
				patterMap[characters[L]]++

			}

			L++
		}
		R++
	}

	return false

}

/*
Given a string and a pattern, find all anagrams of the pattern in the given string.

Example 1:

Input: String="ppqp", Pattern="pq"
Output: [1, 2]
Explanation: The two anagrams of the pattern in the given string are "pq" and "qp".

Example 2:

Input: String="abbcabc", Pattern="abc"
Output: [2, 3, 4]
Explanation: The three anagrams of the pattern in the given string are "bca", "cab", and "abc".

*/
func StringAnagrams(input string, pattern string) []int {

	characters := []rune(input)

	patterMap := make(map[rune]int)

	for _, ch := range pattern { // store all the character of pattern in map
		patterMap[ch]++

	}

	L, R := 0, 0

	matcher := 0

	var result []int

	for R < len(input) { ////"ADOBECODEBANC" T = "ABC"

		if _, ok := patterMap[characters[R]]; ok { // if charcter found in the map decrement the character
			patterMap[characters[R]]--
			if patterMap[characters[R]] == 0 { // if count of character in pattern map is equal to 0 found a match
				matcher++
			}
		}

		if matcher == len(patterMap) { // if the length of patternMap == match count we found the permutation of pattern
			result = append(result, L)
		}

		if R-L+1 >= len(pattern) { // if the window size is greater than size of pattern contract the window by sliding
			_, ok := patterMap[characters[L]]

			if ok {
				if patterMap[characters[L]] == 0 {
					matcher--
				}
				patterMap[characters[L]]++
			}

			L++
		}
		R++
	}

	if len(result) == 0 {
		return []int{}
	}

	return result

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
