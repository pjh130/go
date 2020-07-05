package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//创建 Excel 文档
func CreateExcel() {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet2")
	// 设置单元格的值
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		println(err.Error())
	}
}

//读取 Excel 文档
func ReadExcel() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		println(err.Error())
		return
	}
	println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			print(colCell, "\t")
		}
		println()
	}
}

//在 Excel 文档中创建图表
func CreateChart() {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	if err := f.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`); err != nil {
		println(err.Error())
		return
	}
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		println(err.Error())
	}
}

//向 Excel 文档中插入图片
func InsertPicture() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	// 插入图片
	if err := f.AddPicture("Sheet1", "A2", "image.png", ""); err != nil {
		println(err.Error())
	}
	// 在工作表中插入图片，并设置图片的缩放比例
	if err := f.AddPicture("Sheet1", "D2", "image.jpg", `{"x_scale": 0.5, "y_scale": 0.5}`); err != nil {
		println(err.Error())
	}
	// 在工作表中插入图片，并设置图片的打印属性
	if err := f.AddPicture("Sheet1", "H2", "image.gif", `{"x_offset": 15, "y_offset": 10, "print_obj": true, "lock_aspect_ratio": false, "locked": false}`); err != nil {
		println(err.Error())
	}
	// 保存文件
	if err = f.Save(); err != nil {
		println(err.Error())
	}
}

//复制工作表
// 名称为 Sheet1 的工作表已经存在 ...
//index := f.NewSheet("Sheet2")
//err := f.CopySheet(1, index)
func CopySheet() {
	f, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		println(err.Error())
		return
	}

	index := f.GetSheetIndex("模板")
	index1 := f.NewSheet("表1")
	f.NewSheet("表2")

	//把表2改名成表222
	f.SetSheetName("表2", "表2222")

	f.SetCellValue("表1", "A1", "Hello world.")
	f.SetCellValue("表1", "B2", 100)

	index2 := f.NewSheet("copy")

	if true {
		//拷贝一个简单的表
		err = f.CopySheet(index1, index2)
		if err != nil {
			println(err.Error())
			return
		}
	} else {
		//拷贝一个复杂的表
		err = f.CopySheet(index, index2)
		if err != nil {
			println(err.Error())
			return
		}
	}

	//设置超链接
	err = f.SetCellHyperLink("股票池", "B2", "表1!A1", "Location")
	err = f.SetCellHyperLink("股票池", "B3", "表1!B2", "Location")
	if err != nil {
		println(err.Error())
		return
	}

	err = f.SaveAs("new.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
}
