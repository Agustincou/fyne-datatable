package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TappableCell struct {
	widget.Label
	fyne.Tappable

	table *widget.Table
	columnNumber int
}

func NewTappableCell() *TappableCell {
	tappableCell := &TappableCell{}
	tappableCell.ExtendBaseWidget(tappableCell)

	return tappableCell
}

func (t *TappableCell) Tapped(_ *fyne.PointEvent) {
	if t.table != nil {
		t.table.SetColumnWidth(t.columnNumber, widget.NewLabel("---" + t.Text + "---").MinSize().Width)
	}
}

func (t *TappableCell) TappedSecondary(_ *fyne.PointEvent) {
}
