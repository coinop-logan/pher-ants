package main

import (
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/imdraw"
  "github.com/faiface/pixel/pixelgl"
)

func drawTileHighlight(win *pixelgl.Window) {
  t := gameMap.getTileAtPos(win.MousePosition())

  imd := imdraw.New(nil)

  imd.Color = pixel.RGBA{0,0,1,0.5}
  r := t.getRect()
  imd.Push(r.Min, r.Max)
  imd.Rectangle(0)

  imd.Draw(win)
}
