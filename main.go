package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"statki/connection"
	"statki/game"
	"strings"
)

func main() {
	exit := false
	for !exit {
		client := connection.Client{}
		gameInstance := game.New(&client)
		gameOptions := connection.GameRequestStruct{}

		fmt.Println("Do you want to insert you nick? (yes/no)")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}
		if strings.Contains(line, "yes") {
			fmt.Println("What is your nickname?")
			reader = bufio.NewReader(os.Stdin)
			line, err = reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
				return
			}
			gameOptions.Nick = strings.TrimSuffix(line, "\n")
		} else if !strings.Contains(line, "no") {
			log.Fatal("Wrong input")
			return
		}

		fmt.Println("Do you want to play with bot? (yes/no)")
		reader = bufio.NewReader(os.Stdin)
		line, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}
		if strings.Contains(line, "yes") {
			gameOptions.Wpbot = true
		} else if strings.Contains(line, "no") {
			gameOptions.Wpbot = false
			fmt.Println("Do you want to play with random player? (yes/no)")
			reader = bufio.NewReader(os.Stdin)
			line, err = reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
				return
			}
			if strings.Contains(line, "no") {
				gameOptions.TargetNick = chooseOpponent(gameInstance)
			}
		} else {
			log.Fatal("Wrong input")
			return
		}

		gameInstance.GameInit(gameOptions)
		gameInstance.Start()

		fmt.Println("Do you want to play again? (yes/no)")
		reader = bufio.NewReader(os.Stdin)
		line, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}
		if strings.Contains(line, "no") {
			exit = true
		} else if !strings.Contains(line, "yes") {
			log.Fatal("Wrong input")
			return
		}
	}
}

func chooseOpponent(game *game.Game) string {
	players := game.C.GetPlayers()
	fmt.Println("Available players:")
	for _, player := range players {
		fmt.Println(player)
	}
	fmt.Println("What is your opponent nickname?")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return strings.TrimSuffix(line, "\n")
}
