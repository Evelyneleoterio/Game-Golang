package game

import (
	"fmt"
	"image/color"
	"my-game/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	player           *Player
	lasers           []*Laser
	meteors          []*Meteor
	meteorSpawnTimer *Time
	stars            []*Star
	state            string
	score            int
	gameOverTimer    time.Time
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTime(24),
		state:            "playing",
	}
	player := NewPlayer(g)
	g.player = player

	for i := 0; i < 100; i++ {
		g.stars = append(g.stars, NewStar())
	}
	return g
}

// Responsável por atualizar a lógica do jogo
func (g *Game) Update() error {
	if g.state == "gameover" {
		// Verifica se já passou tempo suficiente para resetar
		if time.Since(g.gameOverTimer) > 2*time.Second {
			g.Reset()
			g.state = "playing"
		}
		return nil
	}

	g.player.Update()
	for _, l := range g.lasers {
		l.Update()
	}
	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.Ready() {
		g.meteorSpawnTimer.Reset()
		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, s := range g.stars {
		s.Update()
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.state = "gameover" // Altera o estado do jogo
			g.gameOverTimer = time.Now()
			return nil
		}
	}

	for i, m := range g.meteors {

		for j, l := range g.lasers {
			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score += 1
			}
		}

	}
	return nil
}

// Responsável por desenhar objetos na tela
func (g *Game) Draw(screen *ebiten.Image) {
	if g.state == "gameover" {
		// Mensagem centralizada
		message := "GAME OVER"
		bounds := text.BoundString(assets.FontUi, message)
		textX := (screenWidth - bounds.Dx()) / 2  // Calcula o X central
		textY := (screenHeight - bounds.Dy()) / 2 // Calcula o Y central
		text.Draw(screen, message, assets.FontUi, textX, textY, color.RGBA{255, 0, 0, 255})
		return
	}

	g.player.Draw(screen)

	for _, s := range g.stars {
		s.Draw(screen)
	}

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), assets.FontUi, 20, 100, color.White)

}

// Responsável por atualizar o tamanho da nossa tela
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.meteorSpawnTimer.Reset()
	g.score = 0

}
