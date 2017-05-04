package user

import (
	"log"

	"github.com/icrowley/fake"
)

func Seeder() {
	users, err := All()

	if err != nil {
		panic(err)
	}

	if len(users) < 5 {
		for i := 0; i < 5; i++ {
			var usr User
			usr.Name = fake.FullName()
			usr.Email = fake.EmailAddress()
			usr.Password = "secret"
			usr.IsAdmin = false
			usr.PhotoIsSet = false
			_, err := usr.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("User seeded")
	}
}
