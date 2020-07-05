package highest_palindrome

import "strconv"

//Complete the highestValuePalindrome function below.
func highestValuePalindrome(s string, n int32, k int32) string {
	source := []int{}
	for _, currentValue := range s {
		i, _ := strconv.Atoi(string(currentValue))
		source = append(source, i)
	}

	leftPart := []int{}
	rightPart := []int{}

	for i := 0; i < len(source)/2; i++ {
		leftPart = append(leftPart, source[i])
	}
	for i := len(source) - (len(source) / 2); i < len(source); i++ {
		rightPart = append(rightPart, source[i])
	}

	rightPart = reverseInts(rightPart)

	//check if we can make palindrome for a given 'k'
	canMakePalindrome, freeReplacementCount := canMakePalindrome(leftPart, rightPart, int(k))
	if !canMakePalindrome {
		return "-1"
	}

	//just return palindrome with the highest replacements possible
	if freeReplacementCount == 0 {
		tL, tR, tC := tryMakePalindrome(leftPart, rightPart, freeReplacementCount)
		if tC == 0 {
			tR = reverseInts(tR)
			return makeResult(source, tL, tR)
		}
	}

	allReplacementCount := k
	var i = 0
	for freeReplacementCount > 0 && i < len(leftPart) {
		if len(leftPart)-1 == i && allReplacementCount%2 == 0 {
			rightPart[i] = 9
			leftPart[i] = 9
			freeReplacementCount -= 2
			allReplacementCount -= 2

			i++
			continue
		}
		if leftPart[i] == 9 && rightPart[i] == 9 {
			i++
			continue
		}
		if leftPart[i] == 9 || rightPart[i] == 9 {
			if leftPart[i] == 9 {
				rightPart[i] = 9
			} else {
				leftPart[i] = 9
			}
			freeReplacementCount -= 1
			allReplacementCount -= 1
			i++
			continue
		}
		if freeReplacementCount >= 2 {
			rightPart[i] = 9
			leftPart[i] = 9
			freeReplacementCount -= 2
			allReplacementCount -= 2

			i++
			continue
		}
		i++
	}

	leftPart, rightPart, _ = tryMakePalindrome(leftPart, rightPart, freeReplacementCount)
	//_ = abc

	//making the highest palindrome
	if freeReplacementCount > 0 && len(s)%2 != 0 {
		source[len(source)/2] = 9
	}

	rightPart = reverseInts(rightPart)
	return makeResult(source, leftPart, rightPart)

}
func makeResult(source []int, leftPart []int, rightPart []int) string {
	var resultString string
	var result []int
	if len(source)%2 != 0 {
		result = append(result, leftPart...)
		result = append(result, source[len(source)/2])
		result = append(result, rightPart...)
	} else {
		result = append(result, leftPart...)
		result = append(result, rightPart...)
	}

	for i := 0; i < len(result); i++ {
		resultString += strconv.Itoa(result[i])
	}

	return resultString
}

func tryMakePalindrome(leftPart []int, rightPart []int, replaceCount int) ([]int, []int, int) {
	leftReplacementCount := replaceCount
	for i := 0; i < len(leftPart); i++ {
		if leftPart[i] == rightPart[i] {
			continue
		}

		leftReplacementCount--
		if leftPart[i] > rightPart[i] {
			rightPart[i] = leftPart[i]
		} else {
			leftPart[i] = rightPart[i]
		}
	}

	return leftPart, rightPart, leftReplacementCount
}

func canMakePalindrome(leftPart []int, rightPart []int, replaceCount int) (bool, int) {
	var notPalindromeCount int
	for i := 0; i < len(leftPart); i++ {
		if leftPart[i] != rightPart[i] {
			if leftPart[i] != 9 && rightPart[i] != 9 {
				notPalindromeCount++
			}
		}
	}

	return (replaceCount - notPalindromeCount) >= 0, replaceCount - notPalindromeCount
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
