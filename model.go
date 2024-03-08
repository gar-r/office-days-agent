package main

import "time"

type Loc int

const (
	Office Loc = iota
	Home
)

type Workday struct {
	Date     time.Time
	Location Loc
}
