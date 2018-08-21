package main

import (
  "fmt"
  "time"
  "log"

  //"github.com/golang/freetype/truetype"
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

  fps := time.Tick(time.Second / 120)

  gameMap.init()

  var gameEntities []gameEntity

  newAnt := ant{
    pos: pixel.V(50,50),
  }

  gameEntities = append(gameEntities, &newAnt)

	for !win.Closed() {
    //background
    win.Clear(colornames.Forestgreen)

    //grid
    gameMap.drawGrid(win)

    //entities
    for i := 0; i<len(gameEntities); i++ {
      gameEntities[i].draw(win)
    }

		win.Update()

    <-fps
	}
}

func main() {
  log.Output(1, fmt.Sprintf("%v", gameMap.tiles[0][0].x))
	pixelgl.Run(run)
}