package entity

import "time"

type SessionID = uint64

type AuthType uint

const (
	AuthTypeSSO AuthType = iota
	AuthTypeOIDC
)

type Session struct {
	ID          SessionID
	Token       string
	Fingerprint Fingerprint
	CreatedAt   time.Time
}

type Fingerprint struct {
	OS        string
	Platform  string
	City      string
	Country   string
	IPAddress string
}
