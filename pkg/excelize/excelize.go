package excelize

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type XLSXReader struct {
	filename string
}

func (xlsxReader *XLSXReader) GetCellValue(sheet string, axis string) (string, error) {
	f := xlsxReader.openFile()

	cell, err := f.GetCellValue(sheet, axis)
	if err != nil {
		return "", err
	}
	return cell, err
}

func (xlsxReader *XLSXReader) openFile() *excelize.File {
	if len(xlsxReader.filename) == 0 {
		return nil
	}

	f, err := excelize.OpenFile(xlsxReader.filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return f
}

func (xlsxReader *XLSXReader) GetRows(sheet string) ([][]string, error) {
	f := xlsxReader.openFile()

	rows, err := f.GetRows(sheet)

	return rows, err
}