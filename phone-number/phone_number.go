package phonenumber

import (
	"fmt"
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`\d`)

func Number(input string) (string, error) {
	digits := regex.FindAllString(input, -1)[:]
	if len(digits) == 11 {
		if digits[0] != "1" {
			return "", fmt.Errorf("Not 1")
		}
		digits = digits[1:]
	}
	if len(digits) != 10 {
		return "", fmt.Errorf("length")
	}
	if digits[0] == "0" || digits[0] == "1" || digits[3] == "0" || digits[3] == "1" {
		return "", fmt.Errorf("0 or 1")
	}
	return strings.Join(digits, ""), nil
}

func Format(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[0:3], number[3:6], number[6:10]), nil
}

func AreaCode(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}
	return number[0:3], nil
}
