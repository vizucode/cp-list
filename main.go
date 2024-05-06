package main

import (
	"fmt"
	"log"
)

func main() {
	var base float32
	var height float32

	_, err := fmt.Scan(&base, &height)
	if err != nil {
		log.Fatal(err)
	}

	// 1/5 * base  * height
	var area = 0.5 * base * height
	fmt.Printf("%.2f", area) // 7.5
}
