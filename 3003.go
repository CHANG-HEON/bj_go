package main

import (
	"bufio"
	"fmt"
	"os"
)

func f_3003() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c, d, e, f int

	fmt.Fscanln(reader, &a, &b, &c, &d, &e, &f)
	fmt.Fprint(writer, 1-a, 1-b, 2-c, 2-d, 2-e, 8-f)
}
