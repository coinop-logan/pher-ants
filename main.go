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

func main() {
	pixelgl.Run(func() {
  	cfg := pixelgl.WindowConfig{
  		Title:  "Pixel Rocks!",
  		Bounds: pixel.R(0, 0, 1024, 768),
  	}
  	win, err := pixelgl.NewWindow(cfg)
  	if err != nil {
  		panic(err)
  	}

    generatePictures()

    fps := time.Tick(time.Second / 60)

    for i := 5; i<6; i++ {
      newAnt := new(ant)
      newAnt.pos = pixel.V(50*float64(i),50)
      gameEntities = append(gameEntities, newAnt)
    }

  	for !win.Closed() {
      handleInput(win)

      for i := 0; i<len(gameEntities); i++ {
        gameEntities[i].iterate()
      }

      drawStuff(win)

      <-fps
  	}
  })
}
