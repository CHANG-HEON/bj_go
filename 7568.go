package main

import (
	"bufio"
	"fmt"
	"os"
)

func f_7568() {
	var N, count int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &N)
	weight := make([]int, N)
	height := make([]int, N)
	num := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &weight[i], &height[i])
	}
	for j := 0; j < N; j++ {
		count = 1
		for k := 0; k < N; k++ {
			if (weight[j] < weight[k]) && (height[j] < height[k]) {
				count++
			}
			num[j] = count
		}
	}
	for _, grade := range num {
		fmt.Fprint(writer, grade, " ")
	}
}
