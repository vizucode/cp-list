package oddeven

import "fmt"

func OddEven(n int) {
	// even number
	for i := 1; i <= n; i++ {
		if n%2 == 0 {
			if i%2 == 0 {
				fmt.Print(i)
			}
		}

		if n%2 == 1 {
			if i%2 == 1 {
				fmt.Print(i)
			}
		}
	}
}
