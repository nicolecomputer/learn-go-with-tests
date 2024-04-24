package iteration

func Repeat(character string) string {
	var repeatedStr string
	for i := 0; i < 5; i++ {
		repeatedStr = repeatedStr + character
	}
	return repeatedStr
}
