package main

import (
	"bufio"
	"fmt"
	"os"
)

func f_2798() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, M int

	fmt.Fscanln(reader, &N, &M)
	numbers := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &numbers[i])
	}

	sum := 300000
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				threesum := numbers[i] + numbers[j] + numbers[k]
				if threesum <= M {
					if (M-sum)*(M-sum) > (M-threesum)*(M-threesum) {
						sum = threesum
					}
				}
			}
		}
	}
	fmt.Fprint(writer, sum)
}
