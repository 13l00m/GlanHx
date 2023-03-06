package utils

import "regexp"

func VerifyHost(host string) bool {
	hostRegex := regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)
	if hostRegex.MatchString(host) {
		return true
	}
	return false
}
