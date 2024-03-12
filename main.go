package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"golang.org/x/image/font"

	"github.com/yohamta/donburi"

	"github.com/Crampustallin/gameDemo/assets"
	"github.com/Crampustallin/gameDemo/figures"
)

type Game struct {
	world donburi.World
	rect *donburi.ComponentType[figures.Character]
	activeEnemy *figures.Character
	fontFace font.Face
	key []ebiten.Key
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
		spawnX := getRandomPosition(minX, float32(maxX))
		spawnY := getRandomPosition(minY, float32(maxY))
		r.SetPlayerPos(spawnX, spawnY)
		r.Word = words[rnd.Intn(len(words))]
		r.SetPlayerBody(float32(10 * len(r.Word)), float32(10 * len(r.Word)))
	})

	var activeEnemy *figures.Character;

	if entry, ok := rect.First(world); ok {
		r := rect.Get(entry)
		r.SetActive()
		activeEnemy = r
	}
	
	fontFace := assets.LoadFont()

	return &Game{
		world: world, 
		rect: rect,
		fontFace: fontFace,
		activeEnemy: activeEnemy,
	}
}


func (g *Game) Update() error {
	g.key = inpututil.AppendPressedKeys(g.key[:0])
	for _, key := range g.key {
		if after, found := strings.CutPrefix(g.activeEnemy.Word, strings.ToLower(key.String())); found { // TODO: find a better way to check if key is right
			g.activeEnemy.Word = after
			if g.activeEnemy.Word == "" {
				entry, _ := g.rect.First(g.world)
				entry.Remove()
				if entry, has := g.rect.First(g.world); has {
					g.activeEnemy = g.rect.Get(entry)
					g.activeEnemy.SetActive()
				}
			} else {
				g.activeEnemy.Damage() // TODO: add animation when a character is taking damage
			}
		}
	}
	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
	g.rect.Each(g.world, func(entry *donburi.Entry) {
		r := g.rect.Get(entry)
		r.DrawCharacter(screen, g.fontFace)
	})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f \nTarget word: %s", ebiten.ActualFPS(), g.activeEnemy.Word))
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
