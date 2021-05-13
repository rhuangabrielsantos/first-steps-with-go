package interation

func Repeater(character string) string {
	var repetitions string

	for i := 0; i < 5; i++ {
		repetitions += character
	}

	return repetitions
}
