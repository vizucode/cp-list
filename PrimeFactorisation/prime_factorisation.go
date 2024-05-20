package primefactorisation

import (
	"fmt"
	"strings"
)

/*
input
N = 12

12 => 12 / 2 = 6
6  => 6 / 2 = 3
3  => 3 / 2 can't be dived

the factorisation is = 2 * 2 * 3 = 2 ^ 2 * 3
*/
func PrimeFactoristation() { // 12
	var n int
	fmt.Scan(&n)

	divider := 2
	arrPower := make([]int, n+1)
	// jml := n
	for n > 1 { // 12
		// check if  modulo < 1 then add powerNumm
		if n%divider < 1 {
			arrPower[divider]++ // powerNum = 1
			n /= divider        // 12 % 2 = 6
		} else {
			// check if n is an prime number than it assume done.
			divider++
			if divider*divider > n {
				arrPower[n]++
				break
			}

			// if checkPrime(n) && n%divider < 2 {
			// 	arrPower[n] = 1
			// 	break
			// }

			// nearPrime := nearestPrime(n)
			// if nearPrime > 0 {
			// 	divider = nearPrime
			// 	arrPower[divider] = 1
			// }
		}
	}

	text := ""
	for i := 0; i < len(arrPower); i++ {
		if arrPower[i] > 1 {
			text += fmt.Sprintf("%d^%d x ", i, arrPower[i])
		} else if arrPower[i] == 1 {
			text += fmt.Sprintf("%d x ", i)
		}
	}
	text = strings.TrimSuffix(text, "x ")

	fmt.Print(text)
}
