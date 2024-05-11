package excel

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"os"
)

type Excel struct {
	filePath   string
	configPath string
}

func NewExcel(filePath string, configPath string) (*Excel, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
	}

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		_, err := os.Create(configPath)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.Create(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	configData := map[string]interface{}{
		"primaryKey":  "",
		"headers":     nil,
		"primaryKeys": nil,
		"autoId":      1,
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(configData)
	if err != nil {
		return nil, err
	}

	return &Excel{filePath: filePath, configPath: configPath}, nil
}

func (e *Excel) GetAllRows() ([][]string, error) {
	file, err := os.Open(e.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (e *Excel) AddHeader(primaryKey string, headers []string) error {
	configFile, err := os.Open(e.configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	var configData map[string]interface{}
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&configData)
	if err != nil {
		return err
	}

	excelFile, err := os.Open(e.filePath)
	if err != nil {
		return err
	}
	defer excelFile.Close()

	if value, ok := configData["headers"]; ok {
		_, ok := value.(map[string]int)
		if !ok {
			return errors.New("header already present")
		}
	} else {
		return errors.New("headers not found in config")
	}

	configData["primaryKey"] = primaryKey
	headersMap := make(map[string]int)
	headersMap[primaryKey] = 0
	for index, field := range headers {
		headersMap[field] = index + 1
	}
	configData["headers"] = headersMap

	rows, err := e.GetAllRows()
	if err != nil {
		return err
	}

	configFileWrite, err := os.Create(e.configPath)
	if err != nil {
		return err
	}
	defer configFileWrite.Close()

	excelFileWrite, err := os.Create(e.configPath)
	if err != nil {
		return err
	}
	defer excelFileWrite.Close()

	headers = append([]string{primaryKey}, headers...)
	rows = append([][]string{headers}, rows...)
	writer := csv.NewWriter(excelFileWrite)
	err = writer.WriteAll(rows)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(configFileWrite)
	err = encoder.Encode(configData)
	if err != nil {
		return err
	}
	return nil
}

// func (e *Excel) AddRow(row []string) error {
// 	if e.primaryKey == "auto" {
// 		row = append([]string{strconv.Itoa(e.autoId)}, row...)
// 		e.autoId = e.autoId + 1
// 	} else {

// 	}
// 	file, err := os.OpenFile(e.filePath, os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()
// 	err = writer.Write(row)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (e *Excel) GetRow(id string) ([]string, error) {
// 	file, err := os.Open(e.filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	rowNum, err := e.GetRowFromId(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if rowNum == -1 {
// 		return nil, errors.New("an error occured")
// 	}

// 	reader := csv.NewReader(file)
// 	rows, err := reader.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if rowNum >= len(rows) {
// 		return nil, errors.New("row number out of range")
// 	}

// 	return rows[rowNum], nil
// }

// func (e *Excel) GetLastId() (int, error) {
// 	file, err := os.Open(e.filePath)
// 	if err != nil {
// 		return -1, err
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	rows, err := reader.ReadAll()
// 	if err != nil {
// 		return -1, err
// 	}

// 	if len(rows) <= 0 || len(rows[len(rows)-1]) <= 0 {
// 		return 1, nil
// 	}
// 	if rows[len(rows)-1][0] == "id" {
// 		return 1, nil
// 	}
// 	id, err := strconv.Atoi(rows[len(rows)-1][0])
// 	if err != nil {
// 		return -1, err
// 	}
// 	return id + 1, nil
// }

// func (e *Excel) GetRowFromId(id string) (int, error) {
// 	file, err := os.Open(e.filePath)
// 	if err != nil {
// 		return -1, err
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	rows, err := reader.ReadAll()
// 	if err != nil {
// 		return -1, err
// 	}

// 	if len(rows) <= 0 {
// 		return -1, errors.New("ID not available")
// 	}

// 	for index, row := range rows {
// 		if len(row) > 0 && row[0] == id {
// 			return index, nil
// 		}
// 	}
// 	return -1, nil
// }

// func (e *Excel) ReplaceRow(id string, row []string) error {
// 	file, err := os.Open(e.filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	row = append([]string{id}, row...)
// 	rowNum, err := e.GetRowFromId(id)
// 	if err != nil {
// 		return err
// 	}
// 	if rowNum == -1 {
// 		return errors.New("an error occured")
// 	}

// 	reader := csv.NewReader(file)
// 	allRows, err := reader.ReadAll()
// 	if err != nil {
// 		return err
// 	}

// 	if rowNum < 0 || rowNum >= len(allRows) {
// 		return errors.New("row number out of range")
// 	}

// 	allRows[rowNum] = row

// 	file, err = os.Create(e.filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	err = writer.WriteAll(allRows)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (e *Excel) ReplaceHeader(header []string) error {
// 	file, err := os.Open(e.filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	header = append([]string{"id"}, header...)
// 	reader := csv.NewReader(file)
// 	allRows, err := reader.ReadAll()
// 	if err != nil {
// 		return err
// 	}

// 	validateHeader, err := e.IsHeaderExist()
// 	if err != nil {
// 		return err
// 	}
// 	if !validateHeader {
// 		return errors.New("header not available")
// 	}

// 	allRows[0] = header

// 	file, err = os.Create(e.filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	err = writer.WriteAll(allRows)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
