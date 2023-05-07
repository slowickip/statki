package game

import (
	"context"
	gui "github.com/grupawp/warships-gui/v2"
	"statki/connection"
	"strconv"
	"time"
)

type Game struct {
	c        connection.Client
	status   connection.StatusResponse
	myShips  [10][10]gui.State
	oppShips [10][10]gui.State
}

type Coordinate struct {
	X int
	Y int
}

func (g *Game) Start() {
	g.status = g.c.GetStatus()
	for !g.Started() {
		time.Sleep(1 * time.Second)
		g.status = g.c.GetStatus()
	}
	ui := gui.NewGUI(true)
	myBoard := gui.NewBoard(1, 3, nil)
	oppBoard := gui.NewBoard(50, 3, nil)

	txt := gui.NewText(1, 1, "Press Ctrl+C to exit", nil)
	statusText := gui.NewText(30, 1, "", nil)
	desc := g.c.GetDesc()
	myNick := gui.NewText(1, 25, desc.Nick, nil)
	myDesc := gui.NewText(1, 26, desc.Desc, nil)
	oppNick := gui.NewText(1, 28, desc.Opponent, nil)
	oppDesc := gui.NewText(1, 29, desc.OppDesc, nil)

	ui.Draw(txt)
	ui.Draw(statusText)
	ui.Draw(myBoard)
	ui.Draw(oppBoard)
	ui.Draw(myNick)
	ui.Draw(myDesc)
	ui.Draw(oppNick)
	ui.Draw(oppDesc)
	go func() {
		ui.Start(nil)
	}()

	g.importShips()
	go func() {
		for {
			if !g.status.ShouldFire {
				statusText.SetText("Opponent's turn")
			} else if g.status.ShouldFire {
				statusText.SetText("Your turn")
				coordString := oppBoard.Listen(context.TODO())
				result := g.c.Fire(coordString)
				coord := parseCoordinate(coordString)
				if result.Result == "hit" {
					g.oppShips[coord.X][coord.Y] = gui.Hit
				} else if result.Result == "miss" {
					g.oppShips[coord.X][coord.Y] = gui.Miss
				} else if result.Result == "sunk" {
					g.oppShips[coord.X][coord.Y] = gui.Hit
				}
				oppBoard.SetStates(g.oppShips)
				ui.Log("Coordinate: %s", coord)
			}
		}
	}()

	for g.Started() {
		g.status = g.c.GetStatus()
		g.updateMyShips()
		myBoard.SetStates(g.myShips)
		time.Sleep(1 * time.Second)
	}
}

func (g *Game) GameInit(wpbot bool) {
	g.c.GameInit(wpbot)
}

func (g *Game) Started() bool {
	return g.status.GameStatus == "game_in_progress"
}

func New(client *connection.Client) *Game {
	return &Game{
		c: *client,
	}
}

func (g *Game) importShips() {
	myShips := g.c.GetBoard()
	states := [10][10]gui.State{}
	for i := range states {
		states[i] = [10]gui.State{}
	}
	for _, ship := range myShips {
		states[ship[0]-'A'][ship[1]-'0'-1] = gui.Ship
	}
	g.myShips = states
}

func (g *Game) updateMyShips() {
	oShots := g.status.OppShots
	for _, shot := range oShots {
		coord := parseCoordinate(shot)
		g.myShips[coord.X][coord.Y] = gui.Hit
	}
}

func parseCoordinate(coordString string) Coordinate {
	convY, _ := strconv.Atoi(coordString[1:])
	convY--
	coord := Coordinate{
		X: int(coordString[0] - 'A'),
		Y: convY,
	}
	return coord
}
