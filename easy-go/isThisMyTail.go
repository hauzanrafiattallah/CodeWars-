package kata

func CorrectTail(body string, tail rune) bool {
	lastLetter := rune(body[len(body)-1])
	return lastLetter == tail
}
