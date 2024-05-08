/*

   var N = 5

   var num // 0,2,3,4,5,6,7,8,9,0,1,2,3,4

   i = 1 => 0,1,2,3,4,5
       j = 1 to eq i => 0,[0,1],[0,1,2],[0,1,2,3],[0,1,2,3,4],[0,1,2,3,4,5]
       num++
       if num > 9 {
           num = 0
       }

       print(num)

*/

package main

import "fmt"

func main() {

	var N int
	fmt.Scan(&N)
	var num int

	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(num)
			num++
			if num > 9 {
				num = 0
			}
		}
		fmt.Println("")
	}

}
