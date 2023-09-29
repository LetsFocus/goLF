package validations

import (
	"regexp"
	"time"

	"github.com/google/uuid"
)

// IsValidEmail checks if an email address is valid.
func IsValidEmail(email string) bool {
	// Regular expression for basic email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(pattern).MatchString(email)
}

// IsValidPhoneNumber checks if a string is a valid phone number.
func IsValidPhoneNumber(phoneNumber string) bool {
	pattern := `^\+?[1-9]\d{1,14}$`
	return regexp.MustCompile(pattern).MatchString(phoneNumber)
}

// IsValidUUID checks if a string is a valid UUID (version 4).
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	if err != nil {
		return false
	}

	return true
}

// IsValidTime checks if a string is a valid time in HH:MM:SS format.
func IsValidTime(timeStr string) bool {
	pattern := `^([0-1][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]$`
	return regexp.MustCompile(pattern).MatchString(timeStr)
}

// IsValidDate checks if a string is a valid date in yyyy-mm-dd format.
func IsValidDate(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
