package common_child

import "math"

//https://www.hackerrank.com/challenges/common-child/problem
//https://en.wikipedia.org/wiki/Longest_common_subsequence_problem

//example for given "ABCDEF", "FBDAMN":
//	0	A	B	C	D	E	F
//0	0	0	0	0	0	0	0
//F 0	0	0	0	1	1	1
//B 0	0	1	1	1	1	1
//D	0	0	1	1	1	1	1
//A	0	0	1	2	2	2	2
//M	0	0	1	2	2	2	2
//N	0	1	1	2	2	2	2
//
// Complete the commonChild function below.
func commonChild(s1 string, s2 string) int32 {
	c := make([][]int, len(s1)+1)
	for i := 0; i < len(c); i++ {
		c[i] = make([]int, len(s2)+1)
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else {
				c[i][j] = int(math.Max(float64(c[i][j-1]), float64(c[i-1][j])))
			}
		}
	}

	return int32(c[len(s1)][len(s2)])
}
