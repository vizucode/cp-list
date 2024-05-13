/*

	count frequency of numbers input
	find the max frequency

	count and select frequency with the numbers

	check if there was numbers bigger, so print the numbers
*/

package modus

import (
	"fmt"
	"math"
)

func main() {

	var N int
	fmt.Scan(&N)

	var numbers = make([]int, N)
	var frequency = make([]int, N+1000)
	var modus = make([]int, N+1000)

	for i := 0; i < N; i++ {
		fmt.Scan(&numbers[i])
		frequency[numbers[i]]++
	}

	// check max frequency
	var maxFrequency float64
	for i := 0; i < 1000; i++ {
		maxFrequency = math.Max(float64(maxFrequency), float64(frequency[i]))
	}

	// check most modus
	for i := 0; i < 1000; i++ {
		if frequency[i] == int(maxFrequency) {
			modus[i] = i
		}
	}

	// find most modus
	var mostModus float64
	for i := 0; i < len(modus); i++ {
		mostModus = math.Max(float64(mostModus), float64(modus[i]))
	}

	fmt.Print(int(mostModus))
}
