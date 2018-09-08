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
)

type gameEntity interface {
  drawable
  iterate() error
}

type pheremone struct {
  command command
}
