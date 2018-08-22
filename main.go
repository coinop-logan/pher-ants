package main

import (
  //"fmt"
  "time"
  //"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
  gameMap gameMapType
  gameEntities []gameEntity
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

  generatePictures()

  fps := time.Tick(time.Second / 120)

  gameMap.init()

  newAnt := ant{
    pos: pixel.V(50,50),
  }

  gameEntities = append(gameEntities, &newAnt)

	for !win.Closed() {
    handleInput(win)

    for i := 0; i<len(gameEntities); i++ {
      gameEntities[i].iterate()
    }

    drawStuff(win)

    <-fps
	}
}

func main() {
	pixelgl.Run(run)
}
