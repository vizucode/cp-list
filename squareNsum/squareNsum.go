package squarensum

func SquareNSum(N []int) int {
	var sum int
	for _, n := range N {
		sum += n * n
	}

	return sum
}
