package sparkline

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"math/rand"
)

type Drawer struct {
	item 			*widgets.Sparkline
	packetChannel	chan []string
}

func (d Drawer) Draw() *widgets.SparklineGroup {
	d.item.Data = d.getData()

	d.item.LineColor = ui.ColorBlue

	slg := widgets.NewSparklineGroup(d.item)
	slg.Title = "Traffic"

	return slg
}

func (d Drawer) Initialize() *widgets.SparklineGroup {
	d.item.Data = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	d.item.LineColor = ui.ColorBlue

	slg := widgets.NewSparklineGroup(d.item)
	slg.Title = "Traffic"

	return slg
}

func New(pc chan []string) *Drawer {
	return &Drawer{widgets.NewSparkline(), pc}
}

func (d Drawer) getData() []float64 {
	data := append(d.item.Data[1:], float64(len(d.packetChannel) + rand.Intn(25)))
	return data
}