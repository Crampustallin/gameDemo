package figures

type Character struct {
	positionX float64
	positionY float64
	width float64
	height float64
}

func newCharacter(x, y, width, height float64) *Character {
	return &Character {
		positionX: x,
		positionY: y,
		width: width,
		height: height,
	}
}

func (c *Character) setPlayerBody(width, height float64) {
	c.width = width
	c.height = height
}

func (c *Character) setPlayerPos(x, y int) error {
	c.positionX = float64(x)
	c.positionY = float64(y)
	return nil
}
