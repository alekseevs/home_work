package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"strconv"
	"os"
	"strings"
	"regexp"
)

var ErrInvalidString = errors.New("invalid string")
var LettersCountRegexp = regexp.MustCompile(`(?P<letter>[\p{L}])(?P<count>[0-9]*)?`)
const MaxLettersCount int = 9


func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	if !LettersCountRegexp.MatchString(s) {
		return "", ErrInvalidString
	}

	matches := LettersCountRegexp.FindAllStringSubmatch(s, -1)
	// fmt.Println(s + "\n")
	// fmt.Println(matches)
	var builder strings.Builder

	for _, match := range matches {
		matchedLetter, rawCount := match[1], match[2]

		count, err := strconv.Atoi(rawCount)
		if err != nil && len(rawCount) == 0 {
			count = 1
		}

		if count == 0 || count > MaxLettersCount {
			return "", ErrInvalidString
		}

		repeatedString := strings.Repeat(matchedLetter, count)
		builder.WriteString(repeatedString)
	}

	return builder.String(), nil
}



// func digit_checker(symbol string) bool{
// 	if _, err := strconv.Atoi(symbol); err == nil {
// 		// fmt.Printf("%q looks like a number.\n", symbol)
// 		return true
// 	} else {
// 		return false
// 	}

// }

// func string_to_int_conventer(s string) int{
// 	i, err := strconv.Atoi(s)
// 	if err != nil {
// 		// handle error
// 		fmt.Println(err)
// 		os.Exit(2)
// 	}
// 	return i
// }