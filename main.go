package exel_parser

import (
	"fmt"
	"github.com/dimanlin/exel_parser/pkg/excelize"
)

type XLSXReader struct {
	filename string
}

func main() {
	xlsxReader := XLSXReader{"example.xlsx"}
	row, err := xlsxReader.GetRows("Sheet1")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(row)

	cell, err := xlsxReader.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cell)
}