/*

   var N = 5

   loops: i; 1 to N (++) = 1,2,3,4,5
       loops: space; N to i (--) = 5, 4, 3, 2, 1
       loops: star; 1 to i (++) = 1, 2, 3, 4, 5

   ####*
   ###**
   ##***
   #****
   *****

*/

package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		for space := N; space > i; space-- {
			fmt.Print(" ")
		}

		for star := 1; star <= i; star++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
