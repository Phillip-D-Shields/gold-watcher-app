package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var currency = "USD"

type Gold struct {
	Prices []Price `json:"items"`
	Client *http.Client
}

type Price struct {
	Currency      string    `json:"currency"`
	Price         float64   `json:"xauPrice"`
	Change        float64   `json:"chgXau"`
	PreviousClose float64   `json:"xauClose"`
	Time          time.Time `json:"-"`
}

func (g *Gold) GetPrices() (*Price, error) {
	if g.Client == nil {
		g.Client = &http.Client{}
	}
	client := g.Client

	url := fmt.Sprintf("https://data-asg.goldprice.org/dbXRates/%s", currency)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	res, err := client.Do(req)
	if err != nil {
		log.Println("Error fetching gold price: ", err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading gold price response: ", err)
		return nil, err
	}

	gold := Gold{}

	var previous, current, change float64

	err = json.Unmarshal(body, &gold)
	if err != nil {
		log.Println("Error unmarshalling gold price: ", err)
		return nil, err
	}

	previous, current, change = gold.Prices[0].PreviousClose, gold.Prices[0].Price, gold.Prices[0].Change

	var currentInfo = Price{
		Currency:      currency,
		Price:         current,
		Change:        change,
		PreviousClose: previous,
		Time:          time.Now(),
	}

	return &currentInfo, nil
}
