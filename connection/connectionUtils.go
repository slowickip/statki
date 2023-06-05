package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	warshipServerAddr = "https://go-pjatk-server.fly.dev"
	httpClientTimeout = 0
)

func (c *Client) GameInit(values GameRequestStruct) {
	c.httpClient = http.Client{
		Timeout: httpClientTimeout,
	}

	json_data, err := json.Marshal(values)
	if err != nil {
		log.Println(err)
	}
	postAddr := warshipServerAddr + "/api/game"

	//resp, err := c.httpClient.Post(postAddr, "application/json", bytes.NewBuffer(json_data))
	//if err != nil {
	//	log.Println(err)
	//}

	resp := &http.Response{}
	ctr := 0
	for ctr < 5 {
		resp, err = c.httpClient.Post(postAddr, "application/json", bytes.NewBuffer(json_data))
		if err != nil {
			//log.Println(err)
			time.Sleep(1 * time.Second)
			ctr++
		} else {
			break
		}
	}

	c.token = resp.Header.Get("x-auth-token")
}

func (c *Client) GetStatus() StatusResponse {

	statusAddr := warshipServerAddr + "/api/game"
	request, err := http.NewRequest(http.MethodGet, statusAddr, nil)
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("x-auth-token", c.token)
	request.Header.Set("content-type", "application/json")

	response := &http.Response{}
	ctr := 0
	for ctr < 3 {
		response, err = c.httpClient.Do(request)
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
			ctr++
		} else {
			break
		}
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var result StatusResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Println(err)
	}

	return result
}

func (c *Client) GetStats(nick string) PlayerStats {
	boardAddr := warshipServerAddr + "/api/stats/" + nick
	request, err := http.NewRequest(http.MethodGet, boardAddr, nil)
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("x-auth-token", c.token)
	request.Header.Set("content-type", "application/json")

	response := &http.Response{}
	ctr := 0
	for ctr < 5 {
		response, err = c.httpClient.Do(request)
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
			ctr++
		} else {
			break
		}
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var result PlayerStats
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Println(err)
	}

	return result
}

func (c *Client) GetRanking() Ranking {
	boardAddr := warshipServerAddr + "/api/stats"
	request, err := http.NewRequest(http.MethodGet, boardAddr, nil)
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("x-auth-token", c.token)
	request.Header.Set("content-type", "application/json")

	response := &http.Response{}
	ctr := 0
	for ctr < 5 {
		response, err = c.httpClient.Do(request)
		if err != nil {
			//log.Println(err)
			time.Sleep(1 * time.Second)
			ctr++
		} else {
			break
		}
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var result Ranking
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Println(err)
	}

	return result
}

func (c *Client) GetBoard() []string {
	boardAddr := warshipServerAddr + "/api/game/board"
	request, err := http.NewRequest(http.MethodGet, boardAddr, nil)
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("x-auth-token", c.token)
	request.Header.Set("content-type", "application/json")

	response := &http.Response{}
	ctr := 0
	for ctr < 5 {
		response, err = c.httpClient.Do(request)
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
			ctr++
		} else {
			break
		}
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var result BoardResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Println(err)
	}

	return result.Board
}

func (c *Client) GetDesc() DescResponse {
	descAddr := warshipServerAddr + "/api/game/desc"
	request, err := http.NewRequest(http.MethodGet, descAddr, nil)
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("x-auth-token", c.token)
	request.Header.Set("content-type", "application/json")

	response := &http.Response{}
	ctr := 0
	for ctr < 5 {
		response, err = c.httpClient.Do(request)
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
			ctr++
		} else {
			break
		}
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var result DescResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Println(err)
	}

	return result
}

func (c *Client) Fire(coord string) FireResponse {
	c.httpClient = http.Client{
		Timeout: httpClientTimeout,
	}

	values := FireRequestStruct{Coord: coord}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Println(err)
	}
	postAddr := warshipServerAddr + "/api/game/fire"

	request, err := http.NewRequest(http.MethodPost, postAddr, bytes.NewBuffer(json_data))
	request.Header.Set("x-auth-token", c.token)

	resp := &http.Response{}
	ctr := 0
	for ctr < 5 {
		resp, err = c.httpClient.Do(request)
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
			ctr++
		} else {
			break
		}
	}

	var result FireResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println(err)
	}

	return result
}

func (c *Client) GetPlayers() []Player {
	playersAddr := warshipServerAddr + "/api/game/list"
	request, err := http.NewRequest(http.MethodGet, playersAddr, nil)
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("x-auth-token", c.token)
	request.Header.Set("content-type", "application/json")

	response := &http.Response{}
	ctr := 0
	for ctr < 5 {
		response, err = c.httpClient.Do(request)
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
			ctr++
		} else {
			break
		}
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var result []Player
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Println(err)
	}

	return result
}

func (c *Client) AbandonGame() {
	fmt.Println("Abandoning game...")
	abandonAddr := warshipServerAddr + "/api/game/abandon"
	request, err := http.NewRequest(http.MethodDelete, abandonAddr, nil)
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("x-auth-token", c.token)
	request.Header.Set("content-type", "application/json")

	_, err = c.httpClient.Do(request)
	if err != nil {
		log.Println(err)
	}
}
