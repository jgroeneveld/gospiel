package main

import (
	"runtime"

	"github.com/jgroeneveld/game/game"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	game := game.Game{}
	game.Start()
}
