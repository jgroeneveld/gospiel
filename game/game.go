package game

import (
	"log"
	"time"

	sf "bitbucket.org/krepa098/gosfml2"
)

type Game struct {
	window *sf.RenderWindow
	entity *Entity
}

func (g *Game) Start() {
	g.window = sf.NewRenderWindow(
		sf.VideoMode{800, 600, 32},
		"Hallo Welt",
		sf.StyleDefault,
		sf.DefaultContextSettings(),
	)

	texture, err := sf.NewTextureFromFile("foo.png", nil)
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
	renderTicker := time.NewTicker(time.Second / 120)
	updateTicker := time.NewTicker(time.Second / 60)

	for g.window.IsOpen() {
		select {
		case <-updateTicker.C:
			g.onUpdate()

		case <-renderTicker.C:
			g.onRender()
		}
	}
}

func (g *Game) onUpdate() {
	for event := g.window.PollEvent(); event != nil; event = g.window.PollEvent() {
		g.onEvent(event)
	}

	// crashes, why?
	// if sf.KeyboardIsKeyPressed(sf.KeyRight) {
	pos := g.entity.GetPosition()
	x := pos.X + 2
	y := pos.Y
	g.entity.SetPosition(sf.Vector2f{x, y})
	// }
}

func (g *Game) onRender() {
	g.window.Clear(sf.ColorMagenta())
	g.window.Draw(g.entity, sf.DefaultRenderStates())
	g.window.Display()
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
