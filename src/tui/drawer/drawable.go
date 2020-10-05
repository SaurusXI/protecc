package drawer

import (
	ui "github.com/gizak/termui/v3"
)

// Interface for widget Drawers
type drawable interface {
	Draw() *ui.Drawable
	Initialize() *ui.Drawable
}