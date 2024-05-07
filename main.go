package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Scanln(&num)

	truncated := math.Trunc(num)
	F := int(truncated)
	C := int(truncated)

	if truncated != num {
		if num < 0 {
			F -= 1
		} else {
			C += 1
		}
	}

	fmt.Printf("%d %d", F, C)
}
