package phonenumber

import "fmt"

func Number(in string) (string, error) {
	out := ""
	//fmt.Printf("in %v\n", in)
	for _, l := range in {
		//fmt.Printf("comparing %v to %v", l, '0')
		if l >= '0' && l <= '9' {
			//fmt.Printf("Adding %v\n", string(l))
			out = out + string(l)
		} else {
			//fmt.Printf("Excluding %v\n", string(l))
		}
	}
	if len(out) == 11 && out[0] == '1' {
		out = out[1:]
	}
	if len(out) != 10 {
		return "", fmt.Errorf("error")
	}
	if out[0] == '0' || out[0] == '1' {
		return "", fmt.Errorf("error")
	}
	if out[3] == '0' || out[3] == '1' {
		return "", fmt.Errorf("error")
	}
	return out, nil
}

func AreaCode(in string) (string, error) {
	number, err := Number(in)
	if err != nil {
		return "", err
	}
	return number[0:3], nil
}

func Format(in string) (string, error) {
	number, err := Number(in)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[0:3], number[3:6], number[6:]), nil
}
