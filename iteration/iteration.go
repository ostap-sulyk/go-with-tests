package iteration


func Repeat(character string, repeatCount int) string {
	var repeted string
	for i := 0; i < repeatCount; i++ {
		repeted += character
	}
	return repeted
}
