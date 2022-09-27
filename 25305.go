package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func f_25305() {
	var N, k int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &k)
	grades := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &grades[i])
	}
	sort.Ints(grades)
	fmt.Fprint(writer, grades[N-k])
}
