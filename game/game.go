package game

import (
	"statki/connection"
)

type Game struct {
	c connection.Client
}

type client interface {
	GameInit(wpbot bool) error
	//GameInit(coords []string, desc, nick, target_nick string, wpbot bool) error
	GetStatus(token string) (*connection.StatusResponse, error)
	Board() ([]string, error)
}

func New(c client) *Game {
	return &Game{
		c,
	}
}
