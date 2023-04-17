package connection

type StatusResponse struct {
	desc           string   `json:"desc"`
	gameStatus     string   `json:"game_status"`
	lastGameStatus string   `json:"last_game_status"`
	nick           string   `json:"nick"`
	oppDesc        string   `json:"opp_desc"`
	oppShots       []string `json:"opp_shots"`
	opponent       string   `json:"opponent"`
	shouldFire     bool     `json:"should_fire"`
	timer          int      `json:"timer"`
}
