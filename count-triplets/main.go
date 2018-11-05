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
	last := int64(len(arr) - 2)
	length := int64(len(arr))
	var n int64
	for i := int64(0); i < last; i++ {
		r = arr[i+1] / arr[i]
		gs := []int64{1}

		for n = i + 1; n < length; n++ {
			if arr[n] == arr[n-1] {
				gs[len(gs)-1]++
				continue
			}

			if arr[n] == arr[n-1]*r {
				gs = append(gs, 1)
				continue
			}
		}

		if len(gs) == 1 && gs[0] > 2 {
			ans += geoOnes(int64(gs[0]))
			i = i + gs[0]
            fmt.Printf("ans: %v, n: %v, r: %v, gs: %v\n", ans, n, r, gs)
			continue
		}

		if len(gs) > 2 {
			offset := int64(0)
			for j := 0; j < len(gs)-2; j++ {
				offset += gs[j]
				ans += int64(gs[j] * gs[j+1] * gs[j+2])
			}
			i = i + offset
		}

		fmt.Printf("n: %v, r: %v, gs: %v\n", n, r, gs)
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
