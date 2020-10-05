package drawer

import (
	ui "github.com/gizak/termui/v3"
)

// HandleUIEvent - draws TUI according to UI event and returns Drawables that can be rendered
func (d Drawer) HandleUIEvent(e ui.Event) []ui.Drawable {
	switch e.ID {
	case "j", "<Down>":
		d.list.ScrollDown()
		return d.drawList()
	case "k", "<Up>":
		d.list.ScrollUp()
		return d.drawList()
	default:
		return d.drawWithoutUpdate()
	}
}

// Draws UI without changing any widgets
func (d Drawer) drawWithoutUpdate() []ui.Drawable {
	slg := d.sparkline.GetSLG()
	l := d.list.Item
	table := d.table.Item

	arrange(l, table, slg)

	return []ui.Drawable {
		slg,
		l,
		table,
	}
}

// Draws UI by updating list widget only
func (d Drawer) drawList() []ui.Drawable {
	slg := d.sparkline.GetSLG()
	l := d.list.Draw()
	table := d.table.Item

	arrange(l, table, slg)

	return []ui.Drawable {
		slg,
		l,
		table,
	}
}