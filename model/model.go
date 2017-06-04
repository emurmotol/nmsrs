package model

import (
	"github.com/emurmotol/nmsrs.v4/database"
	"github.com/emurmotol/nmsrs.v4/env"
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
		down()
		up()
		migrate()
		seed()
	}
}

func up() {
	db := database.Conn()
	defer db.Close()

	db.CreateTable(&User{})
}

func down() {
	db := database.Conn()
	defer db.Close()

	db.DropTableIfExists(&User{})
}

func migrate() {
	db := database.Conn()
	defer db.Close()

	db.AutoMigrate(&User{})
}

func seed() {
	go UserSeeder()
}
