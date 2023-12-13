package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if _, err := strconv.Atoi(str); err == nil {
		return "", ErrInvalidString
	}

	if unicode.IsDigit(rune(str[0])) {
		return "", ErrInvalidString
	}

	var prev rune
	var builder strings.Builder

	for i, char := range str {
		if unicode.IsDigit(char) {
			if unicode.IsDigit(rune(str[i-1])) {
				return "", ErrInvalidString
			}
			m := int(char - '0')

			if m == 0 {
				result := builder.String()
				result = result[:len(result)-1]
				builder.Reset()
				builder.WriteString(result)
				continue
			}

			r := strings.Repeat(string(prev), m-1)
			builder.WriteString(r)
		} else {
			builder.WriteRune(char)
			prev = char
		}
	}

	return builder.String(), nil
}
