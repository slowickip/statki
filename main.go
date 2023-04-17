package main

import (
	"statki/connection"
	"statki/game"
)

func main() {
	client := connection.Client{}
	client.GameInit()

	game := game.New(&client)

}
