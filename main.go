package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"

	"github.com/yohamta/donburi"

	"github.com/Crampustallin/gameDemo/assets"
	"github.com/Crampustallin/gameDemo/figures"
)

type Game struct {
	world donburi.World
	rect *donburi.ComponentType[figures.Character]
	fontFace font.Face
}

func getRandomPosition(min, max float32) float32 {
	return float32(min + rand.Float32()*(max-min))
}

func NewGame() *Game {
	words := []string{ "board", "go", "guts", "proffessor", "despair" }
	maxX, maxY := 270, 190 // TODO: set good values for new dots
	var minX, minY float32 = .0, .0
	world := donburi.NewWorld()
	rect := donburi.NewComponentType[figures.Character]()
	world.CreateMany(len(words), rect)
	rnd  := rand.New(rand.NewSource(time.Now().UnixNano()))

	rect.Each(world, func(entry *donburi.Entry) {
		r := rect.Get(entry)
		r.SetPlayerBody(float32(25), float32(25))
		spawnX := getRandomPosition(minX, float32(maxX))
		spawnY := getRandomPosition(minY, float32(maxY))
		r.SetPlayerPos(spawnX, spawnY)
		r.Word = words[rnd.Intn(len(words))]
	})

	if entry, ok := rect.First(world); ok {
		r := rect.Get(entry)
		r.IsActive = true
	}
	
	fontFace := assets.LoadFont()

	return &Game{
		world: world, 
		rect: rect,
		fontFace: fontFace,
	}
}


func (g *Game) Update() error {
	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
	g.rect.Each(g.world, func(entry *donburi.Entry) {
		r := g.rect.Get(entry)
		r.DrawCharacter(screen, g.fontFace)
	})
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
