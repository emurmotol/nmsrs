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
	db.CreateTable(&EmpStat{})
	db.CreateTable(&UnEmpStat{})
	db.CreateTable(&Country{})
	db.CreateTable(&Emp{})
	db.Model(&Emp{}).AddForeignKey("registrant_id", "registrants(id)", "RESTRICT", "RESTRICT")
	db.Model(&Emp{}).AddForeignKey("es_id", "emp_stats(id)", "RESTRICT", "RESTRICT")
	db.Model(&Emp{}).AddForeignKey("ues_id", "un_emp_stats(id)", "RESTRICT", "RESTRICT")
	db.Model(&Emp{}).AddForeignKey("toc_id", "countries(id)", "RESTRICT", "RESTRICT")
	db.CreateTable(&Region{})
	db.CreateTable(&Province{})
	db.CreateTable(&CityMun{})
	db.CreateTable(&Barangay{})
	db.CreateTable(&Certificate{})
	db.CreateTable(&CivilStat{})
	db.CreateTable(&Course{})
	db.CreateTable(&Disability{})
	db.CreateTable(&EduLevel{})
	db.CreateTable(&Eligibility{})
	db.CreateTable(&Industry{})
	db.CreateTable(&Language{})
	db.CreateTable(&License{})
	db.CreateTable(&OtherSkill{})
	db.CreateTable(&Position{})
	db.CreateTable(&Religion{})
	db.CreateTable(&School{})
	db.CreateTable(&Sex{})
}

func down() {
	db := database.Conn()
	defer db.Close()

	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&Emp{})
	db.DropTableIfExists(&EmpStat{})
	db.DropTableIfExists(&UnEmpStat{})
	db.DropTableIfExists(&Country{})
	db.DropTableIfExists(&PeInfo{})
	db.DropTableIfExists(&Registrant{})
	db.DropTableIfExists(&Region{})
	db.DropTableIfExists(&Province{})
	db.DropTableIfExists(&CityMun{})
	db.DropTableIfExists(&Barangay{})
	db.DropTableIfExists(&Certificate{})
	db.DropTableIfExists(&CivilStat{})
	db.DropTableIfExists(&Course{})
	db.DropTableIfExists(&Disability{})
	db.DropTableIfExists(&EduLevel{})
	db.DropTableIfExists(&Eligibility{})
	db.DropTableIfExists(&Industry{})
	db.DropTableIfExists(&Language{})
	db.DropTableIfExists(&License{})
	db.DropTableIfExists(&OtherSkill{})
	db.DropTableIfExists(&Position{})
	db.DropTableIfExists(&Religion{})
	db.DropTableIfExists(&School{})
	db.DropTableIfExists(&Sex{})
}

func migrate() {
	db := database.Conn()
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&EmpStat{})
	db.AutoMigrate(&UnEmpStat{})
	db.AutoMigrate(&Country{})
	db.AutoMigrate(&Emp{})
	db.AutoMigrate(&Registrant{})
	db.AutoMigrate(&PeInfo{})
	db.AutoMigrate(&Region{})
	db.AutoMigrate(&Province{})
	db.AutoMigrate(&CityMun{})
	db.AutoMigrate(&Barangay{})
	db.AutoMigrate(&Certificate{})
	db.AutoMigrate(&CivilStat{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Disability{})
	db.AutoMigrate(&EduLevel{})
	db.AutoMigrate(&Eligibility{})
	db.AutoMigrate(&Industry{})
	db.AutoMigrate(&Language{})
	db.AutoMigrate(&License{})
	db.AutoMigrate(&OtherSkill{})
	db.AutoMigrate(&Position{})
	db.AutoMigrate(&Religion{})
	db.AutoMigrate(&School{})
	db.AutoMigrate(&Sex{})
}

func seed() {
	go UserSeeder()
	go CountrySeeder()
	go EmpStatSeeder()
	go UnEmpStatSeeder()
	go RegionSeeder()
	go ProvinceSeeder()
	go CityMunSeeder()
	go BarangaySeeder()
	go CertificateSeeder()
	go CivilStatSeeder()
	go CourseSeeder()
	go DisabilitySeeder()
	go EduLevelSeeder()
	go EligibilitySeeder()
	go IndustrySeeder()
	go LanguageSeeder()
	go LicenseSeeder()
	go OtherSkillSeeder()
	go PositionSeeder()
	go ReligionSeeder()
	go SchoolSeeder()
	go SexSeeder()
}
