package main

import (
  "github.com/faiface/pixel"
)

type command interface {
  controlAnt(a *ant)
  rightClick(pos pixel.Vec)
  drawOnTile(t pixel.Target, ti *tile)
}

type moveCommand struct {
  target pixel.Vec
}

func (cmd *moveCommand) rightClick(pos pixel.Vec) {
  cmd.target = pos
}

func (cmd *moveCommand) controlAnt(ant *ant) {
  ant.moveToward(cmd.target)
}
