package main

import (
	"fmt"
	"github.com/dimanlin/exel_parser"
)

func main() {
	var parser = exel_parser.XLSXReader{"example.xlsx"}

	// Cell
	a, _ := parser.GetCellValue("Sheet1", "B2")
	fmt.Println(a)
	fmt.Println("================")

	// Rows
	rows, _ := parser.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
