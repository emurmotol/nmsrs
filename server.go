package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/zneyrl/nmsrs/env"
	"github.com/zneyrl/nmsrs/models/certificate"
	"github.com/zneyrl/nmsrs/models/civilstatus"
	"github.com/zneyrl/nmsrs/models/country"
	"github.com/zneyrl/nmsrs/models/course"
	"github.com/zneyrl/nmsrs/models/disability"
	"github.com/zneyrl/nmsrs/models/educationlevel"
	"github.com/zneyrl/nmsrs/models/eligibility"
	"github.com/zneyrl/nmsrs/models/employmentstatus"
	"github.com/zneyrl/nmsrs/models/industry"
	"github.com/zneyrl/nmsrs/models/language"
	"github.com/zneyrl/nmsrs/models/license"
	"github.com/zneyrl/nmsrs/models/otherskill"
	"github.com/zneyrl/nmsrs/models/position"
	"github.com/zneyrl/nmsrs/models/religion"
	"github.com/zneyrl/nmsrs/models/school"
	"github.com/zneyrl/nmsrs/models/sex"
	"github.com/zneyrl/nmsrs/models/skill"
	"github.com/zneyrl/nmsrs/models/unemployedstatus"
	"github.com/zneyrl/nmsrs/models/user"
	"github.com/zneyrl/nmsrs/routes"
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
	fmt.Printf("Server running at %s\n", host)

	if err := http.ListenAndServe(host, n); err != nil {
		panic(err)
	}
}

func seed() {
	certificate.Seeder()
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
	// registrant.Seeder()
	religion.Seeder()
	school.Seeder()
	sex.Seeder()
	skill.Seeder()
	unemployedstatus.Seeder()
	user.Seeder()
}
