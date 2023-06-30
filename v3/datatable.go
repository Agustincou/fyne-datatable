package datatable

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func New(data interface{}) (*widget.Table, error) {
	var v = reflect.ValueOf(data)
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return nil, errors.New("need array or slice input")
	}

	var headers = tagValues(data)

	labelWidgetsGrid := map[string]*widget.Label{}
	mutex := &sync.Mutex{}

	table := widget.NewTable(

		func() (int, int) {
			rows := reflect.ValueOf(data).Len() + 1 // data rows with header
			cols := len(headers)
			return rows, cols
		},

		// callback fn for Create each cell.
		func() fyne.CanvasObject {
			l := widget.NewLabel("placeholder")
			l.Wrapping = fyne.TextTruncate
			return l
		},

		// callback fn for Update each cell.
		// This may trigger on initial rendering process.
		// override result of second param in NewTable()
		func(id widget.TableCellID, c fyne.CanvasObject) {
			entry := c.(*widget.Label)
			col, row := id.Col, id.Row
			if lockSuccess := mutex.TryLock(); lockSuccess {
				labelWidgetsGrid[fmt.Sprintf("%d%d", row, col)] = entry
				mutex.Unlock()
			}
			if row == 0 { // Header Row
				entry.SetText(headers[col])
				entry.Alignment = fyne.TextAlignCenter
				entry.TextStyle = fyne.TextStyle{Bold: true}

				return
			}
			// Data row
			value := reflect.ValueOf(data).Index(row - 1).Interface()
			entry.SetText(fmt.Sprintf("%v", getFieldValue(value, col)))
			entry.Alignment = fyne.TextAlignLeading
			entry.TextStyle = fyne.TextStyle{Bold: false}
		})

	table.OnSelected = func(id widget.TableCellID) {
		mutex.Lock()
		table.SetColumnWidth(id.Col, widget.NewLabel(labelWidgetsGrid[fmt.Sprintf("%d%d", id.Row, id.Col)].Text).MinSize().Width)
		mutex.Unlock()
	}

	// NOTE: Set width for each columns...
	//
	// Columns for widget.Table is automatically determined from the template object
	// specified in CreateCell (second arg of function NewTable) by default.
	// Here, the size of each column is determined separately from a tmpl of the data.
	for index, header := range headers {
		table.SetColumnWidth(index, widget.NewLabel("---"+header+"---").MinSize().Width)
	}

	return table, nil
}
