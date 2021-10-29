package dob

import "time"

func GetRangeOfDob(dob time.Time) string {
	timeNow := time.Now()

	if dob.After(timeNow) {
		return "you are not burn yet :)"
	}

	currentYear := timeNow.Year()
	year := dob.Year()
	diffAge := currentYear - year

	if diffAge >= 0 && diffAge <= 14 {
		return "kid"
	} else if diffAge >= 15 && diffAge <= 24 {
		return "young"
	} else if diffAge > 24 && diffAge <= 64 {
		return "adult"
	} else {
		return "old"
	}

}
