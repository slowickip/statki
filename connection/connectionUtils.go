package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	warshipServerAddr = "https://go-pjatk-server.fly.dev"
	httpClientTimeout = 0
)

type Client struct {
	httpClient http.Client
	token      string
}

//type client interface {
//	http.Client
//	string
//	GameInit(wpbot bool) error
//	//GameInit(coords []string, desc, nick, target_nick string, wpbot bool) error
//	GetStatus(token string) (*StatusResponse, error)
//	Board() ([]string, error)
//}

type GameRequestStruct struct {
	Wpbot bool `json:"wpbot"`
}

func (c *Client) GameInit(wpbot bool) {
	c.httpClient = http.Client{
		Timeout: httpClientTimeout,
	}

	values := GameRequestStruct{Wpbot: wpbot}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}
	postAddr := warshipServerAddr + "/api/game"

	resp, err := c.httpClient.Post(postAddr, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	c.token = resp.Header.Get("x-auth-token")
}

func (c *Client) GetStatus() StatusResponse {

	statusAddr := warshipServerAddr + "/api/game"
	request, err := http.NewRequest(http.MethodGet, statusAddr, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c.token)
	request.Header.Set("x-auth-token", c.token)
	request.Header.Set("content-type", "application/json")

	response, err := c.httpClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var result StatusResponse
	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	return result
}

func (c *Client) GetBoard() StatusResponse {

	boardAddr := warshipServerAddr + "/api/game/board"
	request, err := http.NewRequest(http.MethodGet, boardAddr, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c.token)
	request.Header.Set("x-auth-token", c.token)
	request.Header.Set("content-type", "application/json")

	response, err := c.httpClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var result StatusResponse
	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	return result
}
