package database

import (
	"fmt"

	"github.com/emurmotol/nmsrs.v4/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	args    string
	dialect string
	logMode bool
)

func init() {
	user, _ := env.Conf.String("pkg.gorm.user")
	pwd, _ := env.Conf.String("pkg.gorm.pwd")
	name, _ := env.Conf.String("pkg.gorm.name")
	charset, _ := env.Conf.String("pkg.gorm.charset")
	parseTime, _ := env.Conf.Bool("pkg.gorm.parseTime")
	loc, _ := env.Conf.String("pkg.gorm.loc")
	args = fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=%t&loc=%s", user, pwd, name, charset, parseTime, loc)
	dialect, _ = env.Conf.String("pkg.gorm.dialect")
	logMode, _ = env.Conf.Bool("pkg.gorm.logMode")
}

func Conn() *gorm.DB {
	db, err := gorm.Open(dialect, args)
	db.LogMode(logMode)

	if err != nil {
		panic(err)
	}
	return db
}

func WrapLike(q string) string {
	return "%" + q + "%"
}
