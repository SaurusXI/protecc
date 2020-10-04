package tui

import (
	"time"
	"log"
	ui "github.com/gizak/termui/v3"
	"github.com/SaurusXI/protecc/src/tui/drawer"
)

func Start() {
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
		case <-ticker:
			render(d.Draw())
		}
	}
}

func render(elements []ui.Drawable) {
	ui.Render(elements...)
}