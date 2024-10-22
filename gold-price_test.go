package main

import (
	"testing"
)

func TestGold_GetPrices(t *testing.T) {

	g := Gold{
		Prices: nil,
		Client: client,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error("Error fetching gold price: ", err)
	}

	if p.Price != 2735.5925 {
		t.Error("Expected 2735.5925, got ", p.Price)
	}
}
