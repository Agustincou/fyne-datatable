package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	datatable "github.com/agustincou/fyne-datatable/v3"
)

type account struct {
	ID             string  `column:"id"`
	Name           string  `column:"name"`
	Phone          string  `column:"-"`
	Birthday       string  `column:"birth day"`
	Integer64Value int64   `column:"int64 value"`
	Floa64Value    float64 `column:"float64 value"`
	Integer32Value int32   `column:"int32 value"`
	Floa32Value    float32 `column:"float32 value"`
}

func main() {
	a := app.New()
	w := a.NewWindow("Account List")
	data := []account{
		{"0f670e49-6252-43dc", "Robert L. Taylor", "706-767-8575", "December 25, 1970", 12, 12.51, 16, 16.6346623},
		{"0223e482-5a3c-4b58", "Kento Agano", "936-347-9392", "May 22, 1943", 17, 17.51, 16, 16},
		{"311cec67-fc48-4bb3", "Mahito Fujiwara", "574-282-4340", "July 27, 1983", 15, 15.51, 16, 99987.6346623},
		{"34977adb-783a-49db", "Suzanne R. Gonzalezxxxxxxxxxxx", "603-736-6867", "June 12, 1994", 7685, 685.51, 76, 45453.6346623},
		{"0b809125-abaf-423b", "Linda M. Bailey", "317-738-7776", "January 2, 1962", 3457, 34566999456.51, 343453, 23436.6346623},
	}
	table, _ := datatable.New(data)
	w.SetContent(container.NewMax(table))
	w.Resize(fyne.NewSize(600, 150))
	w.ShowAndRun()
}
