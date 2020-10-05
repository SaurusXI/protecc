package list

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Drawer for command list widget
type Drawer struct {
	Item 			*widgets.List
	packetChannel 	chan []string
}

// Draw list contained by Drawer
func (d Drawer) Draw() *widgets.List {
	return d.Item
}

// Initialize and draw list contained by Drawer
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

// New (constructor) for List Drawer
func New(pc chan []string) *Drawer {
	return &Drawer{widgets.NewList(), pc}
}