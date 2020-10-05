package table

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Drawer struct {
	item *widgets.Table
}

func (d Drawer) Draw() *widgets.Table {
	return d.item
}

func (d Drawer) Initialize() *widgets.Table {
	d.item.Rows = [][]string{
		[]string{"Src IP", "Src Port", "Dest IP", "Dest Port", "Window", "Checksum"},
	}
	d.item.TextStyle = ui.NewStyle(ui.ColorWhite)
	d.item.RowSeparator = true
	d.item.FillRow = true

	return d.item
}

func (d Drawer) AddRow(row []string) {
	if len(d.item.Rows) > 10 {
		d.item.Rows = append(append(d.item.Rows[:1], d.item.Rows[2:]...), row)
	} else {
		d.item.Rows = append(d.item.Rows, row)
	}
}

func New() *Drawer {
	return &Drawer{widgets.NewTable()}
}
