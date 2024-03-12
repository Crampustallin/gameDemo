package figures

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

type Character struct {
	X, Y float32
	width, height float32
	Word string
	IsActive bool
	figureColor color.RGBA
}

func NewCharacter(x, y, width, height float32, word string) *Character {
	return &Character {
		X: x,
		Y: y,
		width: width,
		height: height,
		Word: word,
		IsActive: false,
	}
}

func (c *Character) SetPlayerBody(width, height float32) {
	c.width = width
	c.height = height
	c.figureColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
}

func (c *Character) SetPlayerPos(x, y float32) error {
	c.X = x
	c.Y = y
	return nil
}

func (c *Character) SetActive() {
	c.IsActive = true
	c.figureColor = color.RGBA{R: 255, G: 0, B: 0, A:255}
}

func (c *Character) GetPlayerPos() (float32, float32) {
	return c.X, c.Y
}

func (c *Character) GetPlayerBody() (float32, float32) {
	return c.width, c.height
}

func (c *Character) Damage() {
	c.width -= 5
	c.height = c.width
}

func (c *Character) DrawCharacter(screen *ebiten.Image, fontFace font.Face) {
	vector.DrawFilledRect(screen, c.X, c.Y, c.width, c.height, c.figureColor, false)
	text.Draw(screen, c.Word, fontFace, int(c.X), int(c.Y + (c.height * 1.5)), color.RGBA{R:255, G:165, B: 0, A:0})
}
