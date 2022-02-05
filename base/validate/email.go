package validate

// based on https://github.com/goware/emailx package

import (
	"net"
	"regexp"
	"strings"
)

var (
	userRegexp = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
	hostRegexp = regexp.MustCompile("^[^\\s]+\\.[^\\s]+$")
	// Dots are not allowed in the beginning, end or in occurances of more than 1 in the email address
	userDotRegexp = regexp.MustCompile("(^[.]{1})|([.]{1}$)|([.]{2,})")
)

// EmailHost checks format of a given email and resolves its host name.
func EmailHost(email string) bool {
	if valid := EmailFormat(email); !valid {
		return false
	}

	at := strings.LastIndex(email, "@")
	host := email[at+1:]

	if _, err := net.LookupMX(host); err != nil {
		if _, err := net.LookupIP(host); err != nil {
			return false
		}
	}

	return true
}

// EmailFormat checks format of a given email.
func EmailFormat(email string) bool {
	if len(email) < 6 || len(email) > 254 {
		return false
	}

	at := strings.LastIndex(email, "@")
	if at <= 0 || at > len(email)-3 {
		return false
	}

	user := email[:at]
	host := email[at+1:]

	if len(user) > 64 {
		return false
	}
	if userDotRegexp.MatchString(user) || !userRegexp.MatchString(user) || !hostRegexp.MatchString(host) {
		return false
	}

	return true
}
