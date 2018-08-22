package main

import (
  "fmt"
  "log"
  "github.com/faiface/pixel/pixelgl"
)

var (
  highlightedTile *tile
  selectedTiles []*tile
)

var cmd command

func handleInput(win *pixelgl.Window) {
  highlightedTile = gameMap.getTileAtPos(win.MousePosition())

  if win.JustPressed(pixelgl.MouseButton1) {
    t := gameMap.getTileAtPos(win.MousePosition())
    selectedTiles = append(selectedTiles, t)
  } else if win.JustPressed(pixelgl.MouseButton2) {
    if cmd != nil {
      cmd.rightClick(win.MousePosition())
      for i := 0; i < len(selectedTiles); i++ {
        selectedTiles[i].pher = pheremone{command:cmd}
      }
    }
  }

  if win.JustPressed(pixelgl.KeyEscape) {
    if selectedTiles != nil {
      selectedTiles = nil
    } else {
      cmd = nil
      log.Output(1, fmt.Sprintf("%v", cmd))
    }
  }
  if win.JustPressed(pixelgl.KeyQ) {
    cmd = new(moveCommand)
  }
}
