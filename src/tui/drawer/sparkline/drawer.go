package sparkline

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Drawer struct {
	item *widgets.Sparkline
}

func (d Drawer) Draw() *widgets.SparklineGroup {
	d.item.Data = getData()

	d.item.LineColor = ui.ColorBlue

	slg := widgets.NewSparklineGroup(d.item)
	slg.Title = "Traffic"

	return slg
}

func (d Drawer) Initialize() *widgets.SparklineGroup {
	return d.Draw()
}

func New() *Drawer {
	return &Drawer{widgets.NewSparkline()}
}

func getData() []float64 {
	data := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	return data
}