package main

import (
	"fmt"
	"log"
)

func main() {

	var (
		num1 uint
		num2 uint
		num3 uint
		num4 uint
		num5 uint
		num6 uint
		num7 uint
		num8 uint
		num9 uint
	)

	_, err := fmt.Scanln(&num1, &num2, &num3)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Scanln(&num4, &num5, &num6)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Scanln(&num7, &num8, &num9)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(num1, num4, num7)
	fmt.Println(num2, num5, num8)
	fmt.Println(num3, num6, num9)
}
