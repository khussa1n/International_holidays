package telegram

import (
	"github.com/joshtronic/go-holidayapi"
	"github.com/spf13/viper"
)

func getHolidays(country, month, day string) (holidayapi.Respone, error) {
	hapi := holidayapi.NewV1(viper.GetString("db.host"))

	return hapi.Holidays(map[string]string{
		"country":  country,
		"year":     "2022",
		"month":    month,
		"day":      day,
		"language": "RU",
	})
}
