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
	} else {
		return false
	}

	dayInt, err := strconv.Atoi(day)
	if err != nil {
		return false
	}
	if dayInt > 0 && dayInt <= 31 {
		output = true
	} else {
		return false
	}

	if monthInt == 2 && dayInt > 29 {
		return false
	}

	return output
}
