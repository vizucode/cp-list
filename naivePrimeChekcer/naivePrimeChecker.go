package naiveprimechekcer

func NaivePrimeChecker(N uint) bool {
	/*
		fact: prime number can divided by 1 or it's self only
		In this way i loop over N then check if the range 2 < N is can't divide N

		So the worst case was loop over N times, or O(N)
	*/
	for i := uint(2); i < N; i++ {
		if N%i == 0 {
			return false
		}
	}

	return true
}
