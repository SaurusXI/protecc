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

	d := drawer.New(packetChannel)
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
				render(d.HandleUIEvent(e))
			}
		case <-ticker:
			render(d.Draw())
		}
	}
}

func render(elements []ui.Drawable) {
	ui.Render(elements...)
}
