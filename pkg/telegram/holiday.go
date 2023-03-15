package telegram

import (
	"github.com/joshtronic/go-holidayapi"
)

func getHolidays(country, month, day string) (holidayapi.Respone, error) {
	hapi := holidayapi.NewV1("baa0a205-4931-4e44-95c2-22ddde2ea585")

	return hapi.Holidays(map[string]string{
		"country":  country,
		"year":     "2022",
		"month":    month,
		"day":      day,
		"language": "RU",
	})
}
