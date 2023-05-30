package utils

import "strings"

func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

func AnyBlank(strings ...string) bool {
	for _, str := range strings {
		if IsBlank(str) {
			return true
		}
	}
	return false
}
