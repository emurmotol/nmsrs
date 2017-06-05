package model

import (
	"os"
	"path/filepath"

	"github.com/emurmotol/nmsrs/database"
	"github.com/emurmotol/nmsrs/env"
)

var (
	count          int
	contentDir     string
	SuperuserEmail string
)

func init() {
	contentDir, _ = env.Conf.String("dir.content")
	SuperuserEmail, _ = env.Conf.String("superuser.email")
}

func Load(reset bool) {
	if reset {
		clearContentDir()
		down()
		up()
		migrate()
		seed()
	}
}

func clearContentDir() {
	dir := filepath.Join(contentDir)

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		if err := os.RemoveAll(dir); err != nil {
			panic(err)
		}
	}
}

func up() {
	db := database.Conn()
	defer db.Close()

	db.CreateTable(&User{})
	db.CreateTable(&Registrant{})
	db.CreateTable(&PeInfo{})
	db.Model(&PeInfo{}).AddForeignKey("registrant_id", "registrants(id)", "RESTRICT", "RESTRICT")
	db.CreateTable(&Stat{})
	db.CreateTable(&UnempStat{})
	db.CreateTable(&Country{})
	db.CreateTable(&Emp{})
	db.Model(&Emp{}).AddForeignKey("registrant_id", "registrants(id)", "RESTRICT", "RESTRICT")
	db.Model(&Emp{}).AddForeignKey("stat_id", "stats(id)", "RESTRICT", "RESTRICT")
	db.Model(&Emp{}).AddForeignKey("unemp_stat_id", "unemp_stats(id)", "RESTRICT", "RESTRICT")
	db.Model(&Emp{}).AddForeignKey("toc_id", "countries(id)", "RESTRICT", "RESTRICT")
}

func down() {
	db := database.Conn()
	defer db.Close()

	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&Emp{})
	db.DropTableIfExists(&Stat{})
	db.DropTableIfExists(&UnempStat{})
	db.DropTableIfExists(&Country{})
	db.DropTableIfExists(&PeInfo{})
	db.DropTableIfExists(&Registrant{})
}

func migrate() {
	db := database.Conn()
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Stat{})
	db.AutoMigrate(&UnempStat{})
	db.AutoMigrate(&Country{})
	db.AutoMigrate(&Emp{})
	db.AutoMigrate(&Registrant{})
	db.AutoMigrate(&PeInfo{})
}

func seed() {
	go UserSeeder()
}
