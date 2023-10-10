package validations

import (
	"math/rand"
	"regexp"
	"strconv"
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

// RemoveSliceDuplicates revomes the duplicate elements in a slice.
func RemoveSliceDuplicates[T comparable](slice []T) []T {
	encountered := map[T]bool{}
	result := []T{}

	for v := range slice {
		if encountered[slice[v]] == false {
			encountered[slice[v]] = true
			result = append(result, slice[v])
		}
	}
	return result
}

// RemoveSliceElement removes a particular element in a slice based on index
func RemoveSliceElement[T comparable](slice []T, index int) []T {
	result := []T{}
	if index < len(slice)-1 {
		result = append(slice[:index], slice[index+1:]...)
	}
	if index == len(slice)-1 {
		result = append(slice[:index])
	}
	return result
}

// StrToInt converts string to integer return in a val, err format
func StrToInt(s string) (int, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// IntToStr converts integer to string
func IntToStr(value int) string {
	val := strconv.Itoa(value)
	return val

}

// RendInt return a random integer from the range[0,n]
func RandInt(a int) int {
	return rand.Intn(a)
}
