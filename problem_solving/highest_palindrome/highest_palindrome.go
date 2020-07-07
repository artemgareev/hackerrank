package highest_palindrome

import (
	"strconv"
)

//https://www.hackerrank.com/challenges/richie-rich/problem
//Complete the highestValuePalindrome function below.
func highestValuePalindrome(s string, n int32, k int32) string {
	source := make([]int, len(s))
	for i, currentValue := range s {
		number, _ := strconv.Atoi(string(currentValue))
		source[i] = number
	}

	leftPart := make([]int, len(source)/2)
	rightPart := []int{}

	for i := 0; i < len(source)/2; i++ {
		leftPart[i] = source[i]
	}
	for i := len(source) - (len(source) / 2); i < len(source); i++ {
		rightPart = append(rightPart, source[i])
	}
	rightPart = reverseInts(rightPart)

	//check if we can make palindrome for a given 'k'
	canMakePalindrome, equalNumbersCount := canMakePalindrome(leftPart, rightPart, int(k))
	if !canMakePalindrome {
		return "-1"
	}

	freeReplacementCount := int(k)
	var i = 0
	var skippedEqualNumbers int
	for k > 0 && i < len(leftPart) {
		if leftPart[i] == rightPart[i] {
			skippedEqualNumbers++
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
			freeReplacementCount--
			i++
			continue
		}

		leftToChange := len(leftPart) - i - 1
		canReplace2Numbers := freeReplacementCount-leftToChange+equalNumbersCount-skippedEqualNumbers-2 >= 0
		if canReplace2Numbers {
			rightPart[i] = 9
			leftPart[i] = 9
			freeReplacementCount -= 2

			i++
			continue
		}

		if leftPart[i] != rightPart[i] {
			if leftPart[i] > rightPart[i] {
				rightPart[i] = leftPart[i]
			} else {
				leftPart[i] = rightPart[i]
			}
			freeReplacementCount -= 1
		}

		i++
	}

	//change middle number if 's' is odd
	if freeReplacementCount > 0 && len(s)%2 != 0 {
		source[len(source)/2] = 9
	}

	rightPart = reverseInts(rightPart)
	return makeResult(source, leftPart, rightPart)
}

func canMakePalindrome(leftPart []int, rightPart []int, replaceCount int) (bool, int) {
	var minimalChangesNeed int
	var equalNumbersCount int

	for i := 0; i < len(leftPart); i++ {
		if leftPart[i] != rightPart[i] {
			minimalChangesNeed++
		} else {
			equalNumbersCount++
		}

	}
	return (replaceCount - minimalChangesNeed) >= 0, equalNumbersCount
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

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
