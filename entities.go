package main

import (
  //"fmt"
  //"log"
  "math"
  "github.com/faiface/pixel"
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

  ANT_TURNING_SPEED =        0.1
  ANT_FORWARD_ACCELERATION = 0.1
  ANT_BRAKE_DECELERATION =   0.2
  ANT_SPEED_MAX =            1.0
)

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
  rightClick(pos pixel.Vec)
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

type ant struct {
  pos pixel.Vec
  angle float64
  speed float64
  cmd command
}

func (ant *ant) getAngleUnitVec() pixel.Vec {
  return pixel.V(math.Cos(ant.angle), math.Sin(ant.angle))
}

func (ant *ant) moveToward(targetPos pixel.Vec) {
  //turn
  angleToTarget := targetPos.Sub(ant.pos).Angle()
  ant.turnTowardAngle(angleToTarget)

  //move
  if math.Cos(angleToTarget) > 0 {
    ant.moveForward()
  }
}

func (ant *ant) moveForward() {
  ant.speed += ANT_FORWARD_ACCELERATION
}

func (ant *ant) turnTowardAngle(targetAngle float64) {
  if math.Sin(targetAngle - ant.angle) > 0 {
    ant.angle += ANT_TURNING_SPEED
  } else {
    ant.angle -= ANT_TURNING_SPEED
  }
}

func (ant *ant) brake() {
  if ant.speed > ANT_BRAKE_DECELERATION {
    ant.speed -= ANT_BRAKE_DECELERATION
  } else {
    ant.speed = 0
  }
}

func (ant *ant) iterate() error {
  onTile := gameMap.getTileAtPos(ant.pos)

  if onTile.pher.command == nil{
    //ant.brake()
  } else {
    ant.cmd = onTile.pher.command
  }

  if ant.cmd != nil {
    ant.cmd.controlAnt(ant)
  }

  if ant.speed > ANT_SPEED_MAX {
    ant.speed = ANT_SPEED_MAX
  }
  ant.pos = ant.pos.Add(ant.getAngleUnitVec().Scaled(ant.speed))

  return nil
}
