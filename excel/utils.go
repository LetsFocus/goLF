package excel

import (
	"encoding/json"
	"os"
)

func validateFilePath(filePath string) (*os.File, *os.File, error) {
	var writeFile, readFile *os.File
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		writeFile, err = os.Create(filePath)
		if err != nil {
			return nil, nil, err
		}
		readFile, err = os.Open(filePath)
		if err != nil {
			return nil, nil, err
		}
	}
	return writeFile, readFile, nil
}

func writeToConfig(autoId int) error {
	config := map[string]int{"AUTO_ID": autoId}
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(CONFIG_PATH, jsonData, 0644)
}

func buildHeaderMap(headers []string, primaryKey string) map[string]int {
	headersMap := make(map[string]int)
	startIndex := 0
	if primaryKey == "auto_id" {
		headersMap[primaryKey] = startIndex
		startIndex = 1
	}
	for index, field := range headers {
		headersMap[field] = index + startIndex
	}
	return headersMap
}

func eraseFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}