package phonenumber

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func Number(uglyNumber string) (string, error) {
	// Remove anything that's not a number
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	numbersOnly := reg.ReplaceAllString(uglyNumber, "")

	if len(numbersOnly) > 11 {
		return "", errors.New(fmt.Sprintf("%s is too long", numbersOnly))
	} else if len(numbersOnly) < 10 {
		return "", errors.New(fmt.Sprintf("%s is too short", numbersOnly))
	}

	tenDigitNumber := numbersOnly
	if len(numbersOnly) == 11 {
		if numbersOnly[0] == '1' {
			tenDigitNumber = numbersOnly[1:]
		} else {
			return "", errors.New(fmt.Sprintf("%s is 11 digits long but country code is not 1", numbersOnly))
		}
	}

	char0, err := strconv.Atoi(string(tenDigitNumber[0]))
	char3, err := strconv.Atoi(string(tenDigitNumber[3]))
	if char0 < 2 || char3 < 2 {
		return "", errors.New(fmt.Sprintf("%s is holding an invalid digit in position 0 or 3", tenDigitNumber))
	}

	return tenDigitNumber, nil
}

func Format(uglyNumber string) (string, error) {
	cleanNumber, err := Number(uglyNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", cleanNumber[0:3], cleanNumber[3:6], cleanNumber[6:]), nil
}

func AreaCode(uglyNumber string) (string, error) {
	cleanNumber, err := Number(uglyNumber)
	if err != nil {
		return "", err
	}
	return cleanNumber[0:3], nil
}
