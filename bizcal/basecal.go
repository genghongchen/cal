package bizcal

import (
	"time"
)

//BaseCal interface, only checks weekday vs weekend
type BaseCal interface {
	IsWeekday(t time.Time) bool
	IsWeekend(t time.Time) bool
}

//Basic Calendar satisfies BaseCal interface
type BasicCal struct{}

//IsWeekend checks if a particular day is weekend
func (cal BasicCal) IsWeekend(t time.Time) bool {
	w := t.Weekday()
	if w == time.Saturday || w == time.Sunday {
		return true
	}

	return false
}

//IsWeekday checks if a particular day is a weekday
func (cal BasicCal) IsWeekday(t time.Time) bool {
	return !(cal.IsWeekend(t))
}
