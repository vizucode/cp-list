package main

import (
	outlier "competitive_programming_practice/FindTheParityOutlier"
	"fmt"
)

func main() {
	ot := outlier.FindOutlier([]int{2, 4, 8, -3, 40, 20})
	fmt.Println(ot)
}
