package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	var result int64
	result = 0
	x := int(n) / len(s)
	r := int(n) % len(s)

	fmt.Println(r, x)

	if x > 0 {
		var count int64
		count = 0
		for _, v := range s {
			if v == 'a' {
				count++
			}
		}
		result = count * int64(x)
	}

	for i := 0; i < r; i++ {
		if s[i] == 'a' {
			result++
		}
	}

	fmt.Println(n, len(s))
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
