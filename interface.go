package main

import (
  "fmt"
  "log"
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/imdraw"
  "github.com/faiface/pixel/pixelgl"
)

var selectedTiles []*tile

var cmd command

func handleInput(win *pixelgl.Window) {

  if win.JustPressed(pixelgl.MouseButton1) {
    t := gameMap.getTileAtPos(win.MousePosition())
    selectedTiles = append(selectedTiles, t)
  } else if win.JustPressed(pixelgl.MouseButton2) {
    if cmd != nil {
      cmd.rightClick(win.MousePosition())
      for i := 0; i < len(selectedTiles); i++ {
        selectedTiles[i].pher = pheremone{command:cmd}
      }
    }
  }

  if win.JustPressed(pixelgl.KeyEscape) {
    if selectedTiles != nil {
      selectedTiles = nil
    } else {
      cmd = nil
      log.Output(1, fmt.Sprintf("%v", cmd))
    }
  }
  if win.JustPressed(pixelgl.KeyQ) {
    cmd = new(moveCommand)
  }
}

func drawTileHighlight(win *pixelgl.Window) {
  t := gameMap.getTileAtPos(win.MousePosition())

  imd := imdraw.New(nil)

  imd.Color = pixel.RGBA{0,0,1,0.5}
  r := t.getRect()
  imd.Push(r.Min, r.Max)
  imd.Rectangle(0)

  imd.Draw(win)
}
