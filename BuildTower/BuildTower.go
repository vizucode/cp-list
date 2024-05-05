package buildtower

import "fmt"

func BuildTower(N int) {

	blockList := []string{}

	for i := 1; i <= N; i++ {
		for j := N; j >= i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}

	fmt.Println(blockList)
}
