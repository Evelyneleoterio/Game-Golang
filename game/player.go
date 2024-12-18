package game

import (
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct{
	image *ebiten.Image
	position Vector
	game *Game
	laserLoadingTime *Time
}

// initializing the image
func NewPlayer(game *Game) *Player{
	image := assets.PlayerSprite

		bounds := image.Bounds()
		halfW := float64(bounds.Dx()) / 2
	position:= Vector{
	X: (screenWidth / 2) - halfW, 
	Y: 500,
}
	return &Player{
		image : image,
		game: game,
		position: position,
		laserLoadingTime: NewTime(12),

	}
}

// responsible for updating the game logic
func (p *Player) Update() {
	speed:= 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	}else if ebiten.IsKeyPressed(ebiten.KeyRight){
		p.position.X += speed
	}
	p.laserLoadingTime.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTime.Ready(){
		p.laserLoadingTime.Reset()
	bounds := p.image.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2.

		spawPos := Vector{

	p.position.X + halfW,
	p.position.Y - halfH/2,
}
		laser := NewLaser(spawPos)
		p.game.AddLasers(laser)
	}
}

//desenha na tela
func (p *Player) Draw(screen *ebiten.Image){
	op := &ebiten.DrawImageOptions{}

	//x and y position of the image that will be drawn on the screen
	op.GeoM.Translate(p.position.X, p.position.Y)

	// draw on screen
	screen.DrawImage(p.image, op)
	
}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()
	return NewRect(p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()))
}