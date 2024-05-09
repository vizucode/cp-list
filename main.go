/*

   var length = scan

   loops i = 0 < length:

       var input = scan // 8
       var divideCounter = 0

       //expect YA

       loops j = 1 < input:
           if j * j <= input
               if input % j == 0 {
               	divideCounter++
				if divideCounter > 1 {
					break
				}
               }
	if divideCounter > 1 {
		print("BUKAN")
	}else {
		print("YA")
	}

*/

package main

import "fmt"

func main() {
	var length int
	fmt.Scan(&length)

	for i := 0; i < length; i++ {
		var input int
		fmt.Scan(&input)

		var divideCounter = 0

		for j := 2; j*j <= input; j++ {
			if input%j == 0 {
				fmt.Println(j)
				divideCounter++
				if divideCounter > 1 {
					break
				}
			}
		}
		if divideCounter > 1 { // Jika ada lebih dari 2 pembagi, maka bukan bilangan agak prima
			fmt.Println("BUKAN")
		} else { // Jika hanya ada 2 pembagi, maka bilangan agak prima
			fmt.Println("YA")
		}
	}
}

// package main

// import "fmt"

// func main() {

// 	var length int
// 	fmt.Scan(&length)

// 	for i := 0; i < length; i++ {
// 		var input int
// 		fmt.Scan(&input)

// 		var divideCounter = 0

// 		for i := 2; i < input; i++ {
// 			if i*i <= input {
// 				divideCounter++
// 				if input%i == 0 && divideCounter > 2 {
// 					fmt.Println("BUKAN")
// 					break
// 				}
// 			} else {
// 				fmt.Println("YA")
// 				break
// 			}
// 		}

// 		if divideCounter <= 2 {
// 			fmt.Println("YA")
// 		}
// 	}
// }
