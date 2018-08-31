package main

import (
  "log"
  //"fmt"
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
  "github.com/faiface/pixel/imdraw"
  "golang.org/x/image/colornames"
)

var (
  selectedTileCanvas *pixelgl.Canvas
  highlightedTileCanvas *pixelgl.Canvas
  antCanvas *pixelgl.Canvas
  moveCommandCanvas *pixelgl.Canvas
  pheremoneOverlayCanvas *pixelgl.Canvas
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

  //pheremoneOverlay

  imd.Clear()

  pheremoneOverlayCanvas = pixelgl.NewCanvas(tileRect)

  imd.Color = pixel.RGBA{1,1,0,0.5}
  imd.Push(tileRect.Min, tileRect.Max)
  imd.Rectangle(0)

  imd.Draw(pheremoneOverlayCanvas)

  //ant

  imd.Clear()

  antCanvas = pixelgl.NewCanvas(pixel.R(0,0,16,16))

  imd.Color = pixel.RGB(1,0.5,0.5)
  imd.Push(pixel.V(1,    0   ).Scaled(10).Add(pixel.V(8,8)),
           pixel.V(-1,   -0.5).Scaled(10).Add(pixel.V(8,8)),
           pixel.V(-0.5, 0   ).Scaled(10).Add(pixel.V(8,8)),
           pixel.V(-1,   0.5 ).Scaled(10).Add(pixel.V(8,8)))
  imd.Polygon(0)

  imd.Draw(antCanvas)

  //move Command

  imd.Clear()

  moveCommandCanvas = pixelgl.NewCanvas(pixel.R(0,0,TILE_WIDTH,TILE_WIDTH))

  imd.Color = pixel.RGB(0,0,1)
  imd.Push(pixel.V(-0.4,  -0.15).Add(pixel.V(0.5,0.5)).Scaled(TILE_WIDTH))
  imd.Push(pixel.V(0.4,   0   ).Add(pixel.V(0.5,0.5)).Scaled(TILE_WIDTH))
  imd.Push(pixel.V(-0.4,  0.15 ).Add(pixel.V(0.5,0.5)).Scaled(TILE_WIDTH))
  imd.Line(1)

  imd.Draw(moveCommandCanvas)

}

type drawable interface {
  draw(t pixel.Target)
}

func drawStuff(win *pixelgl.Window) {
  //background
  win.Clear(colornames.Forestgreen)

  //gameMap.drawGrid(win)

  //entities
  drawPheremones(win)
  drawEntities(win)

  drawHighlightedTiles(win)
  drawSelectedTiles(win)

  win.Update()
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

func drawHighlightedTiles(t pixel.Target) {
  if (tileHighlightStart == nil || tileHighlightEnd == nil) {
    log.Output(1, "Warning: drawing highlighted tiles with nil start and end tiles")
    return
  }
  minCol, maxCol, minRow, maxRow := gameMap.getBoundsFromTiles(tileHighlightStart, tileHighlightEnd)

  v1 := gameMap.tiles[minCol][minRow].originPos()
  v2 := gameMap.tiles[maxCol+1][maxRow+1].originPos()

  imd := imdraw.New(nil)

  imd.Color = pixel.RGB(0,1,0)
  imd.Push(v1, v2)
  imd.Rectangle(1)

  imd.Draw(t)
}

func drawEntities(t pixel.Target) {
  for i := 0; i<len(gameEntities); i++ {
    gameEntities[i].draw(t)
  }
}

func (ant *ant) draw(t pixel.Target) {
  antCanvas.Draw(t, pixel.IM.Rotated(pixel.ZV, ant.angle).Moved(ant.pos))
}

func drawPheremones(t pixel.Target) {
  for _, ti := range(gameMap.drawableTiles()) {
    if ti.pher.command == nil {
      continue
    }
    ti.drawPheremone(t)
  }
}

func (ti* tile) drawPheremone(t pixel.Target) {
  if ti.pher.command != nil {
    pheremoneOverlayCanvas.Draw(t, pixel.IM.Moved(ti.centerPos()))
    ti.pher.command.drawOnTile(t, ti)
  }
}

func (mc *moveCommand) drawOnTile(t pixel.Target, ti *tile) {
  tileCenterPos := ti.centerPos()
  angle := mc.target.Sub(tileCenterPos).Angle()
  moveCommandCanvas.Draw(t, pixel.IM.Rotated(pixel.ZV, angle).Moved(tileCenterPos))
}

func drawSelectedTiles(t pixel.Target) {
  for _, ti := range(selectedTiles) {
    selectedTileCanvas.Draw(t, pixel.IM.Moved(ti.centerPos()))
  }
}
