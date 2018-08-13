package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) (ans int32) {
	swap := func(i, j int32) {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
		fmt.Printf("swap: pos(%d)=%d | pos(%d)=%d\narr: %v\n", i, j, arr[j], arr[i], arr)
	}

	size := int32(len(arr))
	for i := int32(0); i < size; i++ {
		fmt.Printf("i: %d; arr: %v\n", i, arr)
		if arr[i] == i+1 || arr[i] > size {
			continue
		}
		swap(i, arr[i]-1)
		ans++
		i--
	}

	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

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
