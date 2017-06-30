package model

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
		go drop()
		go clearContentDir()
		seed()
	}
}

func drop() {
	cNames, _ := db.Session().DB(db.Name).CollectionNames()

	for _, cName := range cNames {
		if cName == "system.indexes" {
			continue
		}

		if err := db.Session().DB(db.Name).C(cName).DropCollection(); err != nil {
			panic(err)
		}
	}
	defer db.Close()
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

	go func() {
		importDir, _ := env.Conf.String("dir.import")

		err := filepath.Walk(importDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				if filepath.Ext(path) == ".json" {
					args := []string{
						"-d", db.Name,
						"-c", strings.TrimSuffix(info.Name(), filepath.Ext(path)),
						"--jsonArray", path,
					}

					if err := exec.Command("mongoimport", args...).Run(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			panic(err)
		}
	}()
}
