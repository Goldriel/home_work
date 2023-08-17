package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Place your code here.
	newRune := []rune(str)
	var newString string
	var num int

	for i, char := range newRune {
		if unicode.IsDigit(char) && i == 0 {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(char) && unicode.IsDigit(newRune[i-1]) {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(char) {
			num, _ = strconv.Atoi(string(char))
			if num == 0 {
				newString = newString[:len(newString)-1]
				continue
			}
			for j := 0; j < num-1; j++ {
				newString += string(newRune[i-1])
			}
			continue
		}
		newString += string(char)
	}
	return newString, nil
}
