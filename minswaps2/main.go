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
	pos := map[int32]int32{}

	for i, j := range arr {
		pos[j] = int32(i)
	}

	// Build binary tree and hashmap of positions
	// Starting at i = 0, grab next highest value and swap into first position
	// Remove first position mapping and binary tree entry
	// Increment i, and continue until last position
	for i := int32(0); i < int32(len(arr)); i++ {
		fmt.Printf("pos: %+v\n", pos)
		fmt.Printf("arr: %+v\n", arr)

		v := i + 1  // 1 // 2
		j := arr[i] // 1 // 3

		if v != j {
			// 4
			// swap pos
			si := pos[v]

			pos[j] = si
			pos[v] = i

			// swap arr
			arr[si] = j
			arr[i] = v
			ans++
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
