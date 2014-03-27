package game

import (
	"log"
	"time"

	sf "bitbucket.org/krepa098/gosfml2"
)

func New() *Game {
	return &Game{}
}

type Game struct {
	window *sf.RenderWindow
	entity *Entity
}

func (g *Game) Start(window *sf.RenderWindow) {
	g.window = window

	texture, err := sf.NewTextureFromFile("res/foo.png", nil)
	if err != nil {
		log.Fatal(err)
	}

	sprite, err := sf.NewSprite(texture)
	if err != nil {
		log.Fatal(err)
	}
	g.entity = &Entity{sprite}

	g.mainLoop()
}

func (g *Game) mainLoop() {
	ticker := time.NewTicker(time.Second / 60)

	for g.window.IsOpen() {
		select {
		case <-ticker.C:
			g.Update()
			g.Draw()
		}
	}
}

func (g *Game) Update() {
	for event := g.window.PollEvent(); event != nil; event = g.window.PollEvent() {
		g.onEvent(event)
	}

	speed, movement := g.getSpeedAndMovementFromControl()

	pos := g.entity.GetPosition()
	pos.X += movement.X * speed.X
	pos.Y += movement.Y * speed.Y
	g.entity.SetPosition(pos)
}

func (g *Game) Draw() {
	g.window.Clear(sf.ColorMagenta())
	g.window.Draw(g.entity, sf.DefaultRenderStates())
	g.window.Display()
}

func (g *Game) getSpeedAndMovementFromControl() (sf.Vector2f, sf.Vector2f) {
	movement := sf.Vector2f{0, 0}
	speed := sf.Vector2f{2, 2}

	if sf.KeyboardIsKeyPressed(sf.KeySpace) {
		speed.X *= 3
		speed.Y *= 3
	}

	if sf.KeyboardIsKeyPressed(sf.KeyRight) {
		movement.X += 1.0
	}

	if sf.KeyboardIsKeyPressed(sf.KeyLeft) {
		movement.X += -1.0
	}

	if sf.KeyboardIsKeyPressed(sf.KeyUp) {
		movement.Y += -1.0
	}

	if sf.KeyboardIsKeyPressed(sf.KeyDown) {
		movement.Y += 1.0
	}

	return speed, movement
}

func (g *Game) onEvent(event sf.Event) {
	switch ev := event.(type) {
	case sf.EventKeyReleased:
		g.onKeyReleased(ev.Code)
	case sf.EventClosed:
		g.Close()
	}
}

func (g *Game) onKeyReleased(keyCode sf.KeyCode) {
	if keyCode == sf.KeyEscape {
		g.Close()
	}
}

func (g *Game) Close() {
	g.window.Close()
}
