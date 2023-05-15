package connection

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
