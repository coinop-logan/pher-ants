package main

import (
  //"log"
  //"fmt"
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
  "github.com/faiface/pixel/imdraw"
  "golang.org/x/image/colornames"
)

func init() {
  generatePictures()
}

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

  drawHighlightedTiles(win)
  drawSelectedTiles(win)

  win.Update()
}

func (ant *ant) draw(t pixel.Target) {
  antCanvas.Draw(t, pixel.IM.Rotated(pixel.ZV, ant.angle).Moved(ant.pos))
}

func drawHighlightedTiles(t pixel.Target) {
  minCol, maxCol, minRow, maxRow := gameMap.getBoundsFromTiles(tileHighlightStart, tileHighlightEnd)

  v1 := gameMap.tiles[minCol][minRow].originPos()
  v2 := gameMap.tiles[maxCol+1][maxRow+1].originPos()

  imd := imdraw.New(nil)

  imd.Color = pixel.RGB(0,1,0)
  imd.Push(v1, v2)
  imd.Rectangle(1)

  imd.Draw(t)
}

func drawSelectedTiles(t pixel.Target) {
  for _, ti := range(selectedTiles) {
    selectedTileCanvas.Draw(t, pixel.IM.Moved(ti.centerPos()))
  }
}

func (m *gameMapType) drawGrid(t pixel.Target) {
  for i := 0; i<MAP_WIDTH_IN_TILES; i++ {
    fromPoint1 := m.tiles[0][i].originPos()
    toPoint1 := fromPoint1.Add(pixel.V(TILE_WIDTH*MAP_WIDTH_IN_TILES, 0))

    fromPoint2 := m.tiles[i][0].originPos()
    toPoint2 := fromPoint2.Add(pixel.V(0, TILE_WIDTH*MAP_WIDTH_IN_TILES))

    imd := imdraw.New(nil)

    imd.Color = pixel.RGB(0.5, 0.5, 0.5)
    imd.Push(fromPoint1, toPoint1)
    imd.Line(1)
    imd.Push(fromPoint2, toPoint2)
    imd.Line(1)

    imd.Draw(t)
  }
}
