package table

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Drawer struct {
	Item 			*widgets.Table
	packetChannel	chan []string
}

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

func (d Drawer) Initialize() *widgets.Table {
	d.Item.Rows = [][]string{
		[]string{"Src IP", "Src Port", "Dest IP", "Dest Port", "Window", "Checksum"},
	}
	d.Item.TextStyle = ui.NewStyle(ui.ColorWhite)
	d.Item.RowSeparator = true
	d.Item.FillRow = true

	return d.Item
}

func (d Drawer) AddRow(row []string) {
	if len(d.Item.Rows) > 10 {
		d.Item.Rows = append(append(d.Item.Rows[:1], d.Item.Rows[2:]...), row)
	} else {
		d.Item.Rows = append(d.Item.Rows, row)
	}
}

func New(pc chan []string) *Drawer {
	return &Drawer{widgets.NewTable(), pc}
}
