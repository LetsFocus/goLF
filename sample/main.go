package main

import (
	"fmt"

	"github.com/LetsFocus/goLF/excel"
)

func main() {
	csvPath := "csv.csv"
	configPath := "config.json"
	excelFile, _ := excel.NewExcel(csvPath, configPath)
	err := excelFile.AddHeader("auto", []string{"name", "age"})
	if err!= nil{
		fmt.Println(err)
	}
}
