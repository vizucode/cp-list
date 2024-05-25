/*

	f(x) => | Ax + B |

	1 1 1 4
	1 * 1 + 4
	5
*/

package dynamicabsolute

import "fmt"

func Absolute(a, b, x int) (resp int) {
	resp = a*x + b
	if resp < 1 {
		resp *= -1
	}

	return resp
}

func DynamicAbsolute() {
	var (
		a    int
		b    int
		x    int
		k    int
		temp int
	)

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&k)
	fmt.Scan(&x)

	temp = Absolute(a, b, x)
	for i := 0; i < k-1; i++ {
		temp = Absolute(a, b, temp)
	}

	fmt.Println(temp)
}
