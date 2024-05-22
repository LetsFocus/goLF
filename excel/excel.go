package excel

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"os"
	"sort"
	"strconv"
)

func InitExcel(filePath string, primaryKey string, isHeaderExist bool) (*Excel, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
	}

	configPath := "config.json"
	if primaryKey == "auto_id" {
		_, err = os.Stat(configPath)
		if os.IsNotExist(err) {
			_, err := os.Create(configPath)
			if err != nil {
				return nil, err
			}
			config := make(map[string]int, 0)
			config["AUTO_ID"] = 0
			jsonData, err := json.MarshalIndent(config, "", "  ")
			if err != nil {
				return nil, err
			}

			err = os.WriteFile(configPath, jsonData, 0644)
			if err != nil {
				return nil, err
			}
		}
	}

	return &Excel{filePath: filePath, configPath: configPath, primaryKey: primaryKey,
		isHeaderExist: isHeaderExist, headersMap: nil, header: nil}, nil
}

func (e *Excel) checkAndIncrementAutoID() (int, error) {
	file, err := os.Open(e.configPath)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return -1, err
	}

	var config map[string]int
	err = json.Unmarshal(data, &config)
	if err != nil {
		return -1, err
	}

	autoID, ok := config["AUTO_ID"]
	if !ok {
		autoID = 0
	}

	config["AUTO_ID"] = autoID+1
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return -1, err
	}

	err = os.WriteFile(e.configPath, jsonData, 0644)
	if err != nil {
		return -1, err
	}

	return autoID, nil
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

func (e *Excel) checkPrimaryKey(primaryKey string) bool {
	if e.primaryKey == "auto_id" {
		return false
	}
	primaryKeyPosition := e.headersMap[e.primaryKey]
	rows, _ := e.GetAllRows()
	for _, row := range rows {
		if row[primaryKeyPosition] == primaryKey {
			return true
		}
	}
	return false
}

func (e *Excel) buildHeaderMap(headers []string) map[string]int {
	headersMap := make(map[string]int)
	if e.primaryKey == "auto_id" {
		headersMap[e.primaryKey] = 0
		for index, field := range headers {
			headersMap[field] = index + 1
		}
	} else {
		for index, field := range headers {
			headersMap[field] = index
		}
	}
	return headersMap
}

func (e *Excel) AddHeader(headers []string) error {
	excelFile, err := os.Open(e.filePath)
	if err != nil {
		return err
	}
	defer excelFile.Close()

	if e.isHeaderExist {
		return errors.New("header already present")
	}

	e.headersMap = e.buildHeaderMap(headers)

	rows, err := e.GetAllRows()
	if err != nil {
		return err
	}

	file, err := os.Create(e.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if e.primaryKey == "auto_id" {
		headers = append([]string{e.primaryKey}, headers...)
	}
	rows = append([][]string{headers}, rows...)
	writer := csv.NewWriter(file)
	err = writer.WriteAll(rows)
	if err != nil {
		return err
	}
	e.isHeaderExist = true
	return nil
}

func (e *Excel) ReplaceHeaderName(headers map[string]string) error {
	excelFile, err := os.Open(e.filePath)
	if err != nil {
		return err
	}
	defer excelFile.Close()

	rows, err := e.GetAllRows()
	if err != nil {
		return err
	}

	if !e.isHeaderExist {
		return errors.New("header not present")
	}

	if e.primaryKey == "auto_id" {
		e.headersMap = e.buildHeaderMap(rows[0][1:])
	} else {
		e.headersMap = e.buildHeaderMap(rows[0])
	}

	newHeaders := make([]string, len(e.headersMap))
	if e.primaryKey == "auto_id" {
		newHeaders[0] = "auto_id"
	}
	for key, value := range headers {
		if key == e.primaryKey {
			e.primaryKey = value
		}
		val, ok := e.headersMap[key]
		if !ok {
			return errors.New("header doesn't exist")
		}
		delete(e.headersMap, key)
		e.headersMap[value] = val
		if e.primaryKey == "auto_id" {
			newHeaders[val-1] = value
		} else {
			newHeaders[val] = value
		}
	}

	file, err := os.Create(e.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	rows[0] = newHeaders
	writer := csv.NewWriter(file)
	err = writer.WriteAll(rows)
	if err != nil {
		return err
	}
	e.isHeaderExist = true
	return nil
}

func (e *Excel) AddRow(row []string) error {
	rows, err := e.GetAllRows()
	if err != nil {
		return err
	}

	if e.primaryKey == "auto_id" {
		currentId, err := e.checkAndIncrementAutoID()
		if err != nil {
			return err
		}
		row = append([]string{strconv.Itoa(currentId)}, row...)
	} else {
		isPrimaryKeyDuplicated := e.checkPrimaryKey(row[e.headersMap[e.primaryKey]])
		if isPrimaryKeyDuplicated {
			return errors.New("primary key duplicated")
		}
	}
	rows = append(rows, row)

	file, err := os.Create(e.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(rows)
	if err != nil {
		return err
	}
	return nil
}

func (h *Header) AddRow(row []string) error {
	rows, err := h.excel.GetAllRows()
	if err != nil {
		return err
	}

	if h.excel.primaryKey == "auto_id" {
		rowToAdd := make([]string, len(row))
		currentId, err := h.excel.checkAndIncrementAutoID()
		if err != nil {
			return err
		}
		for i := 0; i < len(h.headers); i++ {
			rowToAdd[h.excel.headersMap[h.headers[i]]-1] = row[i]
		}
		row = append([]string{strconv.Itoa(currentId)}, rowToAdd...)
	} else {
		rowToAdd := make([]string, len(row))
		for i := 0; i < len(h.headers); i++ {
			rowToAdd[h.excel.headersMap[h.headers[i]]] = row[i]
		}
		row = rowToAdd
		isPrimaryKeyDuplicated := h.excel.checkPrimaryKey(row[h.excel.headersMap[h.excel.primaryKey]])
		if isPrimaryKeyDuplicated {
			return errors.New("primary key duplicated")
		}
	}
	rows = append(rows, row)

	file, err := os.Create(h.excel.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(rows)
	if err != nil {
		return err
	}
	return nil
}

func (e *Excel) Columns(headers []string) *Header {
	newHeader := &Header{
		headers: headers,
		excel:   e,
	}
	e.header = newHeader
	return e.header
}

func (e *Excel) getRowNumFromId(primaryKey string) (int, error) {
	file, err := os.Open(e.filePath)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return -1, err
	}

	if len(rows) <= 0 {
		return -1, errors.New("ID not available")
	}

	primaryKeyPosition := e.headersMap[e.primaryKey]
	for index, row := range rows {
		if primaryKeyPosition < len(row) && row[primaryKeyPosition] == primaryKey {
			return index, nil
		}
	}
	return -1, nil
}

func (e *Excel) GetRow(primaryKey string) ([]string, error) {
	file, err := os.Open(e.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rowNumber, err := e.getRowNumFromId(primaryKey)
	if err != nil {
		return nil, err
	}
	if rowNumber == -1 {
		return nil, errors.New("an error occured")
	}

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if rowNumber >= len(rows) {
		return nil, errors.New("row number out of range")
	}

	return rows[rowNumber], nil
}

func (e *Excel) ReplaceRow(primaryKey string, row []string) error {
	file, err := os.Open(e.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	rowNumber, err := e.getRowNumFromId(primaryKey)
	if err != nil {
		return err
	}
	if rowNumber == -1 {
		return errors.New("an error occured")
	}

	reader := csv.NewReader(file)
	allRows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if rowNumber < 0 || rowNumber >= len(allRows) {
		return errors.New("row number out of range")
	}

	if e.primaryKey == "auto_id" {
		row = append([]string{primaryKey}, row...)
	}

	allRows[rowNumber] = row

	file, err = os.Create(e.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(allRows)
	if err != nil {
		return err
	}

	return nil
}

func (h *Header) ReplaceRow(primaryKey string, row []string) error {
	file, err := os.Open(h.excel.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	rowNumber, err := h.excel.getRowNumFromId(primaryKey)
	if err != nil {
		return err
	}
	if rowNumber == -1 {
		return errors.New("an error occured")
	}

	reader := csv.NewReader(file)
	allRows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if rowNumber < 0 || rowNumber >= len(allRows) {
		return errors.New("row number out of range")
	}

	if h.excel.primaryKey == "auto_id" {
		rowToAdd := make([]string, len(row))
		for i := 0; i < len(h.headers); i++ {
			rowToAdd[h.excel.headersMap[h.headers[i]]-1] = row[i]
		}
		row = append([]string{primaryKey}, rowToAdd...)
	} else {
		rowToAdd := make([]string, len(row))
		for i := 0; i < len(h.headers); i++ {
			rowToAdd[h.excel.headersMap[h.headers[i]]] = row[i]
		}
		row = rowToAdd
	}
	allRows[rowNumber] = row

	file, err = os.Create(h.excel.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(allRows)
	if err != nil {
		return err
	}
	return nil
}

func (e *Excel) DeleteRows(primaryKeys []string) error {
	file, err := os.Open(e.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	allRows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	rowNumbers := make([]int, 0)
	for _, primaryKey := range primaryKeys {
		rowNumber, err := e.getRowNumFromId(primaryKey)
		if err != nil {
			return err
		}
		if rowNumber == -1 {
			return errors.New("an error occured")
		}
		if rowNumber < 0 || rowNumber >= len(allRows) {
			return errors.New("row number out of range")
		}
		rowNumbers = append(rowNumbers, rowNumber)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(rowNumbers)))
	for _, index := range rowNumbers {
		allRows = append(allRows[:index], allRows[index+1:]...)
	}

	file, err = os.Create(e.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(allRows)
	if err != nil {
		return err
	}

	return nil
}