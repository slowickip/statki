package connection

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	warshipServerAddr = "https://go-pjatk-server.fly.dev"
	httpClientTimeout = 30
)

type Client struct {
	httpClient *http.Client
	token      string
}

type GameRequestStruct struct {
	Wpbot bool `json:"wpbot"`
}

func (c *Client) GameInit(wpbot bool) Client {
	client := Client{httpClient: &http.Client{Timeout: httpClientTimeout}}
	values := GameRequestStruct{Wpbot: true}

	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}
	postAddr := warshipServerAddr + "/api/game"

	resp, err := client.httpClient.Post(postAddr, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	client.token = string(resp.Header.Get("x-auth-token"))

	return client
}

func (c *Client) GetStatus() StatusResponse {
	request, err := http.NewRequest("https://go-pjatk-server.fly.dev/api/game", "GET", nil)

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("x-auth-token", c.token)
	response, err := c.httpClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	result := StatusResponse{}
	err = json.Unmarshal(responseBody, &result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (c *Client) Board() []string {
	return []string{}
}
