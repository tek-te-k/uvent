package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Category            string `gorm:"type:varchar(100)"`
	Participations      []Participation
	UserID              uint
	Title               string    `gorm:"type:varchar(100)"`
	Detail              string    `gorm:"type:varchar(100)"`
	Hold                time.Time `gorm:"type:date"`
	Place               string    `gorm:"type:varchar(100)"`
	ApplicationDeadline time.Time `gorm:"type:timestamp"`
	ImageURL            string    `gorm:"type:varchar(100)"`
}
