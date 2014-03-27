package main

import (
	"runtime"

	sf "bitbucket.org/krepa098/gosfml2"

	"github.com/jgroeneveld/gospiel/game"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	window := sf.NewRenderWindow(
		sf.VideoMode{800, 600, 32},
		"Hallo Welt",
		sf.StyleDefault,
		sf.DefaultContextSettings(),
	)

	game := game.New()
	game.Start(window)
}
