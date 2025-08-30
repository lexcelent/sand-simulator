package src

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lexcelent/sand-simulator/utils"
)

// Empty element
type NoMaterial struct {
	Coord
	nameID int
}

func (m *NoMaterial) Update(g *Game) {

}

func (m *NoMaterial) NameID() int {
	return m.nameID
}

func NewEmpty(x, y int) *NoMaterial {
	return &NoMaterial{
		Coord{x, y},
		utils.Empty,
	}
}

func (m *NoMaterial) Draw(screen *ebiten.Image) {

}
