package main

import (
	"bufio"
	"fmt"
	"os"
)

func f_10926() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var ID string

	fmt.Fscan(reader, &ID)
	fmt.Fprint(writer, ID, "??!")

}
