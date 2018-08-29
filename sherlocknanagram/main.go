package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) (ans int32) {
	ss := map[int][]string{}
	// Make all possible substrings
	for i := 1; i < len(s); i++ {
		for j := 0; j+i-1 < len(s); j++ {
			ss[i] = append(ss[i], s[j:j+i])
		}
	}

	// Compare substrings of equal length
	for _, l := range ss {
		for j, t := range l {
			for i := j + 1; i < len(l); i++ {
				if isAnagram(t, l[i]) {
					ans++
				}
			}
		}
	}
	return ans
}

func isAnagram(a, b string) bool {
	af := [26]int{0}
	bf := [26]int{0}

	for i := 0; i < len(a); i++ {
		af[a[i]-'a']++
		bf[b[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if af[i] != bf[i] {
			return false
		}
	}

	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)

	}

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
