package clock

import "fmt"

type Clock struct {
	minutes int
}

const DAY = 60 * 24

func New(hour, minute int) Clock {
	return Clock{realMod(hour * 60 + minute, DAY)}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", realMod(c.minutes / 60, 24), realMod(c.minutes, 60))
}

func (c Clock) Add(minutes int) Clock {
	return Clock{realMod(c.minutes + minutes, DAY)}
}

func (c Clock) Subtract(minutes int) Clock {
	return Clock{realMod(c.minutes - minutes, DAY)}
}

func realMod(a, m int) int {
	if a < 0 {
		return a % m + m
	} else {
		return a % m
	}
}
