package main

import (
	"bufio"
	"fmt"
	"os"
)

func f_25304() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var X, N int

	fmt.Fscan(reader, &X)
	fmt.Fscan(reader, &N)

	var a, b int
	sum := 0
	for i := 0; i <= N; i++ {
		fmt.Fscanln(reader, &a, &b)
		sum += a * b
	}

	if sum == X {
		fmt.Fprint(writer, "Yes")
	} else {
		fmt.Fprint(writer, "No")
	}
}
