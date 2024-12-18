package game

const (
	screenWidth = 800
	screenHeight = 600
)

type Vector struct{
	X float64
	Y float64
}

type Rect struct{
	X float64
	Y float64
	Width float64
	Height float64	
}

// verifica ss um retângulo está em cima do outro com base na altura e largura das imagens e também da posição x e y que ela esta
func NewRect (x, y, width, height float64) Rect {
	return Rect{
		X: x,
		Y: y, 
	Width: width,
	Height: height, 	
	}
	
}

func (r Rect) Intersects(other Rect) bool{
	return r.X <= other.maX() &&
	other.X <= r.maX() &&
	r.Y <= other.maxY() &&
	other.Y <= r.maxY()
}

func (r Rect) maX() float64{
	return r.X + r.Width
}

func (r Rect) maxY() float64{
	return r.Y + r.Height
}