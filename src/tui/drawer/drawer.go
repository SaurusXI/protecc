package drawer

import (
	ld "github.com/SaurusXI/protecc/src/tui/drawer/list"
	sd "github.com/SaurusXI/protecc/src/tui/drawer/sparkline"
	td "github.com/SaurusXI/protecc/src/tui/drawer/table"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Drawer that combines all widget drawers
type Drawer struct {
	list      		*ld.Drawer
	table     		*td.Drawer
	sparkline 		*sd.Drawer
	packetChannel	chan []string
}

// New (constructor) for Drawer
func New(pc chan []string) *Drawer {
	d := Drawer{nil, nil, nil, nil}
	d.list = ld.New(pc)
	d.sparkline = sd.New(pc)
	d.table = td.New(pc)
	d.packetChannel = pc

	return &d
}

// Initialize all widget Drawers and arrange 
func (d Drawer) Initialize() []ui.Drawable {
	slg := d.sparkline.Initialize()
	l := d.list.Initialize()
	table := d.table.Initialize()

	arrange(l, table, slg)

	return []ui.Drawable{
		slg,
		l,
		table,
	}
}

// Draw all widgets
func (d Drawer) Draw() []ui.Drawable {
	slg := d.sparkline.Draw()
	l := d.list.Draw()
	table := d.table.Draw()

	arrange(l, table, slg)

	return []ui.Drawable{
		slg,
		l,
		table,
	}
}

// Arrange all widgets spatially
func arrange(l *widgets.List, table *widgets.Table, slg *widgets.SparklineGroup) {
	l.SetRect(0, 0, 50, 5)
	slg.SetRect(50, 0, 100, 5)
	table.SetRect(0, 5, 100, 20)
}
