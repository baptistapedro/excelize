package myfuzz

import (
    "github.com/xuri/excelize/v2"
)

func Fuzz(data []byte) int {
    f := excelize.NewFile()
    // Create a new sheet.
    index,_ := f.NewSheet("Sheet2")
    // Set value of a cell.
    f.SetCellValue("Sheet2", string(data), string(data))
    f.SetCellValue("Sheet1", "B2", 100)

    // Set active sheet of the workbook.
    f.SetActiveSheet(index)
    return 0
}
