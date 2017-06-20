package model

import (
	"github.com/emurmotol/nmsrs/database"
)

type RegistEmp struct {
	ID           uint64    `json:"id"`
	RegistrantID uint64    `json:"registrant_id"`
	EmpStat      EmpStat   `gorm:"ForeignKey:EmpStatID"`
	EmpStatID    uint      `json:"emp_stat_id"`
	UnEmpStat    UnEmpStat `gorm:"ForeignKey:UnEmpStatID"`
	UnEmpStatID  uint      `json:"un_emp_stat_id"`
	Toc          Country   `gorm:"ForeignKey:TocID"`
	TocID        uint      `json:"toc_id"`
	Alfw         bool      `gorm:"type:tinyint(1)" json:"alfw"`
	PassportNo   string    `json:"passport_no"`
	Pned         string    `json:"pned"`
}

func (registEmp *RegistEmp) Create() *RegistEmp {
	db := database.Con()
	defer db.Close()

	if err := db.Create(&registEmp).Error; err != nil {
		panic(err)
	}
	return registEmp
}
