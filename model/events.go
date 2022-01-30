package model

import (
	"time"
)

type Event struct {
	Domain    string
	Timestamp time.Time
	Path      string
	Referer   string
	Width     int32
	Height    int32
}
