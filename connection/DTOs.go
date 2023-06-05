package connection

import (
	"net/http"
)

type Client struct {
	httpClient http.Client
	token      string
}

type Stats struct {
	Games  int    `json:"games"`
	Nick   string `json:"nick"`
	Points int    `json:"points"`
	Rank   int    `json:"rank"`
	Wins   int    `json:"wins"`
}

type PlayerStats struct {
	Stats Stats `json:"stats"`
}

type Ranking struct {
	Stats []Stats `json:"stats"`
}

type GameRequestStruct struct {
	Wpbot      bool   `json:"wpbot"`
	Nick       string `json:"nick"`
	TargetNick string `json:"target_nick"`
}

type FireRequestStruct struct {
	Coord string `json:"coord"`
}

type BoardResponse struct {
	Board []string `json:"board"`
}

type FireResponse struct {
	Result string `json:"result"`
}

type StatusResponse struct {
	Nick           string   `json:"nick"`
	GameStatus     string   `json:"game_status"`
	LastGameStatus string   `json:"last_game_status"`
	Opponent       string   `json:"opponent"`
	OppShots       []string `json:"opp_shots"`
	ShouldFire     bool     `json:"should_fire"`
	Timer          int      `json:"timer"`
}

type DescResponse struct {
	Desc     string `json:"desc"`
	Nick     string `json:"nick"`
	OppDesc  string `json:"opp_desc"`
	Opponent string `json:"opponent"`
}

type Player struct {
	Nick       string `json:"nick"`
	GameStatus string `json:"game_status"`
}
