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
	db := database.Con()
	defer db.Close()

	db.CreateTable(&User{})
	db.CreateTable(&Region{})
	db.Model(&Region{}).AddIndex("index_reg_code", "code")
	db.CreateTable(&Province{})
	db.Model(&Province{}).AddIndex("index_prov_code", "code")
	db.Model(&Province{}).AddForeignKey("reg_code", "regions(code)", "CASCADE", "CASCADE")
	db.CreateTable(&CityMun{})
	db.Model(&CityMun{}).AddIndex("index_city_mun_code", "code")
	db.Model(&CityMun{}).AddForeignKey("reg_code", "regions(code)", "CASCADE", "CASCADE")
	db.Model(&CityMun{}).AddForeignKey("prov_code", "provinces(code)", "CASCADE", "CASCADE")
	db.CreateTable(&Barangay{})
	db.Model(&Barangay{}).AddIndex("index_brgy_code", "code")
	db.Model(&Barangay{}).AddForeignKey("reg_code", "regions(code)", "CASCADE", "CASCADE")
	db.Model(&Barangay{}).AddForeignKey("prov_code", "provinces(code)", "CASCADE", "CASCADE")
	db.Model(&Barangay{}).AddForeignKey("city_mun_code", "city_muns(code)", "CASCADE", "CASCADE")
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
	db.CreateTable(&Skill{})
	db.CreateTable(&Registrant{})
	db.CreateTable(&RegistInfo{})
	db.Model(&RegistInfo{}).AddForeignKey("registrant_id", "registrants(id)", "CASCADE", "CASCADE")
	db.Model(&RegistInfo{}).AddForeignKey("city_mun_id", "city_muns(id)", "CASCADE", "CASCADE")
	db.Model(&RegistInfo{}).AddForeignKey("prov_id", "provinces(id)", "CASCADE", "CASCADE")
	db.Model(&RegistInfo{}).AddForeignKey("brgy_id", "barangays(id)", "CASCADE", "CASCADE")
	db.Model(&RegistInfo{}).AddForeignKey("civil_stat_id", "civil_stats(id)", "CASCADE", "CASCADE")
	db.Model(&RegistInfo{}).AddForeignKey("sex_id", "sexes(id)", "CASCADE", "CASCADE")
	db.CreateTable(&EmpStat{})
	db.CreateTable(&UnEmpStat{})
	db.CreateTable(&Country{})
	db.CreateTable(&RegistEmp{})
	db.Model(&RegistEmp{}).AddForeignKey("registrant_id", "registrants(id)", "CASCADE", "CASCADE")
	db.Model(&RegistEmp{}).AddForeignKey("emp_stat_id", "emp_stats(id)", "CASCADE", "CASCADE")
	db.Model(&RegistEmp{}).AddForeignKey("un_emp_stat_id", "un_emp_stats(id)", "CASCADE", "CASCADE")
	db.Model(&RegistEmp{}).AddForeignKey("toc_id", "countries(id)", "CASCADE", "CASCADE")
}

func down() {
	db := database.Con()
	defer db.Close()

	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&RegistEmp{})
	db.DropTableIfExists(&EmpStat{})
	db.DropTableIfExists(&UnEmpStat{})
	db.DropTableIfExists(&Country{})
	db.DropTableIfExists(&RegistInfo{})
	db.DropTableIfExists(&Registrant{})
	db.DropTableIfExists(&Barangay{})
	db.DropTableIfExists(&CityMun{})
	db.DropTableIfExists(&Province{})
	db.DropTableIfExists(&Region{})
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
	db.DropTableIfExists(&Skill{})
}

func migrate() {
	db := database.Con()
	defer db.Close()

	db.AutoMigrate(&User{})
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
	db.AutoMigrate(&Skill{})
	db.AutoMigrate(&EmpStat{})
	db.AutoMigrate(&UnEmpStat{})
	db.AutoMigrate(&Country{})
	db.AutoMigrate(&RegistEmp{})
	db.AutoMigrate(&RegistInfo{})
	db.AutoMigrate(&Registrant{})
}

func seed() {
	go createSuperUser()
	go userSeeder()
	go countrySeeder()
	go empStatSeeder()
	go unEmpStatSeeder()
	go regionSeeder()
	go provinceSeeder()
	go cityMunSeeder()
	go barangaySeeder()
	go certificateSeeder()
	go civilStatSeeder()
	go courseSeeder()
	go disabilitySeeder()
	go eduLevelSeeder()
	go eligibilitySeeder()
	go industrySeeder()
	go languageSeeder()
	go licenseSeeder()
	go otherSkillSeeder()
	go positionSeeder()
	go religionSeeder()
	go schoolSeeder()
	go sexSeeder()
	go skillSeeder()
	// todo seed reg_ assoc last
}
