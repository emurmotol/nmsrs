package model

import (
	"time"
)

type RegistInfo struct {
	ID           int64     `json:"id"`
	RegistrantID int64     `json:"registrant_id"`
	FamilyName   string    `gorm:"not null" json:"family_name"`
	GivenName    string    `gorm:"not null" json:"given_name"`
	MiddleName   string    `gorm:"not null" json:"middle_name"`
	Birthdate    time.Time `json:"birthdate"`
	Password     string    `gorm:"not null" json:"password"`
	HasPhoto     bool      `gorm:"type:tinyint(1);default:false;not null" json:"has_photo"`
}
