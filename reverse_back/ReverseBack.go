package reverseback

import "fmt"

func reverse(x int) int {
	var (
		temp = x
		ret  = 0
	)

	for temp > 0 {
		ret = (ret * 10) + (temp % 10)
		temp /= 10
	}

	return ret
}

func ReverseBack() {

	var (
		a int
		b int
	)

	fmt.Scan(&a)
	fmt.Scan(&b)

	c := reverse(a) + reverse(b)

	fmt.Println(reverse(c))

}
