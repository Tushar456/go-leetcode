package stringproblem

import stringproblem "github.com/Tushar456/go-leetcode/stringproblem/reverse"

func FindPalindromeString(s string) bool {
	rev := stringproblem.ReverseString(s)

	return rev == s
}

func FindPalindromeTwoPointer(s string) bool {

	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {

		if r[i] != r[j] {
			return false
		}

	}
	return true
}
func FindPalindromeTwoPointerOpt(s string) bool {

	r := []rune(s)
	l := len(r)
	count := 0
	i := 0

	for i = 0; i < l/2; i++ {
		//abcba

		if r[i] == r[l-i-1] {
			count++
		}

	}

	return count == i

}

func FindPalindromeInt(num int) bool {

	rev := stringproblem.ReverseInt(num)

	return rev == num
}

func longestPalindrome(s string) string {

	var result string

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; i <= j; j-- {

			if s[i] != s[j] {
				continue
			}

			if isPalindrome(i, j, s) {
				palidrome := s[i : j+1]

				if len(result) < len(palidrome) {
					result = palidrome
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
