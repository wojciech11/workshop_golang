package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HeroClient struct {
	Client *http.Client
	Url    string
}

type Squad struct {
	SquadName string `json:"squadName"`
	HomeTown  string `json:"homeTown"`
	Members   []SquadMember
}

type SquadMember struct {
	Name string
}

func (hc *HeroClient) GetThem() (*Squad, error) {
	resp, err := hc.Client.Get(hc.Url)
	if err != nil {
		return nil, err
	}

	s, _ := io.ReadAll(resp.Body)

	sq := Squad{}

	err = json.Unmarshal(s, &sq)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal: %w", err)
	}

	return &sq, err
}

func main() {
	c := &http.Client{
		Timeout: 2 * time.Second,
	}
	heroClient := HeroClient{Client: c,
		Url: "https://mdn.github.io/learning-area/javascript/oojs/json/superheroes.json",
	}

	squad, err := heroClient.GetThem()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", squad)

}
