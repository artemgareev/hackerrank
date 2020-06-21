package bomber_man

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestBomber(t *testing.T) {
	f, err := os.Open("test_case_default")
	reader := bufio.NewReaderSize(f, 1024*1024)

	rcn := strings.Split(readLine(reader), " ")

	rTemp, err := strconv.ParseInt(rcn[0], 10, 64)
	checkError(err)
	r := int32(rTemp)

	nTemp, err := strconv.ParseInt(rcn[2], 10, 64)
	checkError(err)
	n := int(nTemp)

	var grid []string

	for i := 0; i < int(r); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	BomberMan(n, grid)

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
