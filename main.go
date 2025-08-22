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

// func (g *Game) UpdateSand(x, y int) {
// 	if g.currentGrid[x][y+1] == Empty {
// 		// Если под пикселем пусто, то перемещаем пиксель вниз
// 		// Текущую сетку не меняем!!!
// 		g.nextGrid[x][y] = Empty
// 		g.nextGrid[x][y+1] = Sand
// 	} else if g.currentGrid[x-1][y+1] == Empty && g.currentGrid[x+1][y+1] == Empty {
// 		// Если по бокам пусто, то выбираем направление наугад
// 		direction := rand.Int()
// 		if direction%2 == 0 {
// 			// Направляем песок влево
// 			g.nextGrid[x][y] = Empty
// 			g.nextGrid[x-1][y+1] = Sand
// 		} else {
// 			// Направляем песок вправо
// 			g.nextGrid[x][y] = Empty
// 			g.nextGrid[x+1][y+1] = Sand
// 		}
// 	} else if g.currentGrid[x-1][y+1] == Empty {
// 		// Если пусто только слева
// 		g.nextGrid[x][y] = Empty
// 		g.nextGrid[x-1][y+1] = Sand
// 	} else if g.currentGrid[x+1][y+1] == Empty {
// 		// Если пусто только справа
// 		g.nextGrid[x][y] = Empty
// 		g.nextGrid[x+1][y+1] = Sand

// 	} else if g.currentGrid[x][y+1] == Water {
// 		// Песок проваливается в воду
// 		g.nextGrid[x][y] = Water
// 		g.nextGrid[x][y+1] = Sand
// 	}
// }

// func (g *Game) UpdateWater(x, y int) {
// 	// Обновить пиксель воды. Сделал тут дополнительную проверку для nextGrid, чтобы пиксели не исчезали.
// 	// Наверняка есть другое решение
// 	if g.currentGrid[x][y+1] == Empty && g.nextGrid[x][y+1] == Empty {
// 		// Если под пикселем пусто, то перемещаем пиксель вниз
// 		g.nextGrid[x][y] = Empty
// 		g.nextGrid[x][y+1] = Water
// 	} else if g.currentGrid[x-1][y+1] == Empty {
// 		// Если пусто только слева
// 		if g.nextGrid[x-1][y+1] == Empty {
// 			g.nextGrid[x][y] = Empty
// 			g.nextGrid[x-1][y+1] = Water
// 		}
// 	} else if g.currentGrid[x+1][y+1] == Empty {
// 		// Если пусто только справа
// 		if g.nextGrid[x+1][y+1] == Empty {
// 			g.nextGrid[x][y] = Empty
// 			g.nextGrid[x+1][y+1] = Water
// 		}
// 	} else if g.currentGrid[x-1][y] == Empty && g.currentGrid[x+1][y] == Empty {
// 		// Проверяем слева и справа
// 		direction := rand.Int()
// 		if direction%2 == 0 && g.nextGrid[x-1][y] == Empty {
// 			// Направляем воду влево
// 			g.nextGrid[x][y] = Empty
// 			g.nextGrid[x-1][y] = Water
// 		} else if direction%2 != 0 && g.currentGrid[x+1][y] == Empty {
// 			// Направляем воду вправо
// 			g.nextGrid[x][y] = Empty
// 			g.nextGrid[x+1][y] = Water
// 		}
// 	} else if g.currentGrid[x-1][y] == Empty {
// 		if g.nextGrid[x-1][y] == Empty {
// 			g.nextGrid[x][y] = Empty
// 			g.nextGrid[x-1][y] = Water
// 		}
// 	} else if g.currentGrid[x+1][y] == Empty {
// 		if g.nextGrid[x+1][y] == Empty {
// 			g.nextGrid[x][y] = Empty
// 			g.nextGrid[x+1][y] = Water
// 		}
// 	}
// }
