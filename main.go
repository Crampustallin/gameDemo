package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/Crampustallin/gameDemo/figures"
)

type Game struct{
	player *figures.Character
}

func NewGame() *Game {
	character := figures.NewCharacter(float64(0),float64(0),float64(10.5), float64(10.5))
	return &Game{
		player: character,
	}
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	g.player.SetPlayerPos(x,y)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	playerPosX, playerPosY := g.player.GetPlayerPos()
	playerWidth, playerHeight := g.player.GetPlayerBody()
	ebitenutil.DrawRect(screen, playerPosX, playerPosY,  playerWidth, playerHeight, color.White)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f", ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Word Hunter")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
