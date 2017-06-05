package model

import "time"

type Emp struct {
	ID           int64     `json:"id"`
	RegistrantID int64     `json:"registrant_id"`
	EmpStat      EmpStat   `gorm:"ForeignKey:EmpStatID"`
	EmpStatID    int       `gorm:"column:es_id" json:"es_id"`
	Alfw         bool      `gorm:"type:tinyint(1)" json:"alfw"`
	UnEmpStat    UnEmpStat `gorm:"ForeignKey:UnEmpStatID"`
	UnEmpStatID  int       `gorm:"column:ues_id" json:"ues_id"`
	Toc          Country   `gorm:"ForeignKey:TocID"`
	TocID        int       `json:"toc_id"`
	PassportNo   string    `json:"passport_no"`
	Pned         time.Time `json:"pned"`
}
