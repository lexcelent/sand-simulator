package src

import "github.com/hajimehoshi/ebiten/v2"

// Material should be able to update
type Material interface {
	Update(g *Game)
	Draw(screen *ebiten.Image)
	NameID() int
}
