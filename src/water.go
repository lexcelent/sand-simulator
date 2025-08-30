package src

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/lexcelent/sand-simulator/utils"
)

// Any material have: x,y and color
type Water struct {
	Coord
	color color.RGBA
	name  int
}

func (s *Water) NameID() int {
	return s.name
}

func NewWater(x, y int) *Water {
	return &Water{
		Coord: Coord{x: x, y: y},
		color: utils.BLUE,
		name:  utils.Water}
}

func (w *Water) Update(g *Game) {
	// Обновить пиксель воды. Сделал тут дополнительную проверку для NextGrid, чтобы пиксели не исчезали.
	// Наверняка есть другое решение
	if g.CurrentGrid[w.x][w.y+1].NameID() == utils.Empty && g.NextGrid[w.x][w.y+1].NameID() == utils.Empty {
		// Если под пикселем пусто, то перемещаем пиксель вниз
		g.NextGrid[w.x][w.y] = NewEmpty(w.x, w.y)
		g.NextGrid[w.x][w.y+1] = NewWater(w.x, w.y+1)
	} else if g.CurrentGrid[w.x-1][w.y+1].NameID() == utils.Empty {
		// Если пусто только слева
		if g.NextGrid[w.x-1][w.y+1].NameID() == utils.Empty {
			g.NextGrid[w.x][w.y] = NewEmpty(w.x, w.y)
			g.NextGrid[w.x-1][w.y+1] = NewWater(w.x-1, w.y+1)
		}
	} else if g.CurrentGrid[w.x+1][w.y+1].NameID() == utils.Empty {
		// Если пусто только справа
		if g.NextGrid[w.x+1][w.y+1].NameID() == utils.Empty {
			g.NextGrid[w.x][w.y] = NewEmpty(w.x, w.y)
			g.NextGrid[w.x+1][w.y+1] = NewWater(w.x+1, w.y+1)
		}
	} else if g.CurrentGrid[w.x-1][w.y].NameID() == utils.Empty && g.CurrentGrid[w.x+1][w.y].NameID() == utils.Empty {
		// Проверяем слева и справа
		direction := rand.Int()
		if direction%2 == 0 && g.NextGrid[w.x-1][w.y].NameID() == utils.Empty {
			// Направляем воду влево
			g.NextGrid[w.x][w.y] = NewEmpty(w.x, w.y)
			g.NextGrid[w.x-1][w.y] = NewWater(w.x-1, w.y)
		} else if direction%2 != 0 && g.CurrentGrid[w.x+1][w.y].NameID() == utils.Empty {
			// Направляем воду вправо
			g.NextGrid[w.x][w.y] = NewEmpty(w.x, w.y)
			g.NextGrid[w.x+1][w.y] = NewWater(w.x+1, w.y)
		}
	} else if g.CurrentGrid[w.x-1][w.y].NameID() == utils.Empty {
		if g.NextGrid[w.x-1][w.y].NameID() == utils.Empty {
			g.NextGrid[w.x][w.y] = NewEmpty(w.x, w.y)
			g.NextGrid[w.x-1][w.y] = NewWater(w.x-1, w.y)
		}
	} else if g.CurrentGrid[w.x+1][w.y].NameID() == utils.Empty {
		if g.NextGrid[w.x+1][w.y].NameID() == utils.Empty {
			g.NextGrid[w.x][w.y] = NewEmpty(w.x, w.y)
			g.NextGrid[w.x+1][w.y] = NewWater(w.x+1, w.y)
		}
	}
}

func (w *Water) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(w.x), float32(w.y), float32(1*pixelSize), float32(1*pixelSize), w.color, false)
}
