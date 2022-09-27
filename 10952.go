package main

import (
	"bufio"
	"fmt"
	"os"
)

func f_10952() {
	var A, B int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		fmt.Fscan(reader, &A, &B)
		if A == 0 && B == 0 {
			break
		}
		fmt.Fprintln(writer, A+B)
	}
}
