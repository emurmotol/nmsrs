package religion

import (
	"log"
	"strings"
)

var data = []string{
	"TWELVE TRIBES OF ISRAEL",
	"ASSEMBLY OF GOD",
	"AGLIPAYAN",
	"BORN AGAIN CHRISTIAN",
	"BAPTIST",
	"BUDDIST",
	"CHURCH OF GOD THRU CHRIST JESUS",
	"CHRISTIAN",
	"CHURCH OF CHRIST",
	"CHURCH OF GOD",
	"EPISCOPALIAN ANGELICAN",
	"ESPIRITISM",
	"EVANGELICAL",
	"FOUR SQUARE GOSPEL CHURCH",
	"FAITH TABERNACLE",
	"HINDU",
	"IGLESIA SA DIYOS ESPIRITU SANTO",
	"IGLESIA NI CRISTO",
	"IGLESIA NG DIYOS KAY CRISTO JESUS",
	"ISLAM",
	"JESUS MIRACLE CRUSADE",
	"JEHOVAH'S WITNESSES",
	"LUTHERAN",
	"METHODIST",
	"CHURCH OF LATTER DAY SAINT",
	"NON - SECTORAL CHARISMATIC",
	"ORTHODOX",
	"PENTECOSTAL",
	"PHILIPPINE INDEPENDENT CHRISTIAN CHURCH (PICC/IFI)",
	"FOURTH WATCH",
	"PRESBYTERIAN",
	"PROTESTANT",
	"ROMAN CATHOLIC",
	"RIZALIST",
	"SEVENTH DAY ADVENTIST",
	"UNITED CHURCH CHRISTIAN OF THE PHILIPPINES (UCCP)",
	"UNION ESPIRITISTA CRISTIANA",
	"WESLEYAN CHURCH",
	"WORD OF HOPE",
}

func Seeder() {
	religs, err := All()

	if err != nil {
		panic(err)
	}

	if len(religs) == 0 {
		for index, value := range data {
			var relig Religion
			relig.ID = index + 1
			relig.Name = strings.ToUpper(value)
			_, err := relig.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("Religion seeded")
	}
}
