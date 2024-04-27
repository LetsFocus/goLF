package excel

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

type Excel struct {
	filePath string
}

func NewExcel(filePath string) (*Excel, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
	}

	return &Excel{filePath: filePath}, nil
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

func (e *Excel) GetRow(id string) ([]string, error) {
	file, err := os.Open(e.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rowNum, err := e.GetRowFromId(id)
	if err != nil {
		return nil, err
	}
	if rowNum == -1 {
		return nil, errors.New("an error occured")
	}

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if rowNum >= len(rows) {
		return nil, errors.New("row number out of range")
	}

	return rows[rowNum], nil
}

func (e *Excel) IsHeaderExist() (bool, error) {
	file, err := os.Open(e.filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return false, err
	}

	if len(rows) <= 0 || len(rows[0]) <= 0 {
		return false, errors.New("header not available")
	}
	if rows[0][0] == "id" {
		return true, nil
	}
	return false, nil
}

func (e *Excel) GetLastId() (int, error) {
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

	if len(rows) <= 0 || len(rows[len(rows)-1]) <= 0 {
		return 1, nil
	}
	if rows[len(rows)-1][0] == "id" {
		return 1, nil
	}
	id, err := strconv.Atoi(rows[len(rows)-1][0])
	if err != nil {
		return -1, err
	}
	return id + 1, nil
}

func (e *Excel) GetRowFromId(id string) (int, error) {
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

	for index, row := range rows {
		if len(row) > 0 && row[0] == id {
			return index, nil
		}
	}
	return -1, nil
}

func (e *Excel) AddRow(row []string) error {
	id, err := e.GetLastId()
	if err != nil {
		return err
	}
	if id == -1 {
		return errors.New("an error occured")
	}
	row = append([]string{strconv.Itoa(id)}, row...)
	file, err := os.OpenFile(e.filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(row)
	if err != nil {
		return err
	}

	return nil
}

func (e *Excel) ReplaceRow(id string, row []string) error {
	file, err := os.Open(e.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	row = append([]string{id}, row...)
	rowNum, err := e.GetRowFromId(id)
	if err != nil {
		return err
	}
	if rowNum == -1 {
		return errors.New("an error occured")
	}

	reader := csv.NewReader(file)
	allRows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if rowNum < 0 || rowNum >= len(allRows) {
		return errors.New("row number out of range")
	}

	allRows[rowNum] = row

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

func (e *Excel) ReplaceHeader(header []string) error {
	file, err := os.Open(e.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	header = append([]string{"id"}, header...)
	reader := csv.NewReader(file)
	allRows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	validateHeader, err := e.IsHeaderExist()
	if err != nil {
		return err
	}
	if !validateHeader {
		return errors.New("header not available")
	}

	allRows[0] = header

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

func (e *Excel) AddHeader(header []string) error {
	file, err := os.Open(e.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	header = append([]string{"id"}, header...)
	reader := csv.NewReader(file)
	allRows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	validateHeader, err := e.IsHeaderExist()
	if err != nil {
		return err
	}
	if validateHeader {
		return errors.New("header already exists")
	}

	allRows = append([][]string{header}, allRows...)

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
