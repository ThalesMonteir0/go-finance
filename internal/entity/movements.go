package entity

import "time"

type Movements struct {
	ID    int
	Type  Types
	Value float64
	Desc  string
	Date  time.Time
}
