package main

import (
	"fmt"

	"github.com/LetsFocus/goLF/excel"
)

func main() {
	csvPath := "csv.csv"
	excelFile, _ := excel.InitExcel(csvPath, "age", true)
	headers := make(map[string]string, 3)
	headers["age"] = "new_age"
	headers["name"] = "new_name"
	headers["gpa"] = "new_gpa"
	err := excelFile.ReplaceHeaderName(headers)
	if err!= nil{
		fmt.Println(err)
	}
	err := excelFile.AddRow()
}
