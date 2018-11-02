package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) (jumps int32) {
	graph := map[int][]int{}
	max := len(c) - 1

	for i := 0; i < len(c); i++ {
		if c[i] == 1 {
			continue
		}

		edges := []int{}

		if i+1 <= max && c[i+1] == 0 {
			edges = append(edges, i+1)
		}

		if i+2 <= max && c[i+2] == 0 {
			edges = append(edges, i+2)
		}

		graph[i] = edges
	}

	// fmt.Printf("graph: %v\n", graph)

	edges := graph[0]

	for {
		nextEdges := []int{}

		jumps++

		for _, r := range edges {
			if r == max {
				return
			}

			nextEdges = append(nextEdges, graph[r]...)
		}

		// fmt.Printf("edges: %v, jumps: %v\n", edges, jumps)

		edges = nextEdges
	}
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

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := jumpingOnClouds(c)

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
