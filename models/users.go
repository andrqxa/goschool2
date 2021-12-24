package models

import (
	"fmt"
	"time"
)

type Date struct {
	day   int
	month int
	year  int
}

func NewDate() Date {
	t := time.Now()
	return Date{
		day:   t.Day(),
		month: int(t.Month()),
		year:  t.Year(),
	}
}

func (d Date) String() string {
	return fmt.Sprintf("%d.%d.%d", d.day, d.month, d.year)
}

// User represents the user model.
type User struct {
	Name             string
	Email            string
	RegistrationDate Date
}
