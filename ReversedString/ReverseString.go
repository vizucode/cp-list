package reversedstring

func ReverseString(word string) (resp string) {

	for i := len(word) - 1; i >= 0; i-- {
		resp += string(word[i])
	}

	return resp
}
