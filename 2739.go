package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &N)

	for i := 1; i < 10; i++ {
		fmt.Fprintln(writer, N, "*", i, "=", N*i)
	}
}
