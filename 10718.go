package main

import (
	"bufio"
	"fmt"
	"os"
)

func f_10718() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b int
	fmt.Fscan(reader, &a, &b)

	fmt.Fprintln(writer, a+b)
	fmt.Fprintln(writer, a-b)
	fmt.Fprintln(writer, a*b)
	fmt.Fprintln(writer, a/b)
	fmt.Fprintln(writer, a%b)
}
