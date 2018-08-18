package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the arrayManipulation function below.
func arrayManipulation(n int, queries [][]int) (ans int) {
	// a, b, and k : the left index, right index and summand.
	m := make([]int, n)
	for _, q := range queries {
		a, b, k := q[0], q[1], q[2]
		m[a-1] += k
		if b < n {
			m[b] -= k
		}
	}

	var x int
	for _, v := range m {
		x += v
		if ans < x {
			ans = x
		}
	}

	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int(mTemp)

	var queries [][]int
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
