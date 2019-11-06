package phonenumber

import (
	"strings"
	"errors"
)

func Number(input string) (string, error) {
	output := strings.Replace(input, " ", "", -1)
	output = strings.Replace(output, "(", "", -1)
	output = strings.Replace(output, ")", "", -1)
	output = strings.Replace(output, ".", "", -1)
	output = strings.Replace(output, "-", "", -1)
	output = strings.Replace(output, "+", "", -1)

	if len(output) > 10 {
		output = output[1:]
	}
	return output, nil
}

func Format(input string) (string, error) {
	initial, err := Number(input)
	if err != nil {
		return "", err
	}
	formatted := "(" + initial[0:3] + ") " +initial[3:6] + "-" + initial[6:]
	return formatted, nil
}

func AreaCode(input string) (string, error) {
	number, errNumber := Number(input)

	if errNumber != nil {
		return "", errNumber
	}

	if len(number) != 10 {
		return "", errors.New("error")
	}

	areaCode := number[0:3]
	return areaCode, nil
}