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
	d.item.Rows = [][]string {
		[]string {"Src Port", "Dest Port", "Seq No.", "Data Offset", "Window", "Checksum", "Urgent"},
	}
	d.item.TextStyle = ui.NewStyle(ui.ColorWhite)
	d.item.RowSeparator = true
	d.item.FillRow = true
	
	return d.item
}	

func New() *Drawer {
	return &Drawer{widgets.NewTable()}
}
