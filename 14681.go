package main

import (
	"bufio"
	"fmt"
	"os"
)

func f_14681() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var x, y int

	fmt.Fscanln(reader, &x)
	fmt.Fscanln(reader, &y)

	if x > 0 && y > 0 {
		fmt.Fprint(writer, 1)
	} else if x > 0 && y < 0 {
		fmt.Fprint(writer, 4)
	} else if x < 0 && y > 0 {
		fmt.Fprint(writer, 2)
	} else if x < 0 && y < 0 {
		fmt.Fprint(writer, 3)
	}

}
