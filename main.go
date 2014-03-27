package main

import sf "bitbucket.org/krepa098/gosfml2"
import "log"
import "runtime"
import "time"

func init() {
	runtime.LockOSThread()
}

func main() {
	game := Game{}
	game.start()
}

type Game struct {
	window *sf.RenderWindow
	sprite *sf.Sprite
}

func (g *Game) start() {
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

	g.sprite, err = sf.NewSprite(texture)
	if err != nil {
		log.Fatal(err)
	}

	g.mainLoop()
}

func (g *Game) mainLoop() {
	ticker := time.NewTicker(time.Second / 60)

	for g.window.IsOpen() {
		select {
		case <-ticker.C:
			for event := g.window.PollEvent(); event != nil; event = g.window.PollEvent() {
				g.onEvent(event)
			}

			g.onTick()
		}
	}
}

func (g *Game) onEvent(event sf.Event) {
		switch ev := event.(type) {
		case sf.EventKeyReleased:
			g.onKeyReleased(ev.Code)
		case sf.EventClosed:
			g.Close()
		}
}

func (g *Game) onTick() {
	// g.onUpdate()
	g.window.Clear(sf.ColorMagenta())
	g.onRender()
	g.window.Display()
}

func (g *Game) onRender() {
	g.window.Draw(g.sprite, sf.DefaultRenderStates())
}

func (g *Game) onKeyReleased(keyCode sf.KeyCode) {
	if keyCode == sf.KeyEscape {
		g.Close()
	}
}

// crashes, why?
func (g *Game) onUpdate() {
	if sf.KeyboardIsKeyPressed(sf.KeyRight) {
		pos := g.sprite.GetPosition()
		x := pos.X + 5
		y := pos.Y
		g.sprite.SetPosition(sf.Vector2f{x,y})
	}
}

func (g *Game) Close() {
	g.window.Close()
}
