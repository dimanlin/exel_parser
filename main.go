package main

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type XLSXReader struct {
	filename string
}

func (xlsxReader *XLSXReader) GetCellValue(sheet string, axis string) string {
	if len(xlsxReader.filename) == 0 {
		fmt.Println("No file name.")
		return ""
	}

	f, err := excelize.OpenFile(xlsxReader.filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	cell, err := f.GetCellValue(sheet, axis)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return cell
}

func (xlsxReader *XLSXReader) GetRows(sheet string) ([][]string, error) {
	if len(xlsxReader.filename) == 0 {
		return [][]string{}, errors.New("No file name.")
	}

	f, err := excelize.OpenFile(xlsxReader.filename)
	if err != nil {
		fmt.Println(err)
		return [][]string{}, err
	}

	rows, err := f.GetRows(sheet)

	return rows, err
}

func main() {
	xlsxReader := XLSXReader{"example.xlsx"}
	rows, err := xlsxReader.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	//a := xlsxReader.GetCellValue("Sheet1", "B2")
	fmt.Println(rows)
}