package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/lexcelent/sand-simulator/base"
)

const (
	// Окно приложения
	screenWidth  int = 640
	screenHeight int = 480

	// Сетка
	gridHeight int = 320
	gridWidht  int = 240

	// Размер пикселей
	pixelSize int = 2
)

func main() {
	g := base.NewWorld(gridWidht, gridHeight)

	g.Reset() // Избавляемся от nil

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sand Simulator")
	ebiten.SetTPS(60)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
