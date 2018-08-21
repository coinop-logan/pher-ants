package main

import (
  //"log"
  "math"
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/imdraw"
)

const (
  UP =        0
  UPLEFT =    1
  LEFT =      2
  DOWNLEFT =  3
  DOWN =      4
  DOWNRIGHT = 5
  RIGHT =     6
  UPRIGHT =   7

  ANT_TURNING_SPEED =      0.1
  ANT_SPEED_ACCELERATION = 0.1
  ANT_SPEED_MAX =          0.1
)

type drawable interface {
  draw(t pixel.Target)
}

type gameEntity interface {
  drawable
  iterate() error
}

type pheremone struct {
  gameEntity
  command command
}

type command interface {
  controlAnt(a *ant)
}

type moveCommand struct {
  target pixel.Vec
}

type ant struct {
  pos pixel.Vec
  angle float64
  speed float64
}

func (cmd *moveCommand) controlAnt(ant *ant) {
  ant.moveToward(cmd.target)
}

func (ant *ant) moveToward(targetPos pixel.Vec) {
  //turn
  angleToTarget := targetPos.Sub(ant.pos).Angle()
  ant.turnTowardAngle(angleToTarget)

  //move
  if math.Cos(angleToTarget) < 0 {
    ant.moveForward()
  }
}

func (ant *ant) moveForward() {
  ant.speed += ANT_SPEED_ACCELERATION
}

func (ant *ant) turnTowardAngle(targetAngle float64) {
  if math.Sin(targetAngle - ant.angle) > 0 {
    ant.angle += ANT_TURNING_SPEED
  } else {
    ant.angle -= ANT_TURNING_SPEED
  }
}

func (ant *ant) iterate() error {
  onTile := gameMap.posToTile(ant.pos)

  onTile.pher.command.controlAnt(ant)

  if ant.speed > ANT_SPEED_MAX {
    ant.speed = ANT_SPEED_MAX
  }

  return nil
}

func (ant *ant) draw(t pixel.Target) {
  //TODO: this should be executed elsewhere, perhaps? Just once? And re-use the imd
  imd := imdraw.New(nil)

  imd.Color = pixel.RGB(0,0,1)
  imd.Push(pixel.V(0,   1   ).Scaled(10).Add(ant.pos),
           pixel.V(-0.5,-1  ).Scaled(10).Add(ant.pos),
           pixel.V(0,   -0.5).Scaled(10).Add(ant.pos),
           pixel.V(0.5, -1  ).Scaled(10).Add(ant.pos))
  imd.Polygon(0)

  imd.Draw(t)
}
