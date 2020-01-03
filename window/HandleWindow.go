package window

import (
	"github.com/faiface/pixel/pixelgl"
	color "golang.org/x/image/colornames"
)

func Window(win *pixelgl.Window) {
	win.Clear(color.Black)
	// win.SetSmooth(true)
}
