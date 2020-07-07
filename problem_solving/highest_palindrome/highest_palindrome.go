package highest_palindrome

import (
	"strconv"
)

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
	canMakePalindrome, needChange1, needChange2 := canMakePalindrome(leftPart, rightPart, int(k))
	if !canMakePalindrome {
		return "-1"
	}

	_, _ = needChange1, needChange2

	freeReplacementCount := int(k)
	var i = 0
	for k > 0 && i < len(leftPart) {
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
			needChange1--
			i++
			continue
		}

		leftToChange := len(leftPart) - i - 1
		canReplace2Numbers := freeReplacementCount-leftToChange >= 0
		if freeReplacementCount >= 2 && needChange2 > 0 && canReplace2Numbers {
			rightPart[i] = 9
			leftPart[i] = 9
			freeReplacementCount -= 2
			needChange2--

			i++
			continue
		} else {
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
	}

	//if we have free replacements - make middle number 9
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

func canMakePalindrome(leftPart []int, rightPart []int, replaceCount int) (bool, int, int) {
	var minimalChangesNeed int
	var needChange1 int
	var needChange2 int

	for i := 0; i < len(leftPart); i++ {
		if leftPart[i] == 9 || rightPart[i] == 9 {
			needChange1++
		} else if leftPart[i] != rightPart[i] || leftPart[i] != 9 {
			needChange2++
		}
		if leftPart[i] != rightPart[i] {
			minimalChangesNeed++
		}

	}
	return (replaceCount - minimalChangesNeed) >= 0, needChange1, needChange2
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
