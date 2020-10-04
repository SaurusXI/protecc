package list

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Drawer struct {
	item *widgets.List
}

func (d Drawer) Draw() *widgets.List {
	return d.item
}

func (d Drawer) Initialize() *widgets.List {
	d.item.Title = "Filter By"
	d.item.Rows = []string{
		"[1] Source Port",
		"[2] Destination Port",
		"[3] Sequence Number",
		"[4] Checksum",
	}
	d.item.TextStyle = ui.NewStyle(ui.ColorYellow)
	d.item.WrapText = false

	return d.item
}

func New() *Drawer {
	return &Drawer{widgets.NewList()}
}