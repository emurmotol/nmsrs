package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Religion struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func religionSeeder() {
	data := []string{
		"AGLIPAYAN",
		"ASSEMBLY OF GOD",
		"BAPTIST",
		"BORN AGAIN CHRISTIAN",
		"BUDDIST",
		"CHRISTIAN",
		"CHURCH OF CHRIST",
		"CHURCH OF GOD THRU CHRIST JESUS",
		"CHURCH OF GOD",
		"CHURCH OF LATTER DAY SAINT",
		"EPISCOPALIAN ANGELICAN",
		"ESPIRITISM",
		"EVANGELICAL",
		"FAITH TABERNACLE",
		"FOUR SQUARE GOSPEL CHURCH",
		"FOURTH WATCH",
		"HINDU",
		"IGLESIA NG DIYOS KAY CRISTO JESUS",
		"IGLESIA NI CRISTO",
		"IGLESIA SA DIYOS ESPIRITU SANTO",
		"ISLAM",
		"JEHOVAH'S WITNESSES",
		"JESUS MIRACLE CRUSADE",
		"LUTHERAN",
		"METHODIST",
		"NON - SECTORAL CHARISMATIC",
		"ORTHODOX",
		"PENTECOSTAL",
		"PHILIPPINE INDEPENDENT CHRISTIAN CHURCH (PICC/IFI)",
		"PRESBYTERIAN",
		"PROTESTANT",
		"RIZALIST",
		"ROMAN CATHOLIC",
		"SEVENTH DAY ADVENTIST",
		"TWELVE TRIBES OF ISRAEL",
		"UNION ESPIRITISTA CRISTIANA",
		"UNITED CHURCH CHRISTIAN OF THE PHILIPPINES (UCCP)",
		"WESLEYAN CHURCH",
		"WORD OF HOPE",
	}

	for _, name := range data {
		religion := Religion{Name: strings.ToUpper(name)}
		religion.Create()
	}
}

func (religion *Religion) Create() *Religion {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&religion).Error; err != nil {
		panic(err)
	}
	return religion
}

func (religion Religion) Index(q string) []Religion {
	db := database.Conn()
	defer db.Close()

	religions := []Religion{}
	results := make(chan []Religion)

	go func() {
		db.Find(&religions, "name LIKE ?", database.WrapLike(q))
		results <- religions
	}()
	return <-results
}
