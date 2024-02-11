package validations

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"ValidEmail", "test@example.com", true},
		{"InvalidEmail_NoAtSymbol", "invalid_email", false},
		{"InvalidEmail_NoDomain", "invalid_email@", false},
		{"InvalidEmail_NoTopLevelDomain", "invalid_email@example", false},
		{"InvalidEmail_EmptyString", "", false},
		{"InvalidEmail_OnlyAtSymbol", "@", false},
		{"InvalidEmail_OnlyDomain", "example.com", false},
		{"InvalidEmail_OnlyTopLevelDomain", ".com", false},
		{"InvalidEmail_MissingLocalPart", "@example.com", false},
		{"InvalidEmail_MissingDomainPart", "test@", false},
		{"InvalidEmail_WhiteSpace", "test @ example.com", false},
		{"InvalidEmail_SpecialCharacters", "test!@example.com", false},
		{"InvalidEmail_MultipleAtSymbols", "test@@example.com", false},
		{"InvalidEmail_InvalidCharacters", "test@example!com", false},
		{"ValidEmail_UpperCaseLetters", "TEST@EXAMPLE.COM", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidEmail(tt.email)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsValidPhoneNumber(t *testing.T) {
	tests := []struct {
		name        string
		phoneNumber string
		expected    bool
	}{
		{"ValidPhoneNumber", "+123456789012", true},
		{"ValidPhoneNumber_NoCountryCode", "123456789012", true},
		{"ValidPhoneNumber_NoPlusSign", "123456789012", true},
		{"InvalidPhoneNumber_TooShort", "+1", false},
		{"InvalidPhoneNumber_TooLong", "+12345678901234567890", false},
		{"InvalidPhoneNumber_ContainsLetters", "+1234567a89012", false},
		{"InvalidPhoneNumber_SpecialCharacters", "+1234@6789012", false},
		{"InvalidPhoneNumber_EmptyString", "", false},
		{"InvalidPhoneNumber_WhiteSpace", " 123456789012 ", false},
		{"InvalidPhoneNumber_LeadingZeros", "+00123456789012", false},
		{"InvalidPhoneNumber_OnlyPlusSign", "+", false},
		{"InvalidPhoneNumber_InvalidCountryCode", "+0123456789012", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidPhoneNumber(tt.phoneNumber)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsValidUUID(t *testing.T) {
	tests := []struct {
		name     string
		uuidStr  string
		expected bool
	}{
		{"ValidUUIDv4", "b5fcdf08-9ff9-415f-85f2-ffdb3e03cbb7", true},
		{"ValidUUIDv1", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", true},
		{"InvalidUUID_TooShort", "b5fcdf08-9ff9-415f-85f2", false},
		{"InvalidUUID_TooLong", "b5fcdf08-9ff9-415f-85f2-ffdb3e03cbb7-extra", false},
		{"InvalidUUID_InvalidCharacters", "b5fcdf08-9ff9-415f-85f2-ffdb3e03cbbg", false},
		{"InvalidUUID_InvalidFormat", "b5fcdf08_9ff9_415f_85f2_ffdb3e03cbb7", false},
		{"InvalidUUID_EmptyString", "", false},
		{"InvalidUUID_NonHexCharacters", "b5fcdf08-9ff9-415f-85f2-ffdb3e03cbb-", false},
		{"InvalidUUID_NotDashSeparated", "b5fcdf089ff9415f85f2ffdb3e03cbb7", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidUUID(tt.uuidStr)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsValidTime(t *testing.T) {
	tests := []struct {
		name     string
		timeStr  string
		expected bool
	}{
		{"ValidTime", "12:34:56", true},
		{"ValidTime_ZeroHour", "00:00:00", true},
		{"ValidTime_MaxHour", "23:59:59", true},
		{"InvalidTime_TooShort", "1:2:3", false},
		{"InvalidTime_TooLong", "123:456:789", false},
		{"InvalidTime_InvalidHour", "24:00:00", false},
		{"InvalidTime_InvalidMinute", "12:60:00", false},
		{"InvalidTime_InvalidSecond", "12:34:60", false},
		{"InvalidTime_InvalidFormat", "12-34-56", false},
		{"InvalidTime_EmptyString", "", false},
		{"InvalidTime_WhiteSpace", "   ", false},
		{"InvalidTime_NonNumeric", "12:34:5a", false},
		{"InvalidTime_NoColon", "123456", false},
		{"InvalidTime_ExtraColon", "12:34:56:78", false},
		{"InvalidTime_LeadingZeros", "001:002:003", false},
		{"InvalidTime_NegativeHour", "-12:34:56", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidTime(tt.timeStr)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsValidDate(t *testing.T) {
	tests := []struct {
		name     string
		dateStr  string
		expected bool
	}{
		{"ValidDate", "2024-02-04", true},
		{"ValidDate_ZeroMonth", "2024-00-01", false},
		{"ValidDate_MaxMonth", "2024-12-31", true},
		{"ValidDate_LeapYear", "2024-02-29", true},
		{"InvalidDate_TooShort", "2024-2-4", false},
		{"InvalidDate_TooLong", "2024-02-040", false},
		{"InvalidDate_InvalidYear", "20a4-02-04", false},
		{"InvalidDate_InvalidMonth", "2024-13-04", false},
		{"InvalidDate_InvalidDay", "2024-02-32", false},
		{"InvalidDate_InvalidFormat", "2024/02/04", false},
		{"InvalidDate_EmptyString", "", false},
		{"InvalidDate_WhiteSpace", "   ", false},
		{"InvalidDate_NonNumeric", "2024-02-0a", false},
		{"InvalidDate_ExtraHyphen", "2024--02-04", false},
		{"InvalidDate_LeadingZeros", "02024-002-004", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidDate(tt.dateStr)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRemoveSliceDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{"RemoveDuplicates_Integers", []interface{}{1, 2, 3, 3, 4, 5, 5, 6}, []interface{}{1, 2, 3, 4, 5, 6}},
		{"RemoveDuplicates_Strings", []interface{}{"apple", "banana", "apple", "orange", "banana"}, []interface{}{"apple", "banana", "orange"}},
		{"RemoveDuplicates_Booleans", []interface{}{true, true, false, false, true}, []interface{}{true, false}},
		{"RemoveDuplicates_MixedTypes", []interface{}{1, "apple", true, 2, "apple", false}, []interface{}{1, "apple", true, 2, false}},
		{"RemoveDuplicates_EmptySlice", []interface{}{}, []interface{}{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveSliceDuplicates(tt.input)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func TestRemoveSliceElement(t *testing.T) {
	tests := []struct {
		name         string
		inputSlice   []interface{}
		indexToRemove int
		expected     []interface{}
	}{
		{"RemoveFromIntSlice", []interface{}{1, 2, 3, 4, 5}, 2, []interface{}{1, 2, 4, 5}},
		{"RemoveFromStringSlice", []interface{}{"apple", "banana", "orange", "grape", "kiwi"}, 0, []interface{}{"banana", "orange", "grape", "kiwi"}},
		{"RemoveFromBoolSlice", []interface{}{true, false, true, true, false}, 3, []interface{}{true, false, true, false}},
		{"RemoveOutOfBounds", []interface{}{1, 2, 3, 4, 5}, 5, []interface{}{1, 2, 3, 4, 5}},
		{"RemoveFromEmptySlice", []interface{}{}, 0, []interface{}{}},
		{"RemoveNegativeIndex", []interface{}{1, 2, 3, 4, 5}, -1, []interface{}{1, 2, 3, 4, 5}},
		{"RemoveFromSingleElementSlice", []interface{}{1}, 0, []interface{}{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveSliceElement(tt.inputSlice, tt.indexToRemove)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStrToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		err      error
	}{
		{"ValidIntegerString", "123", 123, nil},
		{"NegativeIntegerString", "-456", -456, nil},
		{"Zero", "0", 0, nil},
		{"EmptyString", "", 0, &strconv.NumError{Func: "Atoi", Num: "", Err: strconv.ErrSyntax}},
		{"NonIntegerString", "abc", 0, &strconv.NumError{Func: "Atoi", Num: "abc", Err: strconv.ErrSyntax}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StrToInt(tt.input)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestIntToStr(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"PositiveInteger", 123, "123"},
		{"NegativeInteger", -456, "-456"},
		{"Zero", 0, "0"},
		{"LargeInteger", 987654321, "987654321"},
		{"SmallInteger", -987654321, "-987654321"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntToStr(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRandInt(t *testing.T) {
	rand.Seed(42)
	a := 100
	numIterations := 1000

	for i := 0; i < numIterations; i++ {
		result := RandInt(a)
		if result < 0 || result >= a {
			t.Errorf("Generated number %d is not within the range [0, %d)", result, a)
		}
	}
}
