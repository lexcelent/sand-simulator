package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/lexcelent/sand-simulator/resource"
)

/*
TODO: Проверить левые и правые границы экрана (иначе index out of range)
TODO: Вместо сетки состояний создать объекты (напр. SAND), которые проверяют коллизии под собой (или с другими объектами)

Later:
TODO: Осуществить работу с увеличенными пикселями (иначе пиксели накладываются друг на друга)
Задача отложена, т.к. для этого возможно нужен полностью другой код
*/

func main() {
	// g := &Game{}
	g := NewWorld(gridWidht, gridHeight)

	img, _, err := image.Decode(bytes.NewReader(resource.Cloud_img))
	if err != nil {
		log.Fatal(err)
	}

	// Облако
	cloudImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sand Simulator")
	ebiten.SetTPS(60)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

// Облако
var (
	cloudImage *ebiten.Image
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

const (
	Empty = iota
	Sand
	Water
	Stone
	Cloud
	Snow
)

var (
	RED    = color.RGBA{255, 0, 0, 0}
	YELLOW = color.RGBA{255, 255, 0, 0}
	BLUE   = color.RGBA{0, 191, 255, 0}
	GRAY   = color.RGBA{128, 128, 128, 0}
	WHITE  = color.RGBA{255, 255, 255, 0}
)

type Game struct {
	currentGrid     [][]int // сетка текущих состояний
	nextGrid        [][]int // сетка следующих состояний
	drag            bool
	currentMaterial int // 1 - SAND, 2 - WATER, 3 - STONE
	waterCount      int
	sandCount       int
}

func NewWorld(width, height int) *Game {
	// Создать две сетки для запоминания состояний

	// rows
	// Код отвечает на вопрос: сколько строк будет в матрице?
	curGrid := make([][]int, height)
	nxtGrid := make([][]int, height)

	// columns
	// Код отвечает на вопрос: сколько столбцов нужно в каждой строке?
	for i := range curGrid {
		curGrid[i] = make([]int, width)
		nxtGrid[i] = make([]int, width)
	}

	return &Game{currentGrid: curGrid, nextGrid: nxtGrid}
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		// Выбрать песок
		g.currentMaterial = Sand
	}

	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		// Выбрать воду
		g.currentMaterial = Water
	}

	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		// Выбрать камень
		g.currentMaterial = Stone
	}

	if inpututil.IsKeyJustPressed(ebiten.Key4) {
		// Выбрать тучу
		g.currentMaterial = Cloud
	}

	if inpututil.IsKeyJustPressed(ebiten.Key5) {
		// Снег
		g.currentMaterial = Snow
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
	}

	// Обработка зажатой кнопки мыши
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.drag = true
	}

	if g.drag {
		x, y := ebiten.CursorPosition()
		g.currentGrid[x][y] = g.currentMaterial
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.drag = false
	}

	// Здесь обрабатываем движение материалов
	// Проходимся по каждой строке. В каждой строке проходимся по столбцу
	for x := 0; x < gridHeight-1; x++ {
		for y := 0; y < gridWidht-1; y++ {
			if g.currentGrid[x][y] == Sand {
				// Обработка песка
				g.UpdateSand(x, y)
			}

			if g.currentGrid[x][y] == Water {
				// Обрабатка воды
				g.UpdateWater(x, y)
			}

			if g.currentGrid[x][y] == Stone {
				g.nextGrid[x][y] = Stone
			}

			if g.currentGrid[x][y] == Cloud {
				g.nextGrid[x][y] = Cloud
			}

			if g.currentGrid[x][y] == Snow {
				g.UpdateSnow(x, y)
			}
		}
	}

	return nil
}

func (g *Game) Reset() {
	for x := 0; x < gridHeight; x++ {
		for y := 0; y < gridWidht; y++ {
			g.currentGrid[x][y] = 0
			g.nextGrid[x][y] = 0
		}
	}
}

func (g *Game) UpdateSand(x, y int) {
	if g.currentGrid[x][y+1] == Empty {
		// Если под пикселем пусто, то перемещаем пиксель вниз
		// Текущую сетку не меняем!!!
		g.nextGrid[x][y] = Empty
		g.nextGrid[x][y+1] = Sand
	} else if g.currentGrid[x-1][y+1] == Empty && g.currentGrid[x+1][y+1] == Empty {
		// Если по бокам пусто, то выбираем направление наугад
		direction := rand.Int()
		if direction%2 == 0 {
			// Направляем песок влево
			g.nextGrid[x][y] = Empty
			g.nextGrid[x-1][y+1] = Sand
		} else {
			// Направляем песок вправо
			g.nextGrid[x][y] = Empty
			g.nextGrid[x+1][y+1] = Sand
		}
	} else if g.currentGrid[x-1][y+1] == Empty {
		// Если пусто только слева
		g.nextGrid[x][y] = Empty
		g.nextGrid[x-1][y+1] = Sand
	} else if g.currentGrid[x+1][y+1] == Empty {
		// Если пусто только справа
		g.nextGrid[x][y] = Empty
		g.nextGrid[x+1][y+1] = Sand

	} else if g.currentGrid[x][y+1] == Water {
		// Песок проваливается в воду
		g.nextGrid[x][y] = Water
		g.nextGrid[x][y+1] = Sand
	}
}

func (g *Game) UpdateSnow(x, y int) {
	if g.currentGrid[x][y+1] == Empty {
		// Если под пикселем пусто, то перемещаем пиксель вниз
		// Текущую сетку не меняем!!!
		g.nextGrid[x][y] = Empty
		g.nextGrid[x][y+1] = Snow
	} else if g.currentGrid[x-1][y+1] == Empty && g.currentGrid[x+1][y+1] == Empty {
		// Если по бокам пусто, то выбираем направление наугад
		direction := rand.Int()
		if direction%2 == 0 {
			// Направляем песок влево
			g.nextGrid[x][y] = Empty
			g.nextGrid[x-1][y+1] = Snow
		} else {
			// Направляем песок вправо
			g.nextGrid[x][y] = Empty
			g.nextGrid[x+1][y+1] = Snow
		}
	} else if g.currentGrid[x-1][y+1] == Empty {
		// Если пусто только слева
		g.nextGrid[x][y] = Empty
		g.nextGrid[x-1][y+1] = Snow
	} else if g.currentGrid[x+1][y+1] == Empty {
		// Если пусто только справа
		g.nextGrid[x][y] = Empty
		g.nextGrid[x+1][y+1] = Snow

	} else if g.currentGrid[x][y+1] == Water {
		// Песок проваливается в воду
		g.nextGrid[x][y] = Water
		g.nextGrid[x][y+1] = Snow
	}
}

func (g *Game) UpdateWater(x, y int) {
	// Обновить пиксель воды. Сделал тут дополнительную проверку для nextGrid, чтобы пиксели не исчезали.
	// Наверняка есть другое решение
	if g.currentGrid[x][y+1] == Empty && g.nextGrid[x][y+1] == Empty {
		// Если под пикселем пусто, то перемещаем пиксель вниз
		g.nextGrid[x][y] = Empty
		g.nextGrid[x][y+1] = Water
	} else if g.currentGrid[x-1][y+1] == Empty {
		// Если пусто только слева
		if g.nextGrid[x-1][y+1] == Empty {
			g.nextGrid[x][y] = Empty
			g.nextGrid[x-1][y+1] = Water
		}
	} else if g.currentGrid[x+1][y+1] == Empty {
		// Если пусто только справа
		if g.nextGrid[x+1][y+1] == Empty {
			g.nextGrid[x][y] = Empty
			g.nextGrid[x+1][y+1] = Water
		}
	} else if g.currentGrid[x-1][y] == Empty && g.currentGrid[x+1][y] == Empty {
		// Проверяем слева и справа
		direction := rand.Int()
		if direction%2 == 0 && g.nextGrid[x-1][y] == Empty {
			// Направляем воду влево
			g.nextGrid[x][y] = Empty
			g.nextGrid[x-1][y] = Water
		} else if direction%2 != 0 && g.currentGrid[x+1][y] == Empty {
			// Направляем воду вправо
			g.nextGrid[x][y] = Empty
			g.nextGrid[x+1][y] = Water
		}
	} else if g.currentGrid[x-1][y] == Empty {
		if g.nextGrid[x-1][y] == Empty {
			g.nextGrid[x][y] = Empty
			g.nextGrid[x-1][y] = Water
		}
	} else if g.currentGrid[x+1][y] == Empty {
		if g.nextGrid[x+1][y] == Empty {
			g.nextGrid[x][y] = Empty
			g.nextGrid[x+1][y] = Water
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.waterCount = 0
	g.sandCount = 0

	// отрисую допустимую сетку
	vector.StrokeLine(screen, float32(gridHeight), 0, float32(gridHeight), float32(gridWidht), float32(pixelSize), RED, false)
	// vector.StrokeLine(screen, 0, float32(gridWidht), float32(gridHeight), float32(gridWidht), float32(pixelSize), RED, false)

	// Отрисовка песка
	for x := 0; x < gridHeight; x++ {
		for y := 0; y < gridWidht; y++ {

			if g.currentGrid[x][y] == Sand {
				// Если здесь есть песок, то нужно его создать и отрисовать
				vector.DrawFilledRect(screen, float32(x), float32(y), float32(1*pixelSize), float32(1*pixelSize), YELLOW, false)
				g.sandCount++
			} else if g.currentGrid[x][y] == Water {
				// Если здесь вода, нужно создать и отрисовать
				vector.DrawFilledRect(screen, float32(x), float32(y), float32(1*pixelSize), float32(1*pixelSize), BLUE, false)
				g.waterCount++
			} else if g.currentGrid[x][y] == Stone {
				vector.DrawFilledRect(screen, float32(x), float32(y), float32(1*pixelSize), float32(1*pixelSize), GRAY, false)
			} else if g.currentGrid[x][y] == Cloud {
				// Рисуем облако (загружаем картинку)
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x), float64(y))
				screen.DrawImage(cloudImage, op)
				g.nextGrid[x+15][y+10] = Water
			} else if g.currentGrid[x][y] == Snow {
				vector.DrawFilledRect(screen, float32(x), float32(y), float32(1*pixelSize), float32(1*pixelSize), WHITE, false)
			}

			// Текущее состояние теперь отрисовано. Обновляем текущее состояние
			g.currentGrid[x][y] = g.nextGrid[x][y]
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Current material: %d", g.currentMaterial))

	x, y := ebiten.CursorPosition()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Current pos x: %d y: %d", x, y), 0, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Water count: %d", g.waterCount), 0, 20)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Sand count: %d", g.sandCount), 0, 30)
}

func (g *Game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}
