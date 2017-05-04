package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emurmotol/nmsrs/env"
	"github.com/emurmotol/nmsrs/models/barangay"
	"github.com/emurmotol/nmsrs/models/certificate"
	"github.com/emurmotol/nmsrs/models/citymunicipality"
	"github.com/emurmotol/nmsrs/models/civilstatus"
	"github.com/emurmotol/nmsrs/models/country"
	"github.com/emurmotol/nmsrs/models/course"
	"github.com/emurmotol/nmsrs/models/disability"
	"github.com/emurmotol/nmsrs/models/educationlevel"
	"github.com/emurmotol/nmsrs/models/eligibility"
	"github.com/emurmotol/nmsrs/models/employmentstatus"
	"github.com/emurmotol/nmsrs/models/industry"
	"github.com/emurmotol/nmsrs/models/language"
	"github.com/emurmotol/nmsrs/models/license"
	"github.com/emurmotol/nmsrs/models/otherskill"
	"github.com/emurmotol/nmsrs/models/position"
	"github.com/emurmotol/nmsrs/models/province"
	"github.com/emurmotol/nmsrs/models/region"
	"github.com/emurmotol/nmsrs/models/religion"
	"github.com/emurmotol/nmsrs/models/school"
	"github.com/emurmotol/nmsrs/models/sex"
	"github.com/emurmotol/nmsrs/models/skill"
	"github.com/emurmotol/nmsrs/models/unemployedstatus"
	"github.com/emurmotol/nmsrs/models/user"
	"github.com/emurmotol/nmsrs/routes"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewStatic(http.Dir("static")))
	n.UseHandler(routes.Register())
	user.SetDefaultUser()
	seed()
	host := fmt.Sprintf("%s:%d", env.SvrHost, env.SvrPort)
	log.Printf("Server running at %s\n", host)

	if err := http.ListenAndServe(host, n); err != nil {
		panic(err)
	}
}

func seed() {
	barangay.Seeder()
	certificate.Seeder()
	citymunicipality.Seeder()
	civilstatus.Seeder()
	country.Seeder()
	course.Seeder()
	disability.Seeder()
	educationlevel.Seeder()
	eligibility.Seeder()
	employmentstatus.Seeder()
	industry.Seeder()
	language.Seeder()
	license.Seeder()
	otherskill.Seeder()
	position.Seeder()
	province.Seeder()
	region.Seeder()
	// registrant.Seeder()
	religion.Seeder()
	school.Seeder()
	sex.Seeder()
	skill.Seeder()
	unemployedstatus.Seeder()
	user.Seeder()
}
