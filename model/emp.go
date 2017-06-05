package model

import "time"

type Emp struct {
	ID           int64 `json:"id"`
	RegistrantID int64 `json:"registrant_id"`
	Stat         Stat
	StatID       int  `json:"stat_id"`
	Alfw         bool `gorm:"type:tinyint(1)" json:"alfw"`
	UnempStat    UnempStat
	UnempStatID  int       `json:"unemp_stat_id"`
	Toc          Country   `gorm:"ForeignKey:TocID"`
	TocID        int       `json:"toc_id"`
	PassportNo   string    `json:"passport_no"`
	Pned         time.Time `json:"pned"`
}
