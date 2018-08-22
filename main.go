package main

import (
  //"fmt"
  "time"
  //"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
  //"github.com/faiface/pixel/text"
  //"github.com/faiface/pixel/imdraw"
  //"golang.org/x/image/font/basicfont"
  //"golang.org/x/image/font"
  "golang.org/x/image/colornames"
)

var gameMap gameMapType

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

  var gameEntities []gameEntity

  newAnt := ant{
    pos: pixel.V(50,50),
  }

  gameEntities = append(gameEntities, &newAnt)

	for !win.Closed() {
    handleInput(win)

    for i := 0; i<len(gameEntities); i++ {
      gameEntities[i].iterate()
    }

    //background
    win.Clear(colornames.Forestgreen)

    //gameMap.drawGrid(win)

    //entities
    for i := 0; i<len(gameEntities); i++ {
      gameEntities[i].draw(win)
    }

    drawTileHighlight(win)

		win.Update()

    <-fps
	}
}

func main() {
	pixelgl.Run(run)
}
