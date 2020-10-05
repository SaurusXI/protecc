package table

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Drawer for Table widget 
type Drawer struct {
	Item 			*widgets.Table
	packetChannel	chan []string
}

// Draw table by flushing all packets in packetChannel
func (d Drawer) Draw() *widgets.Table {
	for {
		select {
		case p := <- d.packetChannel:
			d.AddRow(p)
		default:
			return d.Item
		}
	}
}

// Initialize table with column names
func (d Drawer) Initialize() *widgets.Table {
	d.Item.Rows = [][]string{
		[]string{"Src IP", "Src Port", "Dest IP", "Dest Port", "Window", "Checksum"},
	}
	d.Item.TextStyle = ui.NewStyle(ui.ColorWhite)
	d.Item.RowSeparator = true
	d.Item.FillRow = true

	return d.Item
}

// AddRow inserts a row into the table
func (d Drawer) AddRow(row []string) {
	if len(d.Item.Rows) > 10 {
		d.Item.Rows = append(append(d.Item.Rows[:1], d.Item.Rows[2:]...), row)
	} else {
		d.Item.Rows = append(d.Item.Rows, row)
	}
}

// New (constructor) for Table Drawer
func New(pc chan []string) *Drawer {
	return &Drawer{widgets.NewTable(), pc}
}
