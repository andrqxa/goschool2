package models

import (
	"fmt"
	"time"
)

type date struct {
	day   int
	month int
	year  int
}

func newDate() date {
	t := time.Now()
	return date{
		day:   t.Day(),
		month: int(t.Month()),
		year:  t.Year(),
	}
}

func (d date) String() string {
	return fmt.Sprintf("%d.%d.%d", d.day, d.month, d.year)
}

// User represents the user model.
type User struct {
	Name             string
	Email            string
	RegistrationDate date
}
