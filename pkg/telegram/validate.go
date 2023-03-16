package telegram

import "strconv"

func ValidateDate(month, day string) bool {
	output := false

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		return false
	}
	if monthInt > 0 && monthInt <= 12 {
		output = true
	}

	dayInt, err := strconv.Atoi(day)
	if err != nil {
		return false
	}
	if dayInt > 0 && dayInt <= 31 {
		output = true
	}

	if monthInt == 02 && dayInt >= 30 {
		return false
	}

	return output
}
