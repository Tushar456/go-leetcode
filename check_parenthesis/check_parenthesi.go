package checkparenthesis

import "math"

func checkBalancedParenthesisANdReturnHighest(str string) int { //(((X)))(Y))
	var currentHeight int
	var maxHeight int
	for _, ch := range str {

		if ch == '(' {

			currentHeight++
		} else if ch == ')' {
			if currentHeight > 0 {
				currentHeight--
			} else {
				return -1
			}

		}

		if currentHeight > maxHeight {
			maxHeight = currentHeight
		}

	}

	if currentHeight == 0 {
		return maxHeight
	}
	return -1

}

func longestBalancedParenthesisInUnbalanced(str string) int { //((()())) - > 1  //((()())((())))
	var max int
	var stack []int
	stack = append(stack, -1)
	for i, ch := range str {

		if ch == '(' {

			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				max = int(math.Max(float64(max), float64(i-stack[len(stack)-1])))
			}
		}
	}
	return max

}
