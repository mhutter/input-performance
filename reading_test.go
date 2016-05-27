package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

var n int

func BenchmarkScanf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Scanf("%d", &n)
	}
}

func BenchmarkScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Scan(&n)
	}
}

func BenchmarkBufferedRead(b *testing.B) {
	reader := bufio.NewReader(os.Stdin)
	var line string
	for i := 0; i < b.N; i++ {
		line, _ = reader.ReadString('\n')
		n, _ = strconv.Atoi(line)
	}
}

func BenchmarkBufferedScan(b *testing.B) {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < b.N; i++ {
		n, _ = strconv.Atoi(scanner.Text())
	}
}
