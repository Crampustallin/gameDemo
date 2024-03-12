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
	world := donburi.NewWorld()
	rect := donburi.NewComponentType[figures.Character]()
	world.CreateMany(4, rect)
	rect.Each(world, func(entry *donburi.Entry) {
		rect.Set(entry, SpawnEnemy())
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

func SpawnEnemy() *figures.Character {
	words := []string{ "board", "go", "guts", "proffessor", "despair", "deppression" }
	rnd  := rand.New(rand.NewSource(time.Now().UnixNano()))
	maxX, maxY := 270, 190 / 2 // TODO: set good values for new dots
	var minX, minY float32 = .0, .0
	spawnX := getRandomPosition(minX, float32(maxX))
	spawnY := getRandomPosition(minY, float32(maxY))
	Word := words[rnd.Intn(len(words))]
	width, height := float32(5 * len(Word)), float32(5 * len(Word)) // TODO: need to do something with the figures
	enemy := figures.Character{
		X: spawnX,
		Y: spawnY,
		Word: Word,
	}
	enemy.SetPlayerBody(width, height)
	return &enemy
}


func (g *Game) Update() error {
	g.key = inpututil.AppendJustPressedKeys(g.key[:0])
	for _, key := range g.key {
		if after, found := strings.CutPrefix(g.activeEnemy.Word, strings.ToLower(key.String())); found { // TODO: find a better way to check if key is right
			g.activeEnemy.Word = after
			if g.activeEnemy.Word == "" {
				entry, _ := g.rect.First(g.world)
				entry.Remove()
				entity := g.world.Create(g.rect)
				entry = g.world.Entry(entity)
				g.rect.Set(entry, SpawnEnemy())
				if newTarget, has := g.rect.First(g.world); has {
					g.activeEnemy = g.rect.Get(newTarget)
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
