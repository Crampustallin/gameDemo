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
}

func (c *Character) SetPlayerPos(x, y float32) error {
	c.X = x
	c.Y = y
	return nil
}

func (c *Character) GetPlayerPos() (float32, float32) {
	return c.X, c.Y
}

func (c *Character) GetPlayerBody() (float32, float32) {
	return c.width, c.height
}

func (c *Character) DrawCharacter(screen *ebiten.Image, fontFace font.Face) {
	cl := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	if c.IsActive {
		cl.A, cl.G, cl.B = 0,0,0
	}
	vector.DrawFilledRect(screen, c.X, c.Y, c.width, c.height, cl, false)
	text.Draw(screen, c.Word, fontFace, int(c.X - (c.width/2)), int(c.Y), color.White)
}
