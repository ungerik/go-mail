package email

import (
	"regexp"
	"strings"
)

var emailRegexp *regexp.Regexp = regexp.MustCompile(`[a-z0-9\-+~_%]+[a-z0-9\-+~_%.]*@([a-z]+[a-z0-9\\-]*\.)+[a-z][a-z]+`)

// NormalizeAddress assumes that all email adresses can be
// converted to lower case, which is empirically true
// but not specification conform.
func NormalizeAddress(address string) string {
	return strings.ToLower(strings.TrimSpace(address))
}

// ValidateAddress uses a simplified regular expression for checking
// email adresses. Please report real world adresses that are not working as bug.
func ValidateAddress(address string) (normalizedAddress string, err error) {
	normalizedAddress = NormalizeAddress(address)
	valid := emailRegexp.Match([]byte(normalizedAddress))
	if !valid {
		return "", ErrInvalidEmailAddress(address)
	}
	return normalizedAddress, nil
}

type ErrInvalidEmailAddress string

func (self ErrInvalidEmailAddress) Error() string {
	return "Invalid email address: " + string(self)
}
