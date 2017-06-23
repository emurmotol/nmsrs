package model

import (
	"os"
	"path/filepath"

	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/env"
)

var (
	count          int
	contentDir     string
	SuperUserEmail string
)

func init() {
	contentDir, _ = env.Conf.String("dir.content")
	SuperUserEmail, _ = env.Conf.String("superuser.email")
}

func Load(reset bool) {

	if reset {
		drop()
		clearContentDir()
		seed()
	}
}

func drop() {
	cNames, _ := db.Session().DB(db.Name).CollectionNames()
	dNames := []string{"users"}

	for _, cName := range cNames {
		for _, dName := range dNames {
			if dName == cName {
				if err := db.Session().DB(db.Name).C(dName).DropCollection(); err != nil {
					panic(err)
				}
				defer db.Close()
			}
		}
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
