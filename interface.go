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
    if cmd != nil {
      cmd.rightClick(win.MousePosition())
      for i := 0; i < len(selectedTiles); i++ {
        selectedTiles[i].pher = pheremone{command:cmd}
      }
      cmd = nil
      selectedTiles = nil
    } else {
      if win.Pressed(pixelgl.KeyLeftControl) || win.Pressed(pixelgl.KeyRightControl) {

      } else if ! (win.Pressed(pixelgl.KeyLeftShift) || win.Pressed(pixelgl.KeyRightShift)) {
        selectedTiles = nil
      }
    }
  }

  if win.JustReleased(pixelgl.MouseButton1) {
    addHighlightedTilesToSelection()
    tileHighlightStart = tileAtMouse
    tileHighlightEnd = tileAtMouse
  }

  if win.Pressed(pixelgl.MouseButton1) {
    tileHighlightEnd = tileAtMouse
  } else {
    tileHighlightStart = tileAtMouse
    tileHighlightEnd = tileAtMouse
  }



  if win.JustPressed(pixelgl.KeyEscape) {
    if selectedTiles != nil {
      selectedTiles = nil
    } else {
      cmd = nil
    }
  }
  if win.JustPressed(pixelgl.KeyQ) {
    var mvCmd = command(new(moveCommand))
    cmd = mvCmd
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
