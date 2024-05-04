package kata

func FindOutlier(integers []int) int {

	oddNumber := []int{}
	evenNumber := []int{}

	for number := range integers {
		if number > 0 {
			if number%2 == 1 {
				oddNumber = append(oddNumber, number)
			} else {
				evenNumber = append(evenNumber, number)
			}
		}
	}

	if len(oddNumber) > len(evenNumber) {
		return evenNumber[0]
	} else {
		return oddNumber[0]
	}
}
