package models

import "github.com/gofrs/uuid"

type User struct {
	UserID   uuid.UUID `gorm:"type:uuid;primary_key"`
	Email    string    `gorm:"type:varchar(100);unique_index"`
	Password []byte    `gorm:"type:bytea"`
	Name     string    `gorm:"type:varchar(100)"`
	Age      int       `gorm:"type:int"`
	Address  string    `gorm:"type:varchar(100)"`
	Gender   string    `gorm:"type:varchar(100)"`
	Bio      string    `gorm:"type:varchar(100)"`
}
