package main

import "C"
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"statki/connection"
	"statki/game"
	"strings"
)

type Core struct {
	client       connection.Client
	gameInstance *game.Game
	gameOptions  connection.GameRequestStruct
}

func main() {
	core := Core{
		client:      connection.Client{},
		gameOptions: connection.GameRequestStruct{},
	}
	core.gameInstance = game.New(&core.client)
	core.gameOptions.Wpbot = true

	fmt.Println("Welcome to Warships!")

	for true {
		core.menu()
	}
}

func (core *Core) setNick() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("What is your nickname?")
	reader = bufio.NewReader(os.Stdin)
	line, err = reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println("Your nickname is: " + strings.TrimSuffix(line, "\n"))
	core.gameOptions.Nick = strings.TrimSuffix(line, "\n")
}

func (core *Core) setWPBot() {
	fmt.Println("Do you want to play with bot? (yes/no)")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	if strings.Contains(line, "yes") {
		core.gameOptions.Wpbot = true
	} else if strings.Contains(line, "no") {
		core.gameOptions.Wpbot = false
	} else {
		log.Println("Wrong input")
		return
	}
}

func (core *Core) chooseWhetherRandomOpp() {
	fmt.Println("Do you want to play with random player? (yes/no)")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	if strings.Contains(line, "no") {
		core.chooseOpponent()
	}
}

func (core *Core) chooseOpponent() {
	players := core.gameInstance.C.GetPlayers()
	fmt.Println("Available players:")
	for _, player := range players {
		fmt.Println(player)
	}
	fmt.Println("What is your opponent nickname?")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	core.gameOptions.TargetNick = strings.TrimSuffix(line, "\n")
}

func (core *Core) startGame() {
	core.gameInstance.GameInit(core.gameOptions)
	core.gameInstance.Start()
}

func (core *Core) menu() {
	fmt.Println()
	if core.gameOptions.Nick == "" {
		fmt.Println("Nickname not set")
	} else {
		fmt.Println("Nickname: " + core.gameOptions.Nick)
	}
	if core.gameOptions.TargetNick == "" {
		fmt.Println("Opponent not set")
	} else {
		fmt.Println("Opponent: " + core.gameOptions.TargetNick)
	}
	if core.gameOptions.Wpbot == true {
		fmt.Println("Playing with bot")
	} else {
		fmt.Println("Playing with player")
	}
	fmt.Println()
	fmt.Println("Choose option:")
	fmt.Println("1. Set nickname")
	fmt.Println("2. Set opponent")
	fmt.Println("3. Set whether playing with bot")
	fmt.Println("4. Start game")
	fmt.Println("5. Ranking")
	fmt.Println("6. Exit")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	option := strings.TrimSuffix(line, "\n")
	switch option {
	case "1":
		core.setNick()
	case "2":
		core.chooseWhetherRandomOpp()
	case "3":
		core.setWPBot()
	case "4":
		core.startGame()
	case "5":
		core.printRanking()
	case "6":
		os.Exit(0)
	default:
		fmt.Println("Wrong input")
	}
}

func (core *Core) printRanking() {
	yourStats := core.gameInstance.C.GetStats(core.gameOptions.Nick)
	ranking := core.gameInstance.C.GetRanking()

	fmt.Println("Your stats:")
	printRankingForPlayer(yourStats.Stats)

	fmt.Println("Ranking:")
	for _, player := range ranking.Stats {
		printRankingForPlayer(player)
	}
}

func printRankingForPlayer(stats connection.Stats) {
	fmt.Printf("%d ", stats.Rank)
	fmt.Printf("%s ", stats.Nick)
	fmt.Printf("Punkty: %d ", stats.Points)
	fmt.Printf("Wygrane: %d ", stats.Wins)
	fmt.Printf("Ilość gier: %d", stats.Games)
	fmt.Println()
}
