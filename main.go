package main

import (
	"fmt"
)

func main() {

	var N int
	fmt.Scan(&N)

	var arr = make([]int, 1001)

	for i := 0; i < N; i++ {
		var x int
		fmt.Scan(&x)
		arr[x]++
	}

	var maxFrequency = 0
	var modus = 0

	for i := 0; i < 1001; i++ {
		if arr[i] >= maxFrequency {
			modus = i
			maxFrequency = arr[i]
		}
	}

	fmt.Println(modus)
}
