package clock

import "fmt"

type Clock struct {
	Hour   int
	Minute int
}

func New(hour, minute int) Clock {
	return Clock{
		Hour:   hour,
		Minute: minute,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%v:%v", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	return c
}

func (c Clock) Subtract(minutes int) Clock {
	return c
}
