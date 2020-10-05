package tui

import (
	"log"
	"time"

	"github.com/SaurusXI/protecc/src/tui/drawer"
	ui "github.com/gizak/termui/v3"
)

func Start(packetChannel chan []string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize tui: %v", err)
	}
	defer ui.Close()

	d := drawer.New()
	render(d.Initialize())

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	// Event loop
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			default:
				d.HandleUIEvent(e)
				render(d.Draw())
			}
		case p := <-packetChannel:
			d.AddRow(p)
		case <-ticker:
			render(d.Draw())
		}
	}
}

func render(elements []ui.Drawable) {
	ui.Render(elements...)
}
