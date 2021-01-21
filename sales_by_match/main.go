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

type sock struct {
	color int32
	count int32
}

type pile []int32

func (p pile) Len() int {
	return len(p)
}

func (p pile) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p pile) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	if n == 1 {
		return 0
	}

	socks := make([]sock, n)
	p := make(pile, n)
	for i, v := range ar {
		p[i] = v
	}

	sort.Sort(p)

	aux := 0
	socks[aux].color = p[0]
	socks[aux].count++

	for i := 1; i < int(n); i++ {
		if p[i-1] == p[i] {
			socks[aux].count++
		} else {
			aux++
			socks[aux].color = p[i]
			socks[aux].count++
		}
	}

	var total int32
	total = 0

	for i := 0; i < int(n) && socks[i].count != 0; i++ {
		total += socks[i].count / 2
	}

	return total
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

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

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
