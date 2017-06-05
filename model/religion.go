package model

import "github.com/emurmotol/nmsrs/database"

type Religion struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ReligionSeeder() {
	data := []string{
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

	for _, name := range data {
		religion := Religion{Name: name}

		if _, err := religion.Create(); err != nil {
			panic(err)
		}
	}
}

func (religion *Religion) Create() (*Religion, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&religion).Error; err != nil {
		return nil, err
	}
	return religion, nil
}

func (religion Religion) Search(q string) []Religion {
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
