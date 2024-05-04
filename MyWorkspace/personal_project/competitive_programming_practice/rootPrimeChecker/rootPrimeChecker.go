package rootprimechecker

func RootPrimeChecker(N uint) bool {
	/*
		fact: prime number can divided by 1 or it's self only
		This algorythm will check N number by it's root, for ex: if N = 121 then sqrt(N) = 11,
		so, to have 121 from 11, we can iterate i times until sqrt(N)

		time complexity: O(sqrt(N))
	*/
	for i := uint(2); i*i <= N; i++ { // mean i <= sqrt(N)
		if N%i == 0 {
			return false
		}
	}

	return true
}
