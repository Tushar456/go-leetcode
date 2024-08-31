package stringproblem

import "strings"

func ReverseString(str string) string {
	r := []rune(str)
	rn := make([]rune, len(r))
	j := 0
	for i := len(r) - 1; i >= 0; i-- {
		rn[j] = r[i]
		j++
	}
	return string(rn)
}

func ReverseStringAppened(str string) (result string) {
	// iterate over str and prepend to result
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func ReverseStringTwoPointer(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReverseInt(num int) int { //123  321

	rev := 0

	for num > 0 {
		r := num % 10
		rev = rev*10 + r //321
		num = num / 10

	}

	return rev

}

func ReverseWordofSentence(senetence string) string {

	words := strings.Fields(senetence)

	for i, word := range words {

		words[i] = ReverseString(word)

	}

	return strings.Join(words, " ")

}
