package database

import "strings"

func cleanAddresses(addressesString string) []string {
	addressesList := strings.Split(addressesString, ",")

	var addressesSlice []string

	for _, address := range addressesList {
		trimmedAddress := strings.TrimSpace(address)
		addressesSlice = append(addressesSlice, trimmedAddress)
	}

	return addressesSlice
}
