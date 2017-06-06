package model

import "time"

type RegistEmp struct {
	ID           uint64    `json:"id"`
	RegistrantID uint64    `json:"registrant_id"`
	EmpStat      EmpStat   `gorm:"ForeignKey:EmpStatID"`
	EmpStatID    uint      `json:"emp_stat_id"`
	Alfw         bool      `gorm:"type:tinyint(1)" json:"alfw"`
	UnEmpStat    UnEmpStat `gorm:"ForeignKey:UnEmpStatID"`
	UnEmpStatID  uint      `json:"un_emp_stat_id"`
	Toc          Country   `gorm:"ForeignKey:TocID"`
	TocID        uint      `json:"toc_id"`
	PassportNo   string    `json:"passport_no"`
	Pned         time.Time `json:"pned"`
}
