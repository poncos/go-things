package model

import (
	"time"
)

// Person holds the information of a person
type Person struct {
	Name    string
	Surname string
	Gender  int
	Dob     time.Time
}
