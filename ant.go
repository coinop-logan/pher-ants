package main

import (
  _ "fmt"
  "math"
  "github.com/faiface/pixel"
)

const (
  ANT_TURNING_SPEED =        0.1
  ANT_FORWARD_ACCELERATION = 0.1
  ANT_BRAKE_DECELERATION =   0.2
  ANT_SPEED_MAX =            2.0
)

type ant struct {
  pos pixel.Vec
  angle float64
  speed float64
  cmd command
}

func (a *ant) getAngleUnitVec() pixel.Vec {
  return pixel.V(math.Cos(a.angle), math.Sin(a.angle))
}

func (a *ant) moveToward(targetPos pixel.Vec) {
  //turn
  angleToTarget := targetPos.Sub(a.pos).Angle()
  a.turnTowardAngle(angleToTarget)
  angleDifference := angleToTarget - a.angle

  //move

  if math.Cos(angleDifference) < 0 {
    a.brake()
  } else {
    a.moveForward()
  }
}

func (a *ant) moveForward() {
  a.speed += ANT_FORWARD_ACCELERATION
}

func (a *ant) turnTowardAngle(targetAngle float64) {
  if math.Sin(targetAngle - a.angle) > 0 {
    a.angle += ANT_TURNING_SPEED
  } else {
    a.angle -= ANT_TURNING_SPEED
  }
}

func (a *ant) brake() {
  if a.speed > ANT_BRAKE_DECELERATION {
    a.speed -= ANT_BRAKE_DECELERATION
  } else {
    a.speed = 0
  }
}

func (a *ant) iterate() error {
  onTile := gameMap.getTileAtPos(a.pos)

  if onTile.pher.command == nil{
    //a.brake()
  } else {
    a.cmd = onTile.pher.command
  }

  if a.cmd != nil {
    a.cmd.controlAnt(a)
  }

  if a.speed > ANT_SPEED_MAX {
    a.speed = ANT_SPEED_MAX
  }
  a.pos = a.pos.Add(a.getAngleUnitVec().Scaled(a.speed))

  return nil
}
