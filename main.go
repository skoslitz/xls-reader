package main

import (
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

var applicationRoot string

func init() {

	// set repo root path
	err := os.Chdir("/Users/kOssi/gocode/src/github.com/skoslitz/xlsx-reader/files")

	if err != nil {
		fmt.Println("Can't change working directory")
	}

	applicationRoot, _ = os.Getwd()
}

func main() {
	excelFileName := applicationRoot + "/testfile.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("Read Error ", err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				val, _ := cell.String()
				fmt.Printf("%s ", val)
			}
		}
	}
}
