package model

import (
	"time"
)

type RegistInfo struct {
	ID             uint64    `json:"id"`
	RegistrantID   uint64    `json:"registrant_id"`
	FamilyName     string    `gorm:"not null" json:"family_name"`
	GivenName      string    `gorm:"not null" json:"given_name"`
	MiddleName     string    `gorm:"not null" json:"middle_name"`
	Birthdate      time.Time `json:"birthdate"`
	Password       string    `gorm:"not null" json:"password"`
	HasPhoto       bool      `gorm:"type:tinyint(1);default:false;not null" json:"has_photo"`
	StSub          string    `gorm:"not null" json:"st_sub"`
	CityMun        CityMun   `gorm:"ForeignKey:CityMunCode;AssociationForeignKey:Code"`
	CityMunCode    string    `json:"city_mun_code"`
	Province       Province  `gorm:"ForeignKey:ProvCode;AssociationForeignKey:Code"`
	ProvCode       string    `json:"prov_code"`
	Barangay       Barangay  `gorm:"ForeignKey:BrgyCode;AssociationForeignKey:Code"`
	BrgyCode       string    `json:"brgy_code"`
	CivilStat      CivilStat `gorm:"ForeignKey:CivilStatID"`
	CivilStatID    uint      `json:"civil_stat_id"`
	CivilStatOther string    `json:"civil_stat_other"`
	Sex            Sex       `gorm:"ForeignKey:SexID"`
	SexID          uint      `json:"sex_id"`
	Age            int       `json:"age"`
	Height         float32   `json:"height"`
	Weight         float32   `json:"weight"`
}
