package highest_palindrome

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPalindrome(t *testing.T) {
	assert.Equal(t, highestValuePalindrome("3943", 0, 1), "3993")
	assert.Equal(t, highestValuePalindrome("092282", 0, 4), "992299")
	assert.Equal(t, highestValuePalindrome("092282", 0, 3), "992299")
	assert.Equal(t, highestValuePalindrome("0011", 0, 1), "-1")
	assert.Equal(t, "99399", highestValuePalindrome("11331", 0, 4))
	assert.Equal(t, "99999", highestValuePalindrome("11922", 0, 4))
	assert.Equal(t, "9990999", highestValuePalindrome("1110229", 0, 5))
	assert.Equal(t, highestValuePalindrome("5", 0, 1), "9")
	assert.Equal(t, "12921", highestValuePalindrome("12321", 0, 1))
	assert.Equal(t, "99111199", highestValuePalindrome("11111111", 0, 4))
	assert.Equal(t, "99111199", highestValuePalindrome("11111111", 0, 5))
	assert.Equal(t, "9919199", highestValuePalindrome("1111111", 0, 5))
	assert.Equal(t, "11111111", highestValuePalindrome("11111111", 0, 1))
	assert.Equal(t, "1119111", highestValuePalindrome("1111111", 0, 1))

}

func TestPalindromeFromFixtures(t *testing.T) {
	assertFromFile(
		t,
		"test_cases/test_case_29",
		"test_cases/test_case_29_expected",
	)
}

func assertFromFile(t *testing.T, testCasePath string, testCaseExpectedPath string) {
	f, err := os.Open(testCasePath)
	reader := bufio.NewReaderSize(f, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	s := readLine(reader)

	result := highestValuePalindrome(s, n, k)

	expectedResult, err := ioutil.ReadFile(testCaseExpectedPath) // just pass the file name
	checkError(err)

	df1, df2 := difference(result, string(expectedResult))
	fmt.Printf("%d\n%d\n", len(df1), len(df2))

	assert.Equal(t, string(expectedResult), result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func difference(str1 string, str2 string) (string, string) {
	if len(str1) != len(str2) {
		return "", ""
	}
	var diff1 string
	var diff2 string

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			diff1 += string(str1[i])
			diff2 += string(str2[i])
		}
	}

	return diff1, diff2
}
