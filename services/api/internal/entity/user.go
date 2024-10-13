package entity

import "time"

type UserID = uint64

type User struct {
	ID          UserID
	Username    string
	DisplayName string
	Description string
	AvatarURL   string
	IsBanned    bool
	BanReason   string
	AuthType    AuthType
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
