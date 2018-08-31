package main

import (
  "math"
  "github.com/faiface/pixel"
)

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
