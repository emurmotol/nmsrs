package model

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/emurmotol/nmsrs/database"
	"github.com/emurmotol/nmsrs/env"
)

var (
	count          int
	contentDir     string
	SuperUserEmail string
	db             = database.Con()
)

func init() {
	contentDir, _ = env.Conf.String("dir.content")
	SuperUserEmail, _ = env.Conf.String("superuser.email")
}

func Load(reset bool) {
	if reset {
		drop()
		go clearContentDir()
		seed()
	}
}

func drop() {
	cNames, _ := db.CollectionNames()

	for _, cName := range cNames {
		if cName == "system.indexes" {
			continue
		}

		if err := db.C(cName).DropCollection(); err != nil {
			panic(err)
		}
	}
}

func clearContentDir() {
	if _, err := os.Stat(contentDir); !os.IsNotExist(err) {
		if err := os.RemoveAll(contentDir); err != nil {
			panic(err)
		}
	}
}

func seed() {
	createSuperUser()
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
