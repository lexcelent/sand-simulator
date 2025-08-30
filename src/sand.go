package src

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	// TODO: Не нравится вызов NameID... Но возможно это правильно

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

func (s *Sand) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(s.x), float32(s.y), float32(1*pixelSize), float32(1*pixelSize), s.color, false)
}

func (s *Sand) NameID() int {
	return s.name
}

func NewSand(x, y int) *Sand {

	// Оттенки желтого
	colors := []color.RGBA{
		color.RGBA{255, 255, 0, 0},
		color.RGBA{255, 207, 64, 0},
		color.RGBA{244, 169, 0, 0},
		color.RGBA{205, 164, 52, 0},
	}
	randomIndex := rand.Intn(len(colors))

	return &Sand{
		Coord: Coord{x: x, y: y},
		color: colors[randomIndex],
		name:  utils.Sand,
	}
}
