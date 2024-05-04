package primegenerationeratothenes

/*
	fact: prime number was only can divided by 1 or it's self
	generate prime number that we eliminate  from the non prime number on N.
	so, to check if we had non prime number, we can multiply the point number to N

	Complexity: O(N(SQRT(N)))
	Bech time: 0.433s
*/

func PrimeGenerationEratosthenes(N uint) []uint {
	var primeList []uint
	blackListNumber := make([]bool, N+1) // first time will false all

	for i := uint(2); i <= N; i++ {
		if !blackListNumber[i] {
			primeList = append(primeList, i)
			mul := i * i
			for mul <= N {
				blackListNumber[mul] = true
				mul += i
			}
		}
	}

	return primeList
}
