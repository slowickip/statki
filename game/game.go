package game

import (
	"github.com/grupawp/warships-lightgui/v2"
	"statki/connection"
	"time"
)

type Game struct {
	c connection.Client
}

func (g *Game) Start() {
	for !g.Started() {
		time.Sleep(1 * time.Second)
	}
	gameBoard := board.New(board.NewConfig())
	for true {
		time.Sleep(1 * time.Second)
		gameBoard.Import(g.c.GetBoard())
		gameBoard.Display()
	}
}

func (g *Game) GameInit(wpbot bool) {
	g.c.GameInit(wpbot)
}

func (g *Game) Started() bool {
	return g.c.GetStatus().GameStatus == "game_in_progress"
}

func New(client *connection.Client) *Game {
	return &Game{
		c: *client,
	}
}
