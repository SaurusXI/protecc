package sparkline

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Drawer struct {
	Item 			*widgets.Sparkline
	packetChannel	chan []string
}

func (d Drawer) Draw() *widgets.SparklineGroup {
	d.Item.Data = d.getData()

	d.Item.LineColor = ui.ColorBlue

	slg := widgets.NewSparklineGroup(d.Item)
	slg.Title = "Traffic"

	return slg
}

func (d Drawer) Initialize() *widgets.SparklineGroup {
	d.Item.Data = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	d.Item.LineColor = ui.ColorBlue

	slg := widgets.NewSparklineGroup(d.Item)
	slg.Title = "Traffic"

	return slg
}

func New(pc chan []string) *Drawer {
	return &Drawer{widgets.NewSparkline(), pc}
}

func (d Drawer) getData() []float64 {
	data := append(d.Item.Data[1:], float64(len(d.packetChannel)))
	return data
}

func (d Drawer) GetSLG() *widgets.SparklineGroup {
	slg := widgets.NewSparklineGroup(d.Item)
	slg.Title = "Traffic"

	return slg
}