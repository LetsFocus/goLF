package validations

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
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
	pattern := `^\+?[1-9]\d{6,13}$`
	return regexp.MustCompile(pattern).MatchString(phoneNumber)
}

// IsValidUUID checks if a string is a valid UUID (version 4).
func IsValidUUID(u string) bool {
	if !strings.Contains(u, "-") {
		return false
	}
	_, err := uuid.Parse(u)
	return err == nil
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
func RemoveSliceDuplicates(slice []interface{}) []interface{} {
	encountered := map[interface{}]bool{}
	result := []interface{}{}

	for v := range slice {
		if !encountered[slice[v]] {
			encountered[slice[v]] = true
			result = append(result, slice[v])
		}
	}
	return result
}

// RemoveSliceElement removes a particular element in a slice based on index
func RemoveSliceElement(slice []interface{}, index int) []interface{} {
	if index >= 0 && index < len(slice) {
		if index == len(slice)-1 {
			return slice[:index]
		}
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
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
