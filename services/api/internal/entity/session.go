package entity

import "time"

type SessionID = uint64

type Session struct {
	ID           SessionID
	Token        string
	AccessSource AccessSource
	CreatedAt    time.Time
}

type AccessSource struct {
	OS       string
	Platform string
	City     string
	Country  string
}
