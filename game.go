package main

import (
	_ "image/png"
	"time"

	"github.com/zerofelx/GameP/window"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/zerofelx/GameP/sprites"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Game",
		Bounds: pixel.R(0, 0, 720, 480),
		// VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	Player := sprites.PlayerSprite()

	window.Window(win)
	last := time.Now()
	angle := 0.0
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		angle += 3 * dt

		window.Window(win)
		mat := sprites.PlayerMatrix(win)
		mat = mat.Rotated(win.Bounds().Center(), angle)
		Player.Draw(win, mat)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
