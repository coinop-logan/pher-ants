package main

import (
  //"fmt"
  //"log"
  "github.com/faiface/pixel/pixelgl"
)

var (
  tileHighlightStart, tileHighlightEnd *tile
  selectedTiles []*tile
)

var cmd command

func handleInput(win *pixelgl.Window) {

  //tile highlight and selection
  tileAtMouse := gameMap.getTileAtPos(win.MousePosition())

  if win.JustPressed(pixelgl.MouseButton1) {
    if ! (win.Pressed(pixelgl.KeyLeftShift) || win.Pressed(pixelgl.KeyRightShift)) {
      selectedTiles = nil
    }
  }

  if ! win.Pressed(pixelgl.MouseButton1) {
    if win.JustReleased(pixelgl.MouseButton1) {
      addHighlightedTilesToSelection()
    }
    tileHighlightStart = tileAtMouse
    tileHighlightEnd = tileAtMouse
  } else {
    tileHighlightEnd = tileAtMouse
  }

  //assign commands to tile
  if win.JustPressed(pixelgl.MouseButton2) {
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
    }
  }
  if win.JustPressed(pixelgl.KeyQ) {
    cmd = new(moveCommand)
  }
}

func addHighlightedTilesToSelection() {
  minCol, maxCol, minRow, maxRow := gameMap.getBoundsFromTiles(tileHighlightStart, tileHighlightEnd)

  for col := minCol; col <= maxCol; col ++ {
    for row := minRow; row <= maxRow; row ++ {
      //TODO: check for duplicate
      selectedTiles = append(selectedTiles, &gameMap.tiles[col][row])
    }
  }
}
