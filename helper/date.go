package helper

import "time"

func DateForHumans(sec int64) string {
	return time.Unix(sec, 0).Format("January 2, 2006")
}

func ParsleyDateNow() string {
	return time.Now().Format("2006-01-02")
}

func YearMonth(str string) time.Time {
	t, err := time.Parse("2006-01", str)

	if err != nil {
		panic(err)
	}
	return t
}

func ShortDate(str string) time.Time {
	t, err := time.Parse("2006-01-02", str)

	if err != nil {
		panic(err)
	}
	return t
}
