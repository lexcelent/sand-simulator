package base

import (
	"image/color"
	"math/rand"

	"github.com/lexcelent/sand-simulator/utils"
)

// Any material have: x,y and color
type Sand struct {
	Coord
	color color.RGBA
	name  int
}

// Update solid element
func (s *Sand) Update(g *Game) {
	// TODO: иногда вместо NEW стоит сделать какой-нибудь SWAP или MOVE

	if g.CurrentGrid[s.x][s.y+1].NameID() == utils.Empty {
		// Если под пикселем пусто, то перемещаем пиксель вниз
		// Текущую сетку не меняем!!!
		g.NextGrid[s.x][s.y] = NewEmpty(s.x, s.y)
		g.NextGrid[s.x][s.y+1] = NewSand(s.x, s.y+1)
	} else if g.CurrentGrid[s.x-1][s.y+1].NameID() == utils.Empty && g.CurrentGrid[s.x+1][s.y+1].NameID() == utils.Empty {
		// Если по бокам пусто, то выбираем направление наугад
		direction := rand.Int()
		if direction%2 == 0 {
			// Направляем песок влево
			g.NextGrid[s.x][s.y] = NewEmpty(s.x, s.y)
			g.NextGrid[s.x-1][s.y+1] = NewSand(s.x-1, s.y+1)
		} else {
			// Направляем песок вправо
			g.NextGrid[s.x][s.y] = NewEmpty(s.x, s.y)
			g.NextGrid[s.x+1][s.y+1] = NewSand(s.x+1, s.y+1)
		}
	} else if g.CurrentGrid[s.x-1][s.y+1].NameID() == utils.Empty {
		// Если пусто только слева
		g.NextGrid[s.x][s.y] = NewEmpty(s.x, s.y)
		g.NextGrid[s.x-1][s.y+1] = NewSand(s.x-1, s.y+1)
	} else if g.CurrentGrid[s.x+1][s.y+1].NameID() == utils.Empty {
		// Если пусто только справа
		g.NextGrid[s.x][s.y] = NewEmpty(s.x, s.y)
		g.NextGrid[s.x+1][s.y+1] = NewSand(s.x+1, s.y+1)
	} else if g.CurrentGrid[s.x][s.y+1].NameID() == utils.Water {
		// Взаимодействие с водой
		g.NextGrid[s.x][s.y] = NewWater(s.x, s.y)
		g.NextGrid[s.x][s.y+1] = NewSand(s.x, s.y+1)
	}
}

func (s *Sand) NameID() int {
	return s.name
}

func NewSand(x, y int) *Sand {
	return &Sand{
		Coord: Coord{x: x, y: y},
		color: utils.YELLOW,
		name:  utils.Sand}
}
