package main

import sf "bitbucket.org/krepa098/gosfml2"
import "time"

func main() {
	game := Game{}
	game.start()
}

type Game struct {
	window *sf.RenderWindow
}

func (g *Game) start() {
	ticker := time.NewTicker(time.Second / 60)

	g.window = sf.NewRenderWindow(
		sf.VideoMode{200, 100, 32},
		"Hallo Welt",
		sf.StyleDefault,
		sf.DefaultContextSettings(),
	)


	for g.window.IsOpen() {
		select {
		case <-ticker.C:
			g.onTick()
		}
	}
}

func (g *Game) onTick() {
	for event := g.window.PollEvent(); event != nil; event = g.window.PollEvent() {
		switch ev := event.(type) {
		case sf.EventKeyReleased:
			g.onKeyReleased(ev.Code)
		}
	}

	g.onRender()
}

func (g *Game) onRender() {
	g.window.Clear(sf.ColorMagenta())
	g.window.Display()
}

func (g *Game) onKeyReleased(keyCode sf.KeyCode) {
	if keyCode == sf.KeyEscape {
		g.window.Close()
	}
}
