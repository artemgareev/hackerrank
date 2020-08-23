package pairs

//https://www.hackerrank.com/challenges/pairs/problem
// Complete the pairs function below.
func pairs(k int32, arr []int32) int32 {
	numbersDict := map[int32]int32{}
	var pairsCnt int32
	for i := 0; i < len(arr); i++ {
		numbersDict[arr[i]]++
	}

	for i := 0; i < len(arr); i++ {
		pair, ok := numbersDict[arr[i]-k]
		if ok {
			pairsCnt += pair
		}
	}

	return pairsCnt
}
