package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

	"gopkg.in/mgo.v2/bson"
)

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
		religion := Religion{
			Id:   bson.NewObjectId(),
			Name: strings.ToUpper(name),
		}
		religion.Create()
	}
}

type Religion struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func (religion *Religion) Create() *Religion {
	db.C("religions").Insert(religion)
	defer db.Close()
	return religion
}

func (religion Religion) Index(q string) []Religion {
	religions := []Religion{}
	r := make(chan []Religion)
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	go func() {
		db.C("religions").Find(query).All(&religions)
		defer db.Close()
		r <- religions
	}()

	religions = <-r
	close(r)
	return religions
}
