/*

   var length = scan

   loops i = 0 < length:

       var input = scan // 8
       var divideCounter = 1

       //expect YA

       loops i = 2 < input:
           if i * i <= input
               if input % i == 0 && divideCounter > 2 {
                   print("BUKAN")
                   break
               }
               divideCounter++
           else {
               print("YA") // YA
               break
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

		var divideCounter = 0 // Mulai dari 2 karena setiap bilangan pasti habis dibagi oleh 1 dan bilangan itu sendiri

		for j := 2; j*j <= input; j++ {
			if input%j == 0 {
				divideCounter++
				if divideCounter > 2 { // Jika ada lebih dari dua pembagi, berhenti
					break
				}
			}
		}

		if divideCounter > 2 { // Jika ada lebih dari dua pembagi, maka bukan bilangan agak prima
			fmt.Println("BUKAN")
		} else { // Jika hanya ada dua pembagi, maka bilangan agak prima
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
