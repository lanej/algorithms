package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) (ans int64) {
	geo := make(map[int64]int64, len(arr))

	for _, e := range arr {
		geo[e]++
	}

	var pc int64
	var pn int64
	var pe []int64

	if c, ok := geo[1]; ok {
		delete(geo, 1)

		if c > 2 {
			ans = geoOnes(c)
		}

		for nk, nc := range geo {
			pc = c
			pn = nk

			pn = nk * nk
			pe = []int64{1, nk, pn}
			pc *= c * nc * geo[pn]

			ans += pc
		}
	}

	for k, c := range geo {
		pc = c
		pn = k
		pe = []int64{k}
		for p := int64(1); p < 3; p++ {
			pn = pn * k
			pe = append(pe, pn)
			pc *= geo[pn]
		}
		ans += pc
	}

	return
}

func geoOnes(i int64) int64 {
	switch i {
	case 3:
		return 1
	case 4:
		return 4
	default:
		return 2*geoOnes(i-1) + (i - 4)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
