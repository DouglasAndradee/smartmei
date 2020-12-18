package helper

import (
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

// ValidEmail - Validates an email
func ValidEmail(email string) bool {
	result := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return result.MatchString(email)
}

// IsDup - Avoid duplicates emails
func IsDup(err error) bool {
	wce, ok := err.(mongo.WriteConcernError)
	if !ok {
		return false
	}
	return wce.Code == 11000 || wce.Code == 11001 || wce.Code == 12582 || wce.Code == 16460 && strings.Contains(wce.Message, " E11000 ")
}
