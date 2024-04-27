package main

import (
	"fmt"

	"github.com/LetsFocus/goLF/excel"
)

func main() {
	path := "test.csv"
	excelFile, err := excel.NewExcel(path)
	if err != nil {
		fmt.Println("Error inserting rows:", err)
		return
	}

	row := []string{"row1", "row2", "row3"}
	err = excelFile.AddRow(row)
	if err != nil {
		fmt.Println("Error adding header:", err)
		return
	}
	row = []string{"row1", "row2", "row3"}
	err = excelFile.AddRow(row)
	if err != nil {
		fmt.Println("Error adding header:", err)
		return
	}
	row = []string{"row1", "row2", "row3"}
	err = excelFile.AddRow(row)
	if err != nil {
		fmt.Println("Error adding header:", err)
		return
	}
	row = []string{"row1", "row2", "row3"}
	err = excelFile.AddRow(row)
	if err != nil {
		fmt.Println("Error adding header:", err)
		return
	}

	header := []string{"col1", "col2", "col3"}
	err = excelFile.AddHeader(header)
	if err != nil {
		fmt.Println("Error adding header:", err)
		return
	}

	header = []string{"column1", "column2", "column4"}
	err = excelFile.ReplaceHeader(header)
	if err != nil {
		fmt.Println("Error adding header:", err)
		return
	}

	row = []string{"row4", "row6", "row7"}
	err = excelFile.ReplaceRow("3", row)
	if err != nil {
		fmt.Println("Error adding header:", err)
		return
	}

	row, err = excelFile.GetRow("3")
	fmt.Println(row)
}
