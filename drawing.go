package main

import (
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
  "github.com/faiface/pixel/imdraw"
)

var antCanvas *pixelgl.Canvas

func generatePictures() {

  imd := imdraw.New(nil)

  //ant

  antCanvas = pixelgl.NewCanvas(pixel.R(0,0,16,16))

  imd.Color = pixel.RGB(0,0,1)
  imd.Push(pixel.V(1,    0   ).Scaled(10).Add(pixel.V(8,8)),
           pixel.V(-1,   -0.5).Scaled(10).Add(pixel.V(8,8)),
           pixel.V(-0.5, 0   ).Scaled(10).Add(pixel.V(8,8)),
           pixel.V(-1,   0.5 ).Scaled(10).Add(pixel.V(8,8)))
  imd.Polygon(0)

  imd.Draw(antCanvas)

  imd.Clear()
}

type drawable interface {
  draw(t pixel.Target)
}

func (ant *ant) draw(t pixel.Target) {
  antCanvas.Draw(t, pixel.IM.Rotated(pixel.ZV, ant.angle).Moved(ant.pos))
}
