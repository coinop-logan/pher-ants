package main

import (
  //"fmt"
  //"log"
  "math"
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/imdraw"
)

const (
  TILE_WIDTH = 32
  MAP_WIDTH_IN_TILES = 1024
)

type tile struct {
  col, row uint16
  pher pheremone
  ant *ant
}

func (t *tile) getRect() pixel.Rect {
  return pixel.R(float64(t.col*TILE_WIDTH), float64(t.row*TILE_WIDTH), float64((t.col+1)*TILE_WIDTH), float64((t.row+1)*TILE_WIDTH))
}

type gameMapType struct {
  tiles [MAP_WIDTH_IN_TILES][MAP_WIDTH_IN_TILES]tile
}

func (m *gameMapType) init() {
  for col := uint16(0); col < MAP_WIDTH_IN_TILES; col++ {
    for row := uint16(0); row < MAP_WIDTH_IN_TILES; row++ {
      t := &m.tiles[col][row]
      t.col = col
      t.row = row
    }
  }
}

func (m *gameMapType) getTileAtPos(pos pixel.Vec) *tile {
  tileCol := uint16(math.Floor(pos.X / float64(TILE_WIDTH)))
  tileRow := uint16(math.Floor(pos.Y / float64(TILE_WIDTH)))
  return &m.tiles[tileCol][tileRow]
}

func (t *tile) centerPos() (pixel.Vec) {
  return t.getRect().Center()
}

func (t *tile) originPos() (pixel.Vec) {
  return pixel.V(float64(t.col), float64(t.row)).Scaled(float64(TILE_WIDTH))
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
