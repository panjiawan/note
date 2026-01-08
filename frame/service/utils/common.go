package utils

import "regexp"

type LabelEtl struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}

func ValidMobile(mobile string) bool {
	regRuler := "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(mobile)
}
