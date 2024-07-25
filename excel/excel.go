package excel

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
)

func NewExcel(filePath string, primaryKey string, doesHeaderExist bool) (*Excel, error) {
	writeFile, readFile, err := validateFilePath(filePath)
	if err != nil {
		return nil, err
	}

	var headersMap map[string]int
	reader := csv.NewReader(readFile)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	rowsCount := len(rows)

	if doesHeaderExist {
		headersMap = buildHeaderMap(rows[0], primaryKey)
	}

	var writeConfig, readConfig *os.File
	if primaryKey == "auto_id" {
		var autoId int = 0
		writeConfig, readConfig, err = validateFilePath(CONFIG_PATH)
		if err != nil {
			return nil, err
		}
		if doesHeaderExist {
			autoId = rowsCount - 1
		}
		writeToConfig(autoId)
	}

	return &Excel{
		filePath:        filePath,
		configPath:      CONFIG_PATH,
		writeFile:       writeFile,
		readFile:        readFile,
		writeConfig:     writeConfig,
		readConfig:      readConfig,
		primaryKey:      primaryKey,
		doesHeaderExist: doesHeaderExist,
		headersMap:      headersMap, header: nil}, nil
}

func (e *Excel) GetAllRows() ([][]string, error) {
	reader := csv.NewReader(e.readFile)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (e *Excel) buildHeaderMap(headers []string) map[string]int {
	headersMap := make(map[string]int)
	startIndex := 0
	if e.primaryKey == "auto_id" {
		headersMap[e.primaryKey] = startIndex
		startIndex = 1
	}
	for index, field := range headers {
		headersMap[field] = index + startIndex
	}
	return headersMap
}

func (e *Excel) checkAndIncrementAutoID() (int, error) {
	data, err := io.ReadAll(e.readConfig)
	if err != nil {
		return -1, err
	}

	var config map[string]int
	if err := json.Unmarshal(data, &config); err != nil {
		return -1, err
	}

	autoID := config["AUTO_ID"]
	config["AUTO_ID"] = autoID + 1

	if err := writeToConfig(autoID + 1); err != nil {
		return -1, err
	}

	return autoID, nil
}

func (e *Excel) checkPrimaryKey(primaryKey string) (bool, error) {
	if e.primaryKey == "auto_id" {
		return false, nil
	}

	primaryKeyPosition, exists := e.headersMap[e.primaryKey]
	if !exists {
		return false, errors.New("primary key does not exist in headers")
	}

	rows, err := e.GetAllRows()
	if err != nil {
		return false, err
	}

	for _, row := range rows {
		if row[primaryKeyPosition] == primaryKey {
			return true, nil
		}
	}
	return false, nil
}

func (e *Excel) AddHeader(headers []string) error {
	if e.doesHeaderExist {
		return errors.New("header already present")
	}

	e.headersMap = e.buildHeaderMap(headers)

	rows, err := e.GetAllRows()
	if err != nil {
		return err
	}

	if e.primaryKey == "auto_id" {
		headers = append([]string{e.primaryKey}, headers...)
	}

	allRows := append([][]string{headers}, rows...)
	writer := csv.NewWriter(e.writeFile)
	if err := writer.WriteAll(allRows); err != nil {
		return err
	}

	e.doesHeaderExist = true
	return nil
}

func (e *Excel) ReplaceHeaderName(headers map[string]string) error {
	if !e.doesHeaderExist {
		return errors.New("header not present")
	}

	rows, err := e.GetAllRows()
	if err != nil {
		return err
	}

	newHeaders := make([]string, len(rows[0]))

	for oldName, newName := range headers {
		position, exists := e.headersMap[oldName]
		if !exists {
			return errors.New("header doesn't exist")
		}
		if oldName == e.primaryKey {
			e.primaryKey = newName
		}
		newHeaders[position] = newName
		delete(e.headersMap, oldName)
		e.headersMap[newName] = position
	}
	_, exists := e.headersMap["auto_id"]
	if exists {
		e.headersMap["auto_id"] = 0
		newHeaders[0] = "auto_id"
	}
	rows[0] = newHeaders

	writer := csv.NewWriter(e.writeFile)
	if err := writer.WriteAll(rows); err != nil {
		return err
	}

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
		isPrimaryKeyDuplicated, err := e.checkPrimaryKey(row[e.headersMap[e.primaryKey]])
		if err != nil {
			return err
		}
		if isPrimaryKeyDuplicated {
			return errors.New("primary key duplicated")
		}
	}
	rows = append(rows, row)

	writer := csv.NewWriter(e.writeFile)
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

	rowToAdd := make([]string, len(row)+1)
	if h.excel.primaryKey == "auto_id" {
		currentId, err := h.excel.checkAndIncrementAutoID()
		if err != nil {
			return err
		}
		rowToAdd[0] = strconv.Itoa(currentId)
		for i, header := range h.headers {
			rowToAdd[h.excel.headersMap[header]] = row[i]
		}
	} else {
		for i, header := range h.headers {
			rowToAdd[h.excel.headersMap[header]] = row[i]
		}
		isPrimaryKeyDuplicated, err := h.excel.checkPrimaryKey(row[h.excel.headersMap[h.excel.primaryKey]])
		if err != nil {
			return err
		}
		if isPrimaryKeyDuplicated {
			return errors.New("primary key duplicated")
		}
		rowToAdd = rowToAdd[1:]
	}

	rows = append(rows, rowToAdd)

	writer := csv.NewWriter(h.excel.writeFile)
	if err := writer.WriteAll(rows); err != nil {
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

func (e *Excel) getRowNumFromId(rows [][]string, primaryKey string) (int, error) {
	if len(rows) == 0 {
		return -1, errors.New("ID not available")
	}

	primaryKeyPosition, exists := e.headersMap[e.primaryKey]
	if !exists {
		return -1, errors.New("primary key not found in headers")
	}

	for index, row := range rows {
		if primaryKeyPosition < len(row) && row[primaryKeyPosition] == primaryKey {
			return index, nil
		}
	}
	return -1, nil
}

func (e *Excel) GetRow(primaryKey string) ([]string, error) {
	rows, err := e.GetAllRows()
	if err != nil {
		return nil, err
	}

	rowNumber, err := e.getRowNumFromId(rows, primaryKey)
	if err != nil {
		return nil, err
	}
	if rowNumber == -1 {
		return nil, errors.New("an error occured")
	}

	if rowNumber >= len(rows) {
		return nil, errors.New("row number out of range")
	}

	return rows[rowNumber], nil
}

func (e *Excel) ReplaceRow(primaryKey string, row []string) error {
	rows, err := e.GetAllRows()
	if err != nil {
		return err
	}

	rowNumber, err := e.getRowNumFromId(rows, primaryKey)
	if err != nil {
		return err
	}
	if rowNumber == -1 {
		return errors.New("an error occured")
	}

	if rowNumber < 0 || rowNumber >= len(rows) {
		return errors.New("row number out of range")
	}

	if e.primaryKey == "auto_id" {
		row = append([]string{primaryKey}, row...)
	}

	rows[rowNumber] = row

	writer := csv.NewWriter(e.writeFile)
	err = writer.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

func (h *Header) ReplaceRow(primaryKey string, row []string) error {
	rows, err := h.excel.GetAllRows()
	if err != nil {
		return err
	}

	rowNumber, err := h.excel.getRowNumFromId(rows, primaryKey)
	if err != nil {
		return err
	}
	if rowNumber == -1 {
		return errors.New("row not found")
	}

	if rowNumber < 0 || rowNumber >= len(rows) {
		return errors.New("row number out of range")
	}

	rowToAdd := make([]string, len(row)+1)
	if h.excel.primaryKey == "auto_id" {
		rowToAdd[0] = primaryKey
		for i, header := range h.headers {
			rowToAdd[h.excel.headersMap[header]] = row[i]
		}
	} else {
		for i, header := range h.headers {
			rowToAdd[h.excel.headersMap[header]] = row[i]
		}
		rowToAdd = rowToAdd[1:]
	}

	rows[rowNumber] = rowToAdd

	writer := csv.NewWriter(h.excel.writeFile)
	if err := writer.WriteAll(rows); err != nil {
		return err
	}

	return nil
}

func (e *Excel) UpdateValue(primaryKey, header, value string) error {
	if header == e.primaryKey {
		return errors.New("primary key can't be updated")
	}

	rows, err := e.GetAllRows()
	if err != nil {
		return err
	}

	rowNumber, err := e.getRowNumFromId(rows, primaryKey)
	if err != nil {
		return err
	}
	if rowNumber == -1 {
		return errors.New("an error occured")
	}

	if rowNumber < 0 || rowNumber >= len(rows) {
		return errors.New("row number out of range")
	}

	rows[rowNumber][e.headersMap[header]] = value

	writer := csv.NewWriter(e.writeFile)
	err = writer.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

func (e *Excel) DeleteRows(primaryKeys []string) error {
	rows, err := e.GetAllRows()
	if err != nil {
		return err
	}

	rowNumbers := make(map[int]struct{})
	for _, primaryKey := range primaryKeys {
		rowNumber, err := e.getRowNumFromId(rows, primaryKey)
		if err != nil {
			return err
		}
		if rowNumber == -1 {
			return errors.New("row not found")
		}
		if rowNumber < 0 || rowNumber >= len(rows) {
			return errors.New("row number out of range")
		}
		rowNumbers[rowNumber] = struct{}{}
	}

	newRows := make([][]string, 0, len(rows))
	for i, row := range rows {
		if _, exists := rowNumbers[i]; !exists {
			newRows = append(newRows, row)
		}
	}

	writer := csv.NewWriter(e.writeFile)
	if err := writer.WriteAll(newRows); err != nil {
		return err
	}

	return nil
}

func (e *Excel) CloseExcel() {
	e.readFile.Close()
	e.writeFile.Close()
}
