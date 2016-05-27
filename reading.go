package main

import (
	"fmt"
	"io"
)

func main() {
	var out []int
	var t int
	for {
		_, err := fmt.Scan(&t)
		if err == io.EOF {
			break
		}
		out = append(out, t)
	}

	fmt.Printf("%v\n", out)
}
