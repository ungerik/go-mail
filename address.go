package email

import (
	"errors"
	"regexp"
	"strings"
)

var emailRegexp *regexp.Regexp

func NormalizeAddress(address string) string {
	return strings.ToLower(strings.TrimSpace(address))
}

func ValidateAddress(address string) (normalizedAddress string, err error) {
	if emailRegexp == nil {
		emailRegexp = regexp.MustCompile("[a-z0-9.\\-_%]+@[a-z0-9.\\-]+\\.[a-z][a-z]+")
	}
	normalizedAddress = NormalizeAddress(address)
	valid := emailRegexp.Match([]byte(normalizedAddress))
	if !valid {
		return "", errors.New("Invalid email address: " + address)
	}
	return normalizedAddress, nil
}
