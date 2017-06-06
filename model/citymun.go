package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type CityMun struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Desc     string `json:"desc"`
	PsgcCode string `json:"psgc_code"`
	RegCode  string `json:"reg_code"`
	ProvCode string `json:"prov_code"`
}

type CityMunProv struct {
	CityMunCode string `json:"city_mun_code"`
	CityMunDesc string `json:"city_mun_desc"`
	ProvDesc    string `json:"prov_desc"`
}

type RefCityMun struct {
	PsgcCode    string `json:"psgcCode"`
	CityMunDesc string `json:"cityMunDesc"`
	RegCode     string `json:"regCode"`
	ProvCode    string `json:"provCode"`
	CityMunCode string `json:"cityMunCode"`
}

func cityMunSeeder() {
	data, err := ioutil.ReadFile("model/data/refcitymun.json")

	if err != nil {
		panic(err)
	}
	refCityMuns := []RefCityMun{}

	if err := json.Unmarshal(data, &refCityMuns); err != nil {
		panic(err)
	}

	for _, refCityMun := range refCityMuns {
		cityMun := CityMun{
			Code:     refCityMun.CityMunCode,
			Desc:     strings.ToUpper(refCityMun.CityMunDesc),
			PsgcCode: refCityMun.PsgcCode,
			RegCode:  refCityMun.RegCode,
			ProvCode: refCityMun.ProvCode,
		}
		cityMun.Create()
	}
}

func (cityMun *CityMun) Create() *CityMun {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&cityMun).Error; err != nil {
		panic(err)
	}
	return cityMun
}

func (cityMun CityMun) ProvinceIndex(q string) []CityMunProv {
	db := database.Conn()
	defer db.Close()
	cityMunProv := []CityMunProv{}
	results := make(chan []CityMunProv)

	go func() {
		if err := db.Raw("SELECT city_muns.code as city_mun_code, city_muns.desc as city_mun_desc, provinces.desc as prov_desc, CONCAT(city_muns.desc, ', ', provinces.desc) AS name FROM provinces INNER JOIN city_muns ON city_muns.prov_code = provinces.code HAVING name LIKE ?", database.WrapLike(q)).Scan(&cityMunProv).Error; err != nil {
			panic(err)
		}
		results <- cityMunProv
	}()
	return <-results
}

func (cityMun *CityMun) BarangayIndex(q string) []Barangay {
	db := database.Conn()
	defer db.Close()
	barangays := []Barangay{}
	results := make(chan []Barangay)

	go func() {
		if err := db.Where("city_mun_code = ? AND `desc` LIKE ?", cityMun.Code, database.WrapLike(q)).Find(&barangays).Error; err != nil {
			panic(err)
		}
		results <- barangays
	}()
	return <-results
}
