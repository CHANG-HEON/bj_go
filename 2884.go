package main

import (
	"bufio"
	"fmt"
	"os"
)

func f_2884() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var H, M int

	fmt.Fscanln(reader, &H, &M)

	if M >= 45 {
		fmt.Fprint(writer, H, M-45)
	} else {
		if H > 0 {
			fmt.Fprint(writer, H-1, M+15)
		} else {
			fmt.Fprint(writer, 23, M+15)
		}
	}

}
