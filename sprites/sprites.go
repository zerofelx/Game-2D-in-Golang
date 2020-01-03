package sprites

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func PlayerSprite() *pixel.Sprite {
	pic, err := LoadPicture("assets/sprites/player.png")
	if err != nil {
		panic(err)
	}

	Player := pixel.NewSprite(pic, pic.Bounds())

	return Player
}

func PlayerMatrix(win *pixelgl.Window) pixel.Matrix {
	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(5, 5))

	return mat
}
