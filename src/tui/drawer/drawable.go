package drawer

import (
	ui "github.com/gizak/termui/v3"
)

type drawable interface {
	Draw() *ui.Drawable
	Initialize() *ui.Drawable
}