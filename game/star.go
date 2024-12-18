package game

import (
	"math/rand"
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Star struct{
	image *ebiten.Image
	speed float64
	position Vector
}

func NewStar() *Star {
	image := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))]
	speed := 1.0 + rand.Float64()*2.0
	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: rand.Float64() * screenHeight,
	}
	return &Star{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (s *Star) Update(){
	s.position.Y += s.speed

	if s.position.Y > screenHeight {
        s.position.Y = 0
        s.position.X = rand.Float64() * screenWidth 
        s.speed = 1.0 + rand.Float64()*2.0
    }
}

func (s *Star) Draw(screen *ebiten.Image){
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.position.X, s.position.Y)

	// draw on screen
	screen.DrawImage(s.image, op)
}

func (s *Star) Collider() Rect {
	bounds := s.image.Bounds()
	return NewRect(s.position.X,
		s.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()))
}