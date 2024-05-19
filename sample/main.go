package main

import (
	"fmt"

	"github.com/LetsFocus/goLF/excel"
)

func main() {
	csvPath := "csv.csv"
	// excelFile, _ := excel.InitExcel(csvPath, "age", false)
	// err := excelFile.AddHeader([]string{"name", "age", "marks"})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = excelFile.AddHeader([]string{"name", "age", "marks"})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// newHeaders := make(map[string]string)
	// newHeaders["name"] = "new_name"
	// newHeaders["age"] = "new_age"
	// newHeaders["marks"] = "new_marks"
	// err = excelFile.ReplaceHeaderName(newHeaders)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = excelFile.AddRow([]string{"babaji", "22", "98"})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = excelFile.Columns([]string{"new_age", "new_marks", "new_name"}).AddRow([]string{"20", "99", "purna"})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = excelFile.ReplaceRow("22", []string{"babaji-pattabhiram", "22", "97"})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = excelFile.Columns([]string{"new_age", "new_marks", "new_name"}).ReplaceRow("20", []string{"21", "98", "lakshmi-purna"})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	excelFile, _ := excel.InitExcel(csvPath, "auto_id", false)
	err := excelFile.AddHeader([]string{"name", "age", "marks"})
	if err != nil {
		fmt.Println(err)
	}
	err = excelFile.AddHeader([]string{"name", "age", "marks"})
	if err != nil {
		fmt.Println(err)
	}
	newHeaders := make(map[string]string)
	newHeaders["name"] = "new_name"
	newHeaders["age"] = "new_age"
	newHeaders["marks"] = "new_marks"
	err = excelFile.ReplaceHeaderName(newHeaders)
	if err != nil {
		fmt.Println(err)
	}
	err = excelFile.AddRow([]string{"babaji", "22", "98"})
	if err != nil {
		fmt.Println(err)
	}
	err = excelFile.Columns([]string{"new_age", "new_marks", "new_name"}).AddRow([]string{"22", "99", "purna"})
	if err != nil {
		fmt.Println(err)
	}
	err = excelFile.ReplaceRow("0", []string{"babaji-pattabhiram", "22", "97"})
	if err != nil {
		fmt.Println(err)
	}
	err = excelFile.Columns([]string{"new_age", "new_marks", "new_name"}).ReplaceRow("1", []string{"21", "98", "lakshmi-purna"})
	if err != nil {
		fmt.Println(err)
	}
}
