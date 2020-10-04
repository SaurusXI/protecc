package drawer

import (
	ui "github.com/gizak/termui/v3"
)

func (d Drawer) HandleUIEvent(e ui.Event) {
	switch e.ID {
	case "j", "<Down>":
		d.list.ScrollDown()
	case "k", "<Up>":
		d.list.ScrollUp()
	}
}