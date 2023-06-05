package game

import (
	"context"
	gui "github.com/grupawp/warships-gui/v2"
	"statki/connection"
	"strconv"
	"time"
)

type Coordinate struct {
	X int
	Y int
}

type Game struct {
	C        connection.Client
	status   connection.StatusResponse
	myShips  [10][10]gui.State
	oppShips [10][10]gui.State
	shots    int
	hits     int
}

func (g *Game) Start() {
	g.status = g.C.GetStatus()
	for !g.Started() {
		time.Sleep(1 * time.Second)
		g.status = g.C.GetStatus()
	}
	g.shots = 0
	g.hits = 0
	ui := gui.NewGUI(true)
	myBoard := gui.NewBoard(1, 6, nil)
	oppBoard := gui.NewBoard(50, 6, nil)
	timer := gui.NewText(50, 1, "", nil)
	accuracy := gui.NewText(50, 3, "Accuracy: ", nil)
	txt := gui.NewText(1, 1, "Press Ctrl+C to exit", nil)
	statusText := gui.NewText(1, 3, "", nil)
	desc := g.C.GetDesc()
	myNick := gui.NewText(1, 30, desc.Nick, nil)
	myDesc := gui.NewText(1, 31, desc.Desc, nil)
	oppNick := gui.NewText(1, 33, desc.Opponent, nil)
	oppDesc := gui.NewText(1, 34, desc.OppDesc, nil)
	ctx := context.Background()

	ui.Draw(timer)
	ui.Draw(accuracy)
	ui.Draw(txt)
	ui.Draw(statusText)
	ui.Draw(myBoard)
	ui.Draw(oppBoard)
	ui.Draw(myNick)
	ui.Draw(myDesc)
	ui.Draw(oppNick)
	ui.Draw(oppDesc)
	go func() {
		ui.Start(ctx, nil)
		g.C.AbandonGame()
	}()

	g.importShips()
	go func() {
		for g.Started() {
			c := make(chan bool)
			if !g.status.ShouldFire {
				statusText.SetText("Opponent's turn")
				c <- false
			} else if g.status.ShouldFire {
				statusText.SetText("Your turn")
				go g.startTimer(timer, c)
				coordString := oppBoard.Listen(context.TODO())
				c <- false
				g.shots++
				result := g.C.Fire(coordString)
				coord := parseCoordinate(coordString)
				if result.Result == "hit" {
					g.hits++
					g.oppShips[coord.X][coord.Y] = gui.Hit
				} else if result.Result == "miss" {
					g.oppShips[coord.X][coord.Y] = gui.Miss
				} else if result.Result == "sunk" {
					g.hits++
					g.oppShips[coord.X][coord.Y] = gui.Hit
				}
				oppBoard.SetStates(g.oppShips)
				accuracy.SetText("Accuracy: " + strconv.Itoa(g.hits*100/g.shots) + "%")
				//ui.Log("Coordinate: %s", coord)
			}
		}
	}()

	for g.Started() {
		g.status = g.C.GetStatus()
		g.updateMyShips()
		myBoard.SetStates(g.myShips)
		time.Sleep(1 * time.Second)
	}

	g.status = g.C.GetStatus()
	if g.status.GameStatus == "ended" {
		if g.status.LastGameStatus == "win" {
			statusText.SetText("You won!")
		} else {
			statusText.SetText("You lost!")
		}
	}
	time.Sleep(1 * time.Second)
}

func (g *Game) GameInit(requestStruct connection.GameRequestStruct) {
	g.C.GameInit(requestStruct)
}

func (g *Game) Started() bool {
	return g.status.GameStatus == "game_in_progress"
}

func New(client *connection.Client) *Game {
	return &Game{
		C: *client,
	}
}

func (g *Game) startTimer(timer *gui.Text, event chan bool) {
	ticker := time.NewTicker(time.Second)
	i := 60
	for i > 0 {
		select {
		case <-ticker.C:
			i--
			timer.SetText(strconv.Itoa(i) + "s")
			if i == 0 {
				g.C.AbandonGame()
				timer.SetText("Oddałeś grę walkoverem")
			}
		case <-event:
			return
		}
	}

}

func (g *Game) importShips() {
	myShips := g.C.GetBoard()
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
