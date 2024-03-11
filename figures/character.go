package figures

type Character struct {
	positionX float64
	positionY float64
	width float64
	height float64
}

func NewCharacter(x, y, width, height float64) *Character {
	return &Character {
		positionX: x,
		positionY: y,
		width: width,
		height: height,
	}
}

func (c *Character) SetPlayerBody(width, height float64) {
	c.width = width
	c.height = height
}

func (c *Character) SetPlayerPos(x, y int) error {
	c.positionX = float64(x)
	c.positionY = float64(y)
	return nil
}

func (c *Character) GetPlayerPos() (float64, float64) {
	return c.positionX, c.positionY
}
func (c *Character) GetPlayerBody() (float64, float64) {
	return c.width, c.height
}
