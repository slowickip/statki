package main

import (
	"statki/connection"
	"statki/game"
)

func main() {
	client := connection.Client{}
	gameInstance := game.New(&client)

	gameInstance.GameInit(true)
	gameInstance.Start()
}
