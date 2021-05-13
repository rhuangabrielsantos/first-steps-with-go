package interation

func Repeater(character string, repetitions int) string {
	var result string

	for i := 0; i < repetitions; i++ {
		result += character
	}

	return result
}
