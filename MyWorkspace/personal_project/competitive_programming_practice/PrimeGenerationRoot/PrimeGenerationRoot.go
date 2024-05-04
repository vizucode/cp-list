package primegenerationroot

/*
	fact: prime number was only can divided by 1 or it's self
	prime generation, generate N prime number

	Algorythm: Simple Prime Generation Using Root Checker
	Complexity: O(N(sqrt(N)))

	benchmark test: 0.472s
*/

func PrimeGenerationChecker(N uint) bool {
	for i := uint(2); i*i <= N; i++ {
		if N%i == 0 {
			return false
		}
	}
	return true
}

func PrimeGenerationRoot(N uint) []uint {
	var primeList []uint

	for i := uint(2); i <= N; i++ {
		if PrimeGenerationChecker(i) {
			primeList = append(primeList, i)
		}
	}

	return primeList
}
