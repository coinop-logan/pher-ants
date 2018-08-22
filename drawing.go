package main

import (
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
  "github.com/faiface/pixel/imdraw"
  "golang.org/x/image/colornames"
)

var (
  selectedTileCanvas *pixelgl.Canvas
  highlightedTileCanvas *pixelgl.Canvas
  antCanvas *pixelgl.Canvas
)

func generatePictures() {

  imd := imdraw.New(nil)

  //highlightedTile

  tileRect := pixel.R(0,0,TILE_WIDTH,TILE_WIDTH)

  highlightedTileCanvas = pixelgl.NewCanvas(tileRect)

  imd.Color = pixel.RGBA{0,1,0,0.5}
  imd.Push(tileRect.Min, tileRect.Max)
  imd.Rectangle(0)

  imd.Draw(highlightedTileCanvas)

  //selectedTile

  imd.Clear()

  selectedTileCanvas = pixelgl.NewCanvas(tileRect)

  imd.Color = pixel.RGBA{0,0,1,0.5}
  imd.Push(tileRect.Min, tileRect.Max)
  imd.Rectangle(0)

  imd.Draw(selectedTileCanvas)

  //ant

  imd.Clear()

  antCanvas = pixelgl.NewCanvas(pixel.R(0,0,16,16))

  imd.Color = pixel.RGB(0,0,1)
  imd.Push(pixel.V(1,    0   ).Scaled(10).Add(pixel.V(8,8)),
           pixel.V(-1,   -0.5).Scaled(10).Add(pixel.V(8,8)),
           pixel.V(-0.5, 0   ).Scaled(10).Add(pixel.V(8,8)),
           pixel.V(-1,   0.5 ).Scaled(10).Add(pixel.V(8,8)))
  imd.Polygon(0)

  imd.Draw(antCanvas)
}

type drawable interface {
  draw(t pixel.Target)
}

func drawStuff(win *pixelgl.Window) {
  //background
  win.Clear(colornames.Forestgreen)

  //gameMap.drawGrid(win)

  //entities
  for i := 0; i<len(gameEntities); i++ {
    gameEntities[i].draw(win)
  }

  drawHighlightedTile(win)
  drawSelectedTiles(win)

  win.Update()
}

func (ant *ant) draw(t pixel.Target) {
  antCanvas.Draw(t, pixel.IM.Rotated(pixel.ZV, ant.angle).Moved(ant.pos))
}

func drawHighlightedTile(t pixel.Target) {
  ti := highlightedTile
  highlightedTileCanvas.Draw(t, pixel.IM.Moved(ti.centerPos()))
}

func drawSelectedTiles(t pixel.Target) {
  for _, ti := range(selectedTiles) {
    selectedTileCanvas.Draw(t, pixel.IM.Moved(ti.centerPos()))
  }
}
