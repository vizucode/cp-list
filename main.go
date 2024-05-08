/*
var length = 4
var input = 3

loops i = 2 < input:
i*i <= input

	if input % i == 0
	    print("bukan")

print("YA")
*/
package main

import "fmt"

func main() {

	var length int
	fmt.Scan(&length)

	for i := 0; i < length; i++ {

		var input int
		fmt.Scan(&input)

		//var isPrime = true

		for i := 2; i < input; i++ {
			if i*i <= input {
				if input%i == 0 {
					//isPrime = false
					fmt.Println("BUKAN")
					break
				}
			} else {
				fmt.Println("YA")
				break
			}
		}

		/*if isPrime {
		    fmt.Println("YA")
		}*/

	}

}
