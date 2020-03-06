package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/strange-advertising/problem

// Complete the viralAdvertising function below.
func viralAdvertising(n int32) int32 {
	if n == 1 {
		return 2
	}

	return viralAdvertising(n-1) + liked(n)
}

func shared(n int32) int32 {
	if n == 1 {
		return 5
	}

	return liked(n-1) * 3
}

func liked(n int32) int32 {
	if n == 1 {
		return 5 / 2
	}

	return shared(n) / 2
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

	result := viralAdvertising(n)

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
