package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var limit int32 = 2
var limitExceeded = "Too chaotic"

// Complete the minimumBribes function below.
func minimumBribes(q []int32) interface{} {
	var bribes int32

	for i := range q {
		d := q[i] - 1 - int32(i)
		if d > limit {
			return limitExceeded
		}

		for j := int(math.Max(float64(0), float64(q[i]-2))); j < i; j++ {
			if q[j] > q[i] {
				bribes++
			}
		}

	}

	return bribes
}

func printMinimumBribes(q []int32) {
	fmt.Println(minimumBribes(q))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		printMinimumBribes(q)
	}
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
