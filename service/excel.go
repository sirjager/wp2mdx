package service

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func (w *WordressSite) Run() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetName := "Sheet2"

	// Create a new sheet.
	index, err := f.NewSheet(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set value of a cell.

	headers := []string{
		"Post Name",
		"Post URL",
		"Category",
		"Month",
		"Editor",
		"Folder",
		"PDF Name",
		"PDF URL",
		"Comment",
		"SEO",
		"Image",
		"Social Share",
	} // Replace with your header values
	// Set header row
	for i, header := range headers {
		colName := string(rune('A' + i))
		cell := colName + "1"
		f.SetCellValue(sheetName, cell, header)
	}

	for rowIndex, post := range w.AllPosts {
		for colIndex, header := range headers {
			colName := string('A' + colIndex)
			cell := colName + strconv.Itoa(rowIndex+2)
			switch header {
			case "Post Name":
				f.SetCellValue(sheetName, cell, post.Title.Rendered)
			case "Post URL":
				f.SetCellValue(sheetName, cell, post.Link)
			case "Category":
				categories := w.GetCategories(&post)
				if len(categories) > 0 {
					f.SetCellValue(sheetName, cell, categories[0].Name)
				}
			case "Month":
				f.SetCellValue(sheetName, cell, post.Date)
			case "Editor":
				f.SetCellValue(sheetName, cell, post.Embed.Author[0].Name)
			case "Folder":
				f.SetCellValue(sheetName, cell, "")
			case "PDF Name":
				f.SetCellValue(sheetName, cell, "")
			case "PDF URL":
				f.SetCellValue(sheetName, cell, "")
			case "Comment":
				f.SetCellValue(sheetName, cell, "")
			case "SEO":
				f.SetCellValue(sheetName, cell, "")
			case "Image":
				f.SetCellValue(sheetName, cell, "")
			case "Social Share":
				f.SetCellValue(sheetName, cell, "")
			}
		}
	}

	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
