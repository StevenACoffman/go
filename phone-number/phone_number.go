package phonenumber

import (
	"errors"
	"fmt"
)

// Number cleans up the string
func Number(in string) (string, error) {
	out := toDigits([]rune(in))
	maxLen := 10
	if len(out) < 10 {
		return "", errors.New("Number must be 10 digits long")
	}
	if out[0] == 48 {
		return "", errors.New("Number can't start with a 0")
	}
	if out[0] == 49 {
		out = out[1:]
	}
	if out[0] < 50 || out[3] < 50 {
		return "", errors.New("Area and Exchange codes can't be 1 or 0")
	}
	if len(out) != maxLen {
		return "", errors.New("Number must be 10 digits long")
	}

	return string(out), nil
}

func toDigits(in []rune) []rune {
	out := []rune{}

	for _, c := range in {
		// verify that c is 0-9
		if c >= 48 && c <= 57 {
			out = append(out, c)
		}
	}
	//fmt.Printf("toDigits: %s, %s, len = %d\n", string(in), string(out), len(out))
	return out
}

// AreaCode returns the
func AreaCode(in string) (string, error) {
	n, err := Number(in)
	if err != nil {
		return "", err
	}
	return string([]rune(n)[0:3]), nil
}

// Format make the number pretty ala (613) 995-0253
func Format(in string) (string, error) {
	number, err := Number(in)
	if err != nil {
		return "", err
	}
	n := []rune(number)
	return fmt.Sprintf("(%s) %s-%s", string(n[0:3]), string(n[3:6]),
		string(n[6:10])), nil
}
