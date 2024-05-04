package binarycount

func BinaryCount(N uint) uint {
	var sumOfBiner uint

	/*
		loop over the N
		divide N with 2, then if the gte 0 than n mod 2
		if the result is 1 then add sumOfBiner
	*/
	for i := N; i > 0; i /= 2 {
		if i%2 == 1 {
			sumOfBiner++
		}
	}

	return sumOfBiner
}
