package main

import (
	"bufio"
	"fmt"
	"os"
)

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func f_10872() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &N)
	fmt.Fprint(writer, factorial(N))
}
