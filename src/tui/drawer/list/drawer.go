package list

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Drawer struct {
	Item 			*widgets.List
	packetChannel 	chan []string
}

func (d Drawer) Draw() *widgets.List {
	return d.Item
}

func (d Drawer) Initialize() *widgets.List {
	d.Item.Title = "Filter By"
	d.Item.Rows = []string{
		"[1] Source Port",
		"[2] Destination Port",
		"[3] Sequence Number",
		"[4] Checksum",
	}
	d.Item.TextStyle = ui.NewStyle(ui.ColorYellow)
	d.Item.WrapText = false

	return d.Item
}

func New(pc chan []string) *Drawer {
	return &Drawer{widgets.NewList(), pc}
}