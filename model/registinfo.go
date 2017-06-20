package model

import (
	"github.com/emurmotol/nmsrs/database"
)

type RegistInfo struct {
	ID             uint64    `json:"id"`
	RegistrantID   uint64    `json:"registrant_id"`
	FamilyName     string    `gorm:"not null" json:"family_name"`
	GivenName      string    `gorm:"not null" json:"given_name"`
	MiddleName     string    `gorm:"not null" json:"middle_name"`
	Birthdate      string    `json:"birthdate"`
	Password       string    `gorm:"not null" json:"password"`
	HasPhoto       bool      `gorm:"type:tinyint(1);default:false;not null" json:"has_photo"`
	StSub          string    `gorm:"not null" json:"st_sub"`
	CityMun        CityMun   `gorm:"ForeignKey:CityMunID"`
	CityMunID      uint      `json:"city_mun_id"`
	Province       Province  `gorm:"ForeignKey:ProvID"`
	ProvID         uint      `json:"prov_id"`
	Barangay       Barangay  `gorm:"ForeignKey:BrgyID"`
	BrgyID         uint      `json:"brgy_id"`
	CivilStat      CivilStat `gorm:"ForeignKey:CivilStatID"`
	CivilStatID    uint      `json:"civil_stat_id"`
	CivilStatOther string    `json:"civil_stat_other"`
	Sex            Sex       `gorm:"ForeignKey:SexID"`
	SexID          uint      `json:"sex_id"`
	Age            int       `json:"age"`
	Height         float32   `json:"height"`
	Weight         float32   `json:"weight"`
	LandlineNo     string    `json:"landline_no"`
	MobileNo       string    `json:"mobile_no"`
	Email          string    `json:"email"`
}

func (registInfo *RegistInfo) Create() *RegistInfo {
	db := database.Con()
	defer db.Close()

	if err := db.Create(&registInfo).Error; err != nil {
		panic(err)
	}
	return registInfo
}
