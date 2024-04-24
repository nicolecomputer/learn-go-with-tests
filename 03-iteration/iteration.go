package iteration

func Repeat(character string) string {
	return RepeatWithCount(character, 5)
}

func RepeatWithCount(character string, repeatCount int) string {
	var repeatedStr string
	for i := 0; i < repeatCount; i++ {
		repeatedStr = repeatedStr + character
	}
	return repeatedStr
}
