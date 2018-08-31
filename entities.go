package main

import (
  //"fmt"
  //"log"
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
