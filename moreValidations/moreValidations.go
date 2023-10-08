package moreValidations

import (
	"math/rand"
	"strconv"
)

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

func StrTOInt(s string) (int, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return val, nil
}
func IntTOStr(value int) string {
	val := strconv.Itoa(value)
	return val

}

func RandInt(a int) int {
	return rand.Intn(a)
}
