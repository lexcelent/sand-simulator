package base

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/lexcelent/sand-simulator/utils"
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

type Game struct {
	CurrentGrid     [][]Material // сетка текущих состояний
	NextGrid        [][]Material // сетка следующих состояний
	drag            bool
	currentMaterial int // 1 - SAND, 2 - WATER, 3 - STONE
	waterCount      int
	sandCount       int
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		// Выбрать песок
		g.currentMaterial = utils.Sand
	}

	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		// Выбрать воду
		g.currentMaterial = utils.Water
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
	}

	// Обработка зажатой кнопки мыши
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.drag = true
	}

	if g.drag {
		// TODO: не работает currentMaterial
		x, y := ebiten.CursorPosition()
		if g.currentMaterial == utils.Sand {
			g.CurrentGrid[x][y] = NewSand(x, y)
		} else {
			g.CurrentGrid[x][y] = NewWater(x, y)
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.drag = false
	}

	// Здесь обрабатываем движение материалов
	// Проходимся по каждой строке. В каждой строке проходимся по столбцу
	for x := 0; x < gridHeight-1; x++ {
		for y := 0; y < gridWidht-1; y++ {

			// TODO: Здесь должно быть что-то вроде g.CurrentGrid[x][y].Update() и больше ничего
			// Вся логика будет описана в апдейте каждого элемента
			g.CurrentGrid[x][y].Update(g)

			// switch g.currentGrid[x][y].NameID() {
			// case utils.Sand:
			// 	// Обработка песка
			// case utils.Water:
			// 	// Обрабатка воды
			// }
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.waterCount = 0
	g.sandCount = 0

	// отрисую допустимую сетку
	vector.StrokeLine(screen, float32(gridHeight), 0, float32(gridHeight), float32(gridWidht), float32(pixelSize), utils.RED, false)
	// vector.StrokeLine(screen, 0, float32(gridWidht), float32(gridHeight), float32(gridWidht), float32(pixelSize), RED, false)

	// Отрисовка песка
	for x := 0; x < gridHeight; x++ {
		for y := 0; y < gridWidht; y++ {
			switch g.CurrentGrid[x][y].NameID() {
			case utils.Sand:
				// Если здесь есть песок, то нужно его создать и отрисовать
				vector.DrawFilledRect(screen, float32(x), float32(y), float32(1*pixelSize), float32(1*pixelSize), utils.YELLOW, false)
				g.sandCount++
			case utils.Water:
				// Если здесь вода, нужно создать и отрисовать
				vector.DrawFilledRect(screen, float32(x), float32(y), float32(1*pixelSize), float32(1*pixelSize), utils.BLUE, false)
				g.waterCount++
			}
			// Текущее состояние теперь отрисовано. Обновляем текущее состояние
			g.CurrentGrid[x][y] = g.NextGrid[x][y]
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

func (g *Game) Reset() {
	for x := 0; x < gridHeight; x++ {
		for y := 0; y < gridWidht; y++ {
			g.CurrentGrid[x][y] = NewEmpty(x, y)
			g.NextGrid[x][y] = NewEmpty(x, y)
		}
	}
}

func NewWorld(width, height int) *Game {
	// Создать две сетки для запоминания состояний

	// rows
	// Код отвечает на вопрос: сколько строк будет в матрице?
	curGrid := make([][]Material, height)
	nxtGrid := make([][]Material, height)

	// columns
	// Код отвечает на вопрос: сколько столбцов нужно в каждой строке?
	for i := range curGrid {
		curGrid[i] = make([]Material, width)
		nxtGrid[i] = make([]Material, width)
	}

	return &Game{CurrentGrid: curGrid, NextGrid: nxtGrid}
}
