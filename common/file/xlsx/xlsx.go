package xlsx

import (
	"fmt"

	"github.com/tealeg/xlsx"
	//	"github.com/Luxurioust/excelize" //这个库最完善
)

func TestRead() {
	excelFileName := "/home/tealeg/foo.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Printf("%s\n", cell.String())
			}
		}
	}
}

func TestWrite() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet = file.AddSheet("Sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "I am a cell!"
	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
