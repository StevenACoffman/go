package clock

import (
	"fmt"
)

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	return Clock{hour, minute}
}

func (c Clock) String() string {
	hours := c.hour
	minutes := c.minute

	totalMinutes := hours*60 + minutes + 14400 // HACK

	hours = totalMinutes / 60 % 24
	minutes = totalMinutes % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func (c Clock) Add(minutes int) Clock {
	hoursToAdd := minutes / 60
	minutesToAdd := minutes % 60
	return Clock{c.hour + hoursToAdd, c.minute + minutesToAdd}
}

func (c Clock) Subtract(minutes int) Clock {
	hoursToSubtract := minutes / 60
	minutesToSubtract := minutes % 60
	return Clock{c.hour - hoursToSubtract, c.minute - minutesToSubtract}
}
