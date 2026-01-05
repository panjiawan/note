package dao

import "gorm.io/gorm"

func DbIsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == gorm.ErrRecordNotFound.Error()
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}
