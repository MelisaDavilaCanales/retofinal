package models

import "time"

type Session struct {
	Uid        uint
	ExpiryTime time.Time
}
