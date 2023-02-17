package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 160
	screenHeight = 120
	TPS          = 10
)

type Game struct {
	pixels []byte
	state  [][]bool
}

func main() {
	g := newGame()

	ebiten.SetWindowSize(screenWidth*8, screenHeight*8)
	ebiten.SetWindowTitle("Conway's Game of Life")
	ebiten.SetTPS(TPS)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func newGame() *Game {
	g := &Game{
		pixels: make([]byte, screenWidth*screenHeight*4),
		state:  make([][]bool, screenWidth),
	}

	for i := range g.state {
		g.state[i] = make([]bool, screenHeight)
	}

	g.prepopulate()

	return g
}

func (g *Game) prepopulate() {
	population := screenWidth * screenHeight / 10

	for i := 0; i <= population; i++ {
		x := rand.Intn(screenWidth - 1)
		y := rand.Intn(screenHeight - 1)
		g.state[x][y] = true
	}
}

func (g *Game) calculateState() {
	nextGeneration := make([][]bool, screenWidth)

	for x := 0; x < screenWidth; x++ {
		nextGeneration[x] = make([]bool, screenHeight)

		for y := 0; y < screenHeight; y++ {
			neighbours := g.neighboursCount(x, y)

			switch {
			case (neighbours == 2 || neighbours == 3) && g.state[x][y]:
				nextGeneration[x][y] = true
			case neighbours < 2:
				nextGeneration[x][y] = false
			case neighbours > 3:
				nextGeneration[x][y] = false
			case neighbours == 3:
				nextGeneration[x][y] = true
			}
		}
	}

	g.state = nextGeneration
}

func (g *Game) neighboursCount(x int, y int) int {
	count := 0

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {

			if dx == 0 && dy == 0 {
				continue
			}

			nx := x + dx
			ny := y + dy

			if nx < 0 {
				nx = screenWidth - 1
			}

			if ny < 0 {
				ny = screenHeight - 1
			}

			if nx >= screenWidth {
				nx = 0
			}

			if ny >= screenHeight {
				ny = 0
			}

			if g.state[nx][ny] {
				count++
			}
		}
	}

	return count
}

func (g *Game) renderState() {
	for x := range g.state {
		for y := range g.state[x] {

			if g.state[x][y] {
				g.drawCell(x, y)
			} else {
				g.killCell(x, y)
			}
		}
	}
}

func (g *Game) drawCell(x int, y int) {
	i := y*screenWidth + x

	g.pixels[4*i] = 0xff
	g.pixels[4*i+1] = 0xff
	g.pixels[4*i+2] = 0xff
	g.pixels[4*i+3] = 0xff
}

func (g *Game) killCell(x int, y int) {
	i := y*screenWidth + x

	g.pixels[4*i] = 0
	g.pixels[4*i+1] = 0
	g.pixels[4*i+2] = 0
	g.pixels[4*i+3] = 0
}

func (g *Game) Update() error {
	g.calculateState()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderState()
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
