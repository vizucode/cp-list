package countoddnumbern

func CountOddNumberN(N int) (resp int) {

	for i := 1; i < N; i++ {
		if i%2 == 1 {
			resp++
		}
	}

	return resp
}
