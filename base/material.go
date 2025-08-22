package base

import (
	"image/color"

	"github.com/lexcelent/sand-simulator/utils"
)

// Material should be able to update
type Material interface {
	Update(g *Game)
	NameID() int
}

// Any material have: x,y and color
type Solid struct {
	Coord
	color color.RGBA
	name  int
}

// Пока не работает...
func (s *Solid) Update(g *Game) {
	if g.CurrentGrid[s.x][s.y+1].NameID() == utils.Empty {
		// Если под пикселем пусто, то перемещаем пиксель вниз
		// Текущую сетку не меняем!!!
		g.NextGrid[s.x][s.y] = NewEmpty(s.x, s.y)
		g.NextGrid[s.x][s.y+1] = NewSand(s.x, s.y)
	}
}

func (s *Solid) NameID() int {
	return s.name
}

func NewSand(x, y int) *Solid {
	return &Solid{
		Coord: Coord{x: x, y: y},
		color: utils.YELLOW,
		name:  utils.Sand}
}

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
