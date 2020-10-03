package tui

import (
	"time"
	"log"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "Hello World!"
	p.SetRect(0, 0, 25, 5)


	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			draw()
		}
	}
}

func draw() {
	slg := drawSparklineGroup()
	table := drawTable()
	l := drawList()

	l.SetRect(0, 0, 50, 5)
	slg.SetRect(50, 0, 100, 5)
	table.SetRect(0, 5, 100, 20)

	ui.Render(l, slg, table)
}

func drawSparklineGroup() *widgets.SparklineGroup {
	sl := widgets.NewSparkline()
	sl.Data = getData()
	sl.LineColor = ui.ColorBlue

	slg := widgets.NewSparklineGroup(sl)
	slg.Title = "Traffic"

	return slg
}

func drawList() *widgets.List {
	l := widgets.NewList()
	l.Title = "Filter By"
	l.Rows = []string{
		"[1] Source Port",
		"[2] Destination Port",
		"[3] Sequence Number",
		"[4] Checksum",
	}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false

	return l
}

func drawTable() *widgets.Table {
	table := widgets.NewTable()
	table.Rows = [][]string {
		[]string {"Src Port", "Dest Port", "Seq No.", "Data Offset", "Window", "Checksum", "Urgent"},
	}
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.RowSeparator = true
	table.FillRow = true

	return table
}

func getData() []float64 {
	data := []float64{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}
	return data
}