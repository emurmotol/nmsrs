package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Country struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func countrySeeder() {
	data := []string{
		"AFGHANISTAN",
		"ALBANIA",
		"ALGERIA",
		"ANDORRA",
		"ANGOLA",
		"ANTIGUA AND DEPS",
		"ARGENTINA",
		"ARMENIA",
		"AUSTRALIA",
		"AUSTRIA",
		"AZERBAIJAN",
		"BAHAMAS",
		"BAHRAIN",
		"BANGLADESH",
		"BARBADOS",
		"BELARUS",
		"BELGIUM",
		"BELIZE",
		"BENIN",
		"BHUTAN",
		"BOLIVIA",
		"BOSNIA HERZEGOVINA",
		"BOTSWANA",
		"BRAZIL",
		"BRUNEI",
		"BULGARIA",
		"BURKINA",
		"BURUNDI",
		"CAMBODIA",
		"CAMEROON",
		"CANADA",
		"CAPE VERDE",
		"CENTRAL AFRICAN REP",
		"CHAD",
		"CHILE",
		"CHINA",
		"COLOMBIA",
		"COMOROS",
		"CONGO",
		"CONGO (DEMOCRATIC REP)",
		"COSTA RICA",
		"CROATIA",
		"CUBA",
		"CYPRUS",
		"CZECH REPUBLIC",
		"DENMARK",
		"DJIBOUTI",
		"DOMINICA",
		"DOMINICAN REPUBLIC",
		"EAST TIMOR",
		"ECUADOR",
		"EGYPT",
		"EL SALVADOR",
		"EQUATORIAL GUINEA",
		"ERITREA",
		"ESTONIA",
		"ETHIOPIA",
		"FIJI",
		"FINLAND",
		"FRANCE",
		"GABON",
		"GAMBIA",
		"GEORGIA",
		"GERMANY",
		"GHANA",
		"GREECE",
		"GRENADA",
		"GUATEMALA",
		"GUINEA",
		"GUINEA - BISSAU",
		"GUYANA",
		"HAITI",
		"HONDURAS",
		"HUNGARY",
		"ICELAND",
		"INDIA",
		"INDONESIA",
		"IRAN",
		"IRAQ",
		"IRELAND (REPUBLIC)",
		"ISRAEL",
		"ITALY",
		"IVORY COAST",
		"JAMAICA",
		"JAPAN",
		"JORDAN",
		"KAZAKHSTAN",
		"KENYA",
		"KIRIBATI",
		"KOREA NORTH",
		"KOREA SOUTH",
		"KOSOVO",
		"KUWAIT",
		"KYRGYZSTAN",
		"LAOS",
		"LATVIA",
		"LEBANON",
		"LESOTHO",
		"LIBERIA",
		"LIBYA",
		"LIECHTENSTEIN",
		"LITHUANIA",
		"LUXEMBOURG",
		"MACEDONIA",
		"MADAGASCAR",
		"MALAWI",
		"MALAYSIA",
		"MALDIVES",
		"MALI",
		"MALTA",
		"MARSHALL ISLANDS",
		"MAURITANIA",
		"MAURITIUS",
		"MEXICO",
		"MICRONESIA",
		"MOLDOVA",
		"MONACO",
		"MONGOLIA",
		"MONTENEGRO",
		"MOROCCO",
		"MOZAMBIQUE",
		"MYANMAR, (BURMA)",
		"NAMIBIA",
		"NAURU",
		"NEPAL",
		"NETHERLANDS",
		"NEW ZEALAND",
		"NICARAGUA",
		"NIGER",
		"NIGERIA",
		"NORWAY",
		"OMAN",
		"PAKISTAN",
		"PALAU",
		"PANAMA",
		"PAPUA NEW GUINEA",
		"PARAGUAY",
		"PERU",
		"PHILIPPINES",
		"POLAND",
		"PORTUGAL",
		"QATAR",
		"ROMANIA",
		"RUSSIAN FEDERATION",
		"RWANDA",
		"ST KITTS AND NEVIS",
		"ST LUCIA",
		"SAINT VINCENT AND THE GRENADINES",
		"SAMOA",
		"SAN MARINO",
		"SAO TOME AND PRINCIPE",
		"SAUDI ARABIA",
		"SENEGAL",
		"SERBIA",
		"SEYCHELLES",
		"SIERRA LEONE",
		"SINGAPORE",
		"SLOVAKIA",
		"SLOVENIA",
		"SOLOMON ISLANDS",
		"SOMALIA",
		"SOUTH AFRICA",
		"SOUTH SUDAN",
		"SPAIN",
		"SRI LANKA",
		"SUDAN",
		"SURINAME",
		"SWAZILAND",
		"SWEDEN",
		"SWITZERLAND",
		"SYRIA",
		"TAIWAN",
		"TAJIKISTAN",
		"TANZANIA",
		"THAILAND",
		"TOGO",
		"TONGA",
		"TRINIDAD AND TOBAGO",
		"TUNISIA",
		"TURKEY",
		"TURKMENISTAN",
		"TUVALU",
		"UGANDA",
		"UKRAINE",
		"UNITED ARAB EMIRATES",
		"UNITED KINGDOM",
		"UNITED STATES",
		"URUGUAY",
		"UZBEKISTAN",
		"VANUATU",
		"VATICAN CITY",
		"VENEZUELA",
		"VIETNAM",
		"YEMEN",
		"ZAMBIA",
		"ZIMBABWE",
	}

	for _, name := range data {
		country := Country{Name: strings.ToUpper(name)}
		country.Create()
	}
}

func (country *Country) Create() *Country {
	db := database.Con()
	defer db.Close()

	if err := db.Create(&country).Error; err != nil {
		panic(err)
	}
	return country
}

func (country Country) Index(q string) []Country {
	db := database.Con()
	defer db.Close()

	countries := []Country{}
	results := make(chan []Country)

	go func() {
		db.Find(&countries, "name LIKE ?", database.WrapLike(q))
		results <- countries
	}()

	countries = <-results
	close(results)
	return countries
}
