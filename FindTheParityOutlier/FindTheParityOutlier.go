package kata

func FindOutlier(integers []int) int {

	oddNumber := []int{}
	evenNumber := []int{}

	for _, number := range integers {
		if number%2 == 1 || number%2 == -1 {
			oddNumber = append(oddNumber, number)
		} else {
			evenNumber = append(evenNumber, number)
		}
	}

	if len(oddNumber) > len(evenNumber) {
		if len(evenNumber) > 0 {
			return evenNumber[0]
		}
	} else {
		if len(oddNumber) > 0 {
			return oddNumber[0]
		}
	}

	return 0
}
