package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func activityNotifications(expenditure []int32, d int32) (ans int32) {
	rmediane := make([]int32, d)
	copy(rmediane[:], expenditure[:d])
	rmedian := median(rmediane)

	for i := int(d); i < len(expenditure); i++ {
		if float32(expenditure[i]) >= 2*rmedian {
			ans++
		}

		rmediane = append(rmediane[1:], expenditure[i])
		rmedian = median(rmediane)
	}

	return
}

func median(a []int32) float32 {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	if len(a)%2 == 0 {
		return float32(a[len(a)/2]) + float32(a[len(a)/2+1])/2
	}

	return float32(a[len(a)/2])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
