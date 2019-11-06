package clock

import "fmt"

type Clock struct {
	Hour   int
	Minute int
}

func New(hour, minute int) Clock {
	hourCarry, minute := normalize(minute, 60)
	_, hour = normalize(hour+hourCarry, 24)
	return Clock{hour, minute}
}

func (c Clock) Add(minute int) Clock {
	return New(c.Hour, c.Minute+minute)
}

func normalize(val int, max int) (int, int) {
	cary := val / max
	m := val % max
	if m < 0 {
		cary--
		m += max
	}
	return cary, m
}

func (c Clock) Subtract(minute int) Clock {
	return New(c.Hour, c.Minute-minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)

}
