package main

import "fmt"

func main() {

	var arr = []int{}

	for {

		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		arr = append(arr, num)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}

}
