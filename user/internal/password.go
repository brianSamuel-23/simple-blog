package internal

import (
	"strings"
	"unicode"
)

const minPasswordClasses = 3

// commonPasswords blocks a small set of the most frequently breached
// passwords; a real deployment would check against a much larger corpus
// (e.g. HaveIBeenPwned's Pwned Passwords range API), but even this narrow
// list stops the most trivial automated credential-stuffing attempts.
var commonPasswords = map[string]struct{}{
	"password":    {},
	"12345678":    {},
	"123456789":   {},
	"qwerty123":   {},
	"letmein123":  {},
	"iloveyou123": {},
	"admin1234":   {},
	"welcome123":  {},
	"password123": {},
	"changeme123": {},
}

// isPasswordStrong rejects passwords that are trivially guessable: ones
// pulled from a common-password list, ones built from a single character
// class (e.g. all lowercase letters), and ones that just echo the account's
// own name or email back, which defeats the point of a separate secret.
func isPasswordStrong(password, name, email string) bool {
	lower := strings.ToLower(password)

	if _, ok := commonPasswords[lower]; ok {
		return false
	}

	if name = strings.ToLower(strings.TrimSpace(name)); name != "" && strings.Contains(lower, name) {
		return false
	}

	if localPart, _, ok := strings.Cut(email, "@"); ok {
		if localPart = strings.ToLower(localPart); localPart != "" && strings.Contains(lower, localPart) {
			return false
		}
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r), unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	classes := 0
	for _, present := range [...]bool{hasUpper, hasLower, hasDigit, hasSpecial} {
		if present {
			classes++
		}
	}

	return classes >= minPasswordClasses
}
