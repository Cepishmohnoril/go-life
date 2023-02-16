package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 160
	screenHeight = 120
)

type Game struct {
	pixels []byte
	state  [][]bool
}

func main() {
	g := newGame()

	ebiten.SetWindowSize(screenWidth*8, screenHeight*8)
	ebiten.SetWindowTitle("Conway's Game of Life")
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

	//mb prepopulate???

	return g
}

func (g *Game) calculateState() {
	for x := range g.state {
		for y := range g.state[x] {

			if !g.state[x][y] {
				g.state[x][y] = true
				return
			}

		}
	}
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
